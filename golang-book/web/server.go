package main

import (
	"fmt"
	"net/http"
)

type welcome string

// Handler Interface
func (wc welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to AppsCode Future server!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Please!")
}

func getJson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		w.Write([]byte(`"message: GET is called"`))
	case "POST":
		w.Write([]byte(`"message: POST is called"`))
	}

}

func main() {

	//Router

	router := http.NewServeMux()

	//Handler
	var wc welcome
	router.Handle("/", wc)

	//Handle funcs

	router.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Logout Page")
	})

	router.HandleFunc("/login", login)

	router.HandleFunc("/json", getJson)

	//Server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
