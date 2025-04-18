package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Starting front end service on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("error to init the server")
	}

}

func helloWorld(rw http.ResponseWriter, r *http.Request) {

	tmlp, err := template.ParseFiles("template/index.html")

	if err != nil {
		http.Error(rw, "error in load the template", http.StatusInternalServerError)
		return
	}
	data := struct {
		Name string
	}{
		Name: "daniel",
	}

	maybeErro := tmlp.Execute(rw, data)

	if maybeErro != nil {
		http.Error(rw, "error in execute the tamplate", http.StatusInternalServerError)
	}

}
