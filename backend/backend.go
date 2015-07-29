package backend

import (
	"fmt"
	"net/http"
    //"log"
    //"appengine"
)

func init() {
    http.HandleFunc("/api/mogo/", handlerMogo)
	http.HandleFunc("/api/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
    //log.Println("Move maker supreme")
    fmt.Fprint(w, "Hekkraaae")
}

func handlerMogo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Mogo crew")
}
