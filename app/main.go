package main

import (
	"infra-lab/user"
	"net/http"
)

func main() {

	service := &user.UserService{}
	h := &user.UserHandler{Service: service}

	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HelloUser)

	http.ListenAndServe(":8080", mux)
}
