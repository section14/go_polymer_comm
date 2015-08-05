package backend

import (
	//standard library
	"fmt"
	"net/http"
	"log"
	//"appengine"

	//third party
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	//"github.com/gorilla/securecookie"
	"github.com/dgrijalva/jwt-go"

	//my imports
	"github.com/section14/go_polymer_comm_pkg/model"
)

//setup session
var store = sessions.NewCookieStore([]byte("auto-intro-amigo-status"))
var signString []byte = []byte("oboeMadSauceSupremeGammaTrainSuprippp$%&*%^@@@vsmsoiosvh")

func init() {
	r := mux.NewRouter();

	r.HandleFunc("/api/", handler)
	r.HandleFunc("/api/test/", testHandler)
	r.HandleFunc("/api/bro/", broHandler)
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "app-session")
	session.Values["rude boy"] = "on the scene run bwoy select"
	session.Save(r, w)

	//jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userType"] = "regular user"
	token.Claims["name"] = "Don Draper"
	//token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(signString)

	if err != nil {
		//handle error
	}

    fmt.Fprint(w, "Root level, nothing here")
	fmt.Fprintf(w, "{\"token\": %s}", tokenString)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "app-session")
	rudeBoy := session.Values["rude boy"]

	userModel := model.User {
		Email: "test2",
		Password: "test2",
		Role: 1,
	}

	err := userModel.AddUser(w,r)

	if err != nil {
		//hanle error
		log.Println(err)
	}

	log.Println(rudeBoy)
}

func broHandler(w http.ResponseWriter, r *http.Request) {
	userModel := model.User{}

	res, err := userModel.GetAllUsers(w,r)

	if err != nil {
		//hanle error
	}

	fmt.Fprint(w, res)
}

func adminHander(w http.ResponseWriter, r *http.Request) {
	//need to make sure user is logged in before they can get here
	//just return a 500, omit error
}
