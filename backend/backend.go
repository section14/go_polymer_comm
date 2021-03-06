package backend

import (
	//standard library
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
	//"appengine"

	//third party
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	//"github.com/gorilla/securecookie"

	//my imports
	"github.com/section14/go_polymer_comm_pkg/controller"
	"github.com/section14/go_polymer_comm_pkg/user-auth"
)

func init() {
	r := mux.NewRouter();

	//user
	r.HandleFunc("/api/", handler)
	r.HandleFunc("/api/user", UserGetHandler).Methods("GET")
	r.HandleFunc("/api/user", UserCreateHandler).Methods("POST")
	r.HandleFunc("/api/user/email", UserGetEmailHandler).Methods("POST")
	r.HandleFunc("/api/user/login", UserLoginHandler).Methods("POST")
	r.HandleFunc("/api/user/logout", UserLogoutHandler).Methods("GET")

	//user address
	r.HandleFunc("/api/user/address", UserGetAllAddressHandler).Methods("GET")
	r.HandleFunc("/api/user/address", UserSetAddressHandler).Methods("POST")
	r.HandleFunc("/api/user/address/update", UserUpdateAddressHandler).Methods("POST")
	r.HandleFunc("/api/user/address/{addressId}", UserGetAddressHandler).Methods("GET")
	//r.HandleFunc("/api/user/admin", AdminCreateHandler).Methods("GET")

	//admin
	r.HandleFunc("/api/admin/login", AdminLoginHandler).Methods("POST")
	r.HandleFunc("/api/admin", AdminVerifyHandler).Methods("GET")

	//product
	r.HandleFunc("/api/category", CategoriesGetHandler).Methods("GET")
	r.HandleFunc("/api/category/all", CategoriesGetAllHandler).Methods("GET")
	r.HandleFunc("/api/category/tree", CategoryGetTree).Methods("GET")
	r.HandleFunc("/api/category", CategoryCreateHandler).Methods("POST")
	r.HandleFunc("/api/product", ProductCreateHandler).Methods("POST")
	r.HandleFunc("/api/product/{id}", ProductGetHandler).Methods("GET")

	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
    //do a re-direct or something here
	fmt.Fprint(w, "Root level, nothing here")
}

/*

Everything below needs error handling

*/

/*
func AdminCreateHandler(w http.ResponseWriter, r *http.Request) {
	userController := controller.User{}

	status, err := userController.CreateAdmin(r)

	if err != nil {
		log.Println(err)
	}

	log.Println(status)
}
*/

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	userController := controller.User{}

	//verify user
	jwtToken := auth.JwtToken{}
	userToken,err := jwtToken.ParseToken(r)

	if err != nil {
		http.Error(w, "Invalid user", 400)
	} else {
		userId := userToken.Claims["userId"].(float64)
		user, err := userController.GetUser(r,int64(userId))

		if err != nil {
			http.Error(w, "Invalid user", 400)
		} else {
			jsonRes,_ := json.Marshal(user)
			fmt.Fprint(w, string(jsonRes))
		}
	}
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	userController := controller.User{}
	_, err := userController.CreateUser(w,r)

	if err != nil {
		log.Println(err)
	}
}

func UserGetEmailHandler(w http.ResponseWriter, r *http.Request) {
	userController := controller.User{}
	userStatus, err := userController.CheckEmail(w,r)

	if err != nil {
		//handle err
	}

	fmt.Fprint(w, userStatus)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	type JwtToken struct {
		Token string `json:"token"`
		Status bool `json:"status"`
	}

	userController := controller.User{}
	user, err := userController.Login(w,r)

	var jsonJwt *JwtToken

	if err != nil {
		//incorrect login data
		jsonJwt = &JwtToken {
			Token: " ",
			Status: false,
		}
	} else {
		//issue jwt token
		jwtToken := auth.JwtToken{}
		jwt := jwtToken.GenerateToken(user.Id, user.Role)
		jsonJwt = &JwtToken {
			Token: jwt,
			Status: true,
		}
	}

	token,_ := json.Marshal(jsonJwt)
	fmt.Fprint(w, string(token))
}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
	/*

	This method may be unnecessary since the frontend will be
	logging the user out by removing the token

	*/

	type JwtToken struct {
		Token string `json:"token"`
		Status bool `json:"status"`
	}

	var jsonJwt *JwtToken

	jsonJwt = &JwtToken {
		Token: " ",
		Status: false,
	}

	token,_ := json.Marshal(jsonJwt)
	fmt.Fprint(w, string(token))
}

func UserGetAllAddressHandler(w http.ResponseWriter, r *http.Request) {
	addressController := controller.Address{}

	//verify user
	jwtToken := auth.JwtToken{}
	userToken,err := jwtToken.ParseToken(r)

	if err != nil {
		http.Error(w, "Invalid user", 400)
	} else {
		userId := userToken.Claims["userId"].(float64)
		addressController.UserId = int64(userId)

		addresses, err := addressController.GetAllAddress(r)

		if err != nil {
			http.Error(w, "Invalid user", 400)
		} else {
			jsonRes,_ := json.Marshal(addresses)
			fmt.Fprint(w, string(jsonRes))
		}
	}
}

func UserGetAddressHandler(w http.ResponseWriter, r *http.Request) {
	addressController := controller.Address{}

	//verify user
	jwtToken := auth.JwtToken{}
	userToken,err := jwtToken.ParseToken(r)

	if err != nil {
		http.Error(w, "Invalid user", 400)
	} else {
		userId := userToken.Claims["userId"].(float64)
		addressController.UserId = int64(userId)
		address, err := addressController.GetAddress(r)

		if err != nil {
			http.Error(w, "Invalid user", 400)
		} else {
			jsonRes,_ := json.Marshal(address)
			fmt.Fprint(w, string(jsonRes))
		}
	}
}

func UserSetAddressHandler(w http.ResponseWriter, r *http.Request) {
	addressController := controller.Address{}

	//verify user
	jwtToken := auth.JwtToken{}
	userToken,err := jwtToken.ParseToken(r)

	if err != nil {
		http.Error(w, "Invalid user", 400)
	} else {
		userId := userToken.Claims["userId"].(float64)
		addressController.UserId = int64(userId)
		_, err := addressController.CreateAddress(r)

		if err != nil {
			http.Error(w, "Invalid user", 400)
		}
	}

	fmt.Fprint(w, true)
}

func UserUpdateAddressHandler(w http.ResponseWriter, r *http.Request) {
	addressController := controller.Address{}

	//verify user
	jwtToken := auth.JwtToken{}
	userToken,err := jwtToken.ParseToken(r)

	if err != nil {
		http.Error(w, "Invalid user", 400)
	} else {
		userId := userToken.Claims["userId"].(float64)
		addressController.UserId = int64(userId)
		_, err := addressController.UpdateAddress(r)

		if err != nil {
			http.Error(w, "Invalid user", 400)
		}
	}

	fmt.Fprint(w, true)
}

/*

Admin

*/

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	type JwtToken struct {
		Token string `json:"token"`
		Status bool `json:"status"`
	}

	userController := controller.User{}
	user, err := userController.Login(w,r)

	var jsonJwt *JwtToken

	//ensure role is admin (2)
	if err != nil || user.Role == 1 {
		//incorrect login data
		jsonJwt = &JwtToken {
			Token: " ",
			Status: false,
		}
	} else {
		//issue jwt token
		jwtToken := auth.JwtToken{}
		jwt := jwtToken.GenerateToken(user.Id, user.Role)
		jsonJwt = &JwtToken {
			Token: jwt,
			Status: true,
		}
	}

	token,_ := json.Marshal(jsonJwt)
	fmt.Fprint(w, string(token))
}

func AdminVerifyHandler(w http.ResponseWriter, r *http.Request) {
	adminAuth := auth.Auth{}
	status := adminAuth.VerifyAdmin(r)

	if status != true {
		http.Error(w, "Invalid user", 400)
	} else {
		fmt.Fprint(w, true)
	}
}

/*

Category

*/

func CategoryCreateHandler(w http.ResponseWriter, r *http.Request) {
	categoryController := controller.Category{}

	//verify user is admin
	adminAuth := auth.Auth{}
	status := adminAuth.VerifyAdmin(r)

	if status == true {

		_, err := categoryController.CreateCategory(r)

		if err != nil {
			log.Println(err)
		} else {
			fmt.Fprint(w, true)
		}
	}
}

func CategoriesGetHandler(w http.ResponseWriter, r *http.Request) {
	categoryController := controller.Category{}
	categories, err := categoryController.GetCategories(r)

	if err != nil {
		log.Println(err)
	} else {
		jsonRes,_ := json.Marshal(categories)
		fmt.Fprint(w, string(jsonRes))
	}
}

func CategoriesGetAllHandler(w http.ResponseWriter, r *http.Request) {
	categoryController := controller.Category{}
	categories, err := categoryController.GetAllCategories(r)

	if err != nil {
		log.Println(err)
	} else {
		jsonRes,_ := json.Marshal(categories)
		fmt.Fprint(w, string(jsonRes))
	}
}

func CategoryGetTree(w http.ResponseWriter, r *http.Request) {
	categoryController := controller.Category{}
	categories, err := categoryController.GetCategoryTree(r, 0)

	if err != nil {
		log.Println(err)
	} else {
		jsonRes,_ := json.Marshal(categories)
		fmt.Fprint(w, string(jsonRes))
	}
}

func ProductCreateHandler(w http.ResponseWriter, r *http.Request) {
	productController := controller.Product{}

	//verify user is admin
	adminAuth := auth.Auth{}
	status := adminAuth.VerifyAdmin(r)

	if status == true {
		err := productController.CreateProduct(r)

		if err != nil {
			log.Println(err)
		} else {
			fmt.Fprint(w, true)
		}
	}
}

func ProductGetHandler(w http.ResponseWriter, r *http.Request) {
	//product id
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		log.Println(err)
	}

	productController := controller.Product{}
	product, err := productController.GetProduct(r, id)

	if err != nil {
		log.Println(err)
	} else {
		jsonRes,_ := json.Marshal(product)
		fmt.Fprint(w, string(jsonRes))
	}
}

/*
func testHandler(w http.ResponseWriter, r *http.Request) {
	userModel := model.User {
		Email: "frog@gmail.com",
		Password: "test3",
		Role: 1,
	}

	err := userModel.AddUser(w,r)

	if err != nil {
		//hanle error
		log.Println(err)
	}
}
*/
