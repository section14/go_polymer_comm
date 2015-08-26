package backend

import (
	//standard library
	"fmt"
	"net/http"
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
	//"github.com/section14/go_polymer_comm_pkg/model"
)

//setup session
var signString []byte = []byte("oboeMadSauceSupremeGammaTrainSuprippp$%&*%^@@@vsmsoiosvh")

func init() {
	r := mux.NewRouter();

	r.HandleFunc("/api/", handler)
	r.HandleFunc("/api/user/{name}/", UserGetHandler).Methods("GET")
	r.HandleFunc("/api/user/", UserCreateHandler).Methods("POST")
	r.HandleFunc("/api/user/email/", UserGetEmailHandler).Methods("POST")
	r.HandleFunc("/api/user/login/", UserLoginHandler).Methods("POST")

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
    fmt.Fprint(w, "Root level, nothing here")
}

/*

Everything below needs error handling

*/

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["name"]

	userController := controller.User{}
	userController.TestHit()

	log.Println(user)
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	userController := controller.User{}
	_, err := userController.CreateUser(w,r)

	if err != nil {
		//handle err
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
	userController := controller.User{}
	userStatus, err := userController.Login(w,r)

	if err != nil || userStatus == false {
		//handle err
	}

	//issue jwt token
	token := getToken()

	//encode json return here?
	//Need to pass the token to user somehow

	fmt.Fprint(w, token)
}

func getToken() string {
	//jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userId"] = "some Id string"
	token.Claims["userRole"] = "1"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(signString)

	if err != nil {
		//handle error
	}

	//return token
	return tokenString
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
