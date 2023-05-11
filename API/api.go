package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type welcome string

//user struct

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

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

func checkUser(w http.ResponseWriter, r *http.Request) {
	var user User
	dbPass := "aliban"
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal("error decoding ito structure")
	}
	if user.Password == dbPass {
		fmt.Println("Successfully logged in")
	}

	fmt.Fprintf(w, "Response: %v", user)
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

	router.HandleFunc("/user", checkUser)

	//Server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
