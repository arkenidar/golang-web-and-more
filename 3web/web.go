package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	check(http.ListenAndServe(":81", nil))
}

// http://localhost:81/?s=xyz
func index(w http.ResponseWriter, req *http.Request) {
	check(templ.Execute(w, struct{ Value string }{Value: req.FormValue("s")}))
}

var templ = template.Must(template.New("test").Parse("template...{{.Value}}"))

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
