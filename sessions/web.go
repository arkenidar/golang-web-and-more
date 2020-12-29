// shell: go run web.go
// shell: go get "github.com/gorilla/sessions"

package main

import (
	"fmt"      // text formatter
	"log"      // debug logger
	"net/http" // HTTP server

	"github.com/gorilla/sessions" // HTTP sessions
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/del", delete)
	log.Print("running http.ListenAndServe :1050")
	log.Print(http.ListenAndServe(":1050", nil))
}

var store = sessions.NewCookieStore([]byte("key"))

func delete(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["counter"] = nil
	session.Save(r, w)
}
func index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	var newCounter = 1
	if session.Values["counter"] != nil {
		var counter, _ = session.Values["counter"].(int)
		newCounter = counter + 1
	}
	session.Values["counter"] = newCounter
	session.Save(r, w)
	fmt.Fprintln(w, "counter (reload page to count):", session.Values["counter"])
}
