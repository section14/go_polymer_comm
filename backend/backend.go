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

	//my imports
	"github.com/section14/go_polymer_comm_pkg/model"
)

//setup session
var store = sessions.NewCookieStore([]byte("auto-intro-amigo-status"))

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

    fmt.Fprint(w, "Root level, nothing here")
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
