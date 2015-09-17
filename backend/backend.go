package backend

import (
	//standard library
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"log"
	//"appengine"

	//third party
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	//"github.com/gorilla/securecookie"
	"github.com/dgrijalva/jwt-go"

	//my imports
	"github.com/section14/go_polymer_comm_pkg/controller"
	"github.com/section14/go_polymer_comm_pkg/user-auth"
)

//jwt token signing string
var signString []byte = []byte("oboeMadSauceSupremeGammaTrainSuprippp$%&*%^@@@vsmsoiosvh")

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

	//product
	r.HandleFunc("/api/category", CategoryCreateHandler).Methods("POST")

	//r.HandleFunc("/api/test/", testHandler)

	/*
	r.HandleFunc("/api/bro/", broHandler)
	r.HandleFunc("/api/flo/", floHandler)
	r.HandleFunc("/api/upt/", updateHandler)
	r.HandleFunc("/api/del/", deleteHandler)
	*/

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
		jwt := generateToken(user.Id, user.Role)
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
		jwt := generateToken(user.Id, user.Role)
		jsonJwt = &JwtToken {
			Token: jwt,
			Status: true,
		}
	}

	token,_ := json.Marshal(jsonJwt)
	fmt.Fprint(w, string(token))
}

/*

Category

*/

func CategoryCreateHandler(w http.ResponseWriter, r *http.Request) {
	categoryController := controller.Category{}
	_, err := categoryController.CreateCategory(w,r)

	if err != nil {
		log.Println(err)
	}
}

/*

Token

*/

func generateToken(Id int64, Role int) string {
	//jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userId"] = Id
	token.Claims["userRole"] = Role
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(signString)

	if err != nil {
		//handle error
		return ""
	}

	//return token
	return tokenString
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	myToken :=r.Header.Get("User-Token")

	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {

		//verify signing type
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return signString, nil
	})

	//return error or decoded token
	if err == nil && token.Valid {
		return token, nil
	} else {
		return nil, err
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

/*

func broHandler(w http.ResponseWriter, r *http.Request) {
	userModel := model.User{}

	res, err := userModel.GetAllUsers(w,r)

	if err != nil {
		//hanle error
	}

	fmt.Fprint(w, res)
}

func floHandler(w http.ResponseWriter, r *http.Request) {
	userModel := model.User{}

	res, err := userModel.GetUser(w,r,6192449487634432)

	if err != nil {
		//hanle error
	}

	fmt.Fprint(w, res)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	userModel := model.User{}

	err := userModel.UpdateEmail(w,r,6192449487634432,"oboes-never-gojos")

	if err != nil {
		//hanle error
		log.Println(err)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	userModel := model.User{}

	err := userModel.DeleteUser(w,r,5838406743490560)

	if err != nil {
		//handle error
	}
}

func adminHander(w http.ResponseWriter, r *http.Request) {
	//need to make sure user is logged in before they can get here
	//just return a 500, omit error
}
*/
