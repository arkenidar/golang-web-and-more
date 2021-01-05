package main

import (
	"log"
	"net/http"
    "fmt"
)

func main() {
	// curl http://localhost:3000/static/1.txt
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", index)
    
	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<a href='/static/1.txt'>static file link</a>")
}
