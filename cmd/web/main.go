package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "8080", "webserver port")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on " + *port)
	err := http.ListenAndServe(":"+*port, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Sever Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal server Error", 500)
	}
}
