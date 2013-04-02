package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	api_key := os.Getenv("API_KEY")
	port := os.Getenv("PORT")

	if api_key == "" || port == "" {
	  panic("API_KEY and PORT environment variables has to be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", InputHandler).Methods("POST").Queries("api_key", api_key)
	router.HandleFunc("/", ReadmeHandler)
	http.Handle("/", router)

	fmt.Println("Listening at port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
