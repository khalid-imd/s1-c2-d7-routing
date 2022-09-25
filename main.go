package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/hi", helloworld).Methods("GET")

	route.HandleFunc("/home", home).Methods("GET")

	route.HandleFunc("/project", project).Methods("GET")

	route.HandleFunc("/submit", submit).Methods("POST")

	route.HandleFunc("/contact", contact).Methods("GET")

	fmt.Println("server is Running")
	http.ListenAndServe("localhost:8000", route)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset-utf8")

	var tmplt, err = template.ParseFiles("pages/home.html")

	if err != nil {
		w.Write([]byte("file doesn't exist: " + err.Error()))
		return
	}

	// w.Write([]byte("home"))
	//w.WriteHeader(http.StatusAccepted)
	tmplt.Execute(w, "")
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset-utf8")

	var tmplt, err = template.ParseFiles("pages/project.html")

	if err != nil {
		w.Write([]byte("file doesn't exist: " + err.Error()))
		return
	}

	// w.Write([]byte("home"))
	//w.WriteHeader(http.StatusAccepted)
	tmplt.Execute(w, "")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset-utf8")

	var tmplt, err = template.ParseFiles("pages/contact.html")

	if err != nil {
		w.Write([]byte("file doesn't exist: " + err.Error()))
		return
	}

	// w.Write([]byte("home"))
	//w.WriteHeader(http.StatusAccepted)
	tmplt.Execute(w, "")
}

func submit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("title: " + r.PostForm.Get("addTitle"))
	fmt.Println("start date: " + r.PostForm.Get("addStartDate"))
	fmt.Println("end date: " + r.PostForm.Get("addEndDate"))
	fmt.Println("description: " + r.PostForm.Get("addDescription"))
	fmt.Println(r.PostForm.Get("addNode"))
	fmt.Println(r.PostForm.Get("addReact"))
	fmt.Println(r.PostForm.Get("addNext"))
	fmt.Println(r.PostForm.Get("addTypeScript"))

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)

}
