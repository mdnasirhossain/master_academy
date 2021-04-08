package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/Nasir", Nasir)
	http.HandleFunc("/About", About)
	http.HandleFunc("/Contact", Contact)
	http.HandleFunc("/Blog", Blog)
	http.ListenAndServe(":8888", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `My fast golang Website`)
}

func Nasir(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Welcome to Nasir page`)
}

func About(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Welcome to About page`)

}
func Contact(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Welcome to Contact page`)
}
func Blog(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Welcome to Blog page`)
}
