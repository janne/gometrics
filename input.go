package main

import (
	"net/http"
)

func InputHandler(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Invalid request", 400)
}
