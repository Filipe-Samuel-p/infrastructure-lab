package main

import (
	"net/http"

	"github.com/Filipe-Samuel-p/infrastructure-lab/user"
)

func main() {

	service := &user.UserService{}
	h := &user.UserHandler{Service: service}

	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HelloUser)

	http.ListenAndServe(":8080", mux)
}
