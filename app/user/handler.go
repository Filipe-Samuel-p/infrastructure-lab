package user

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
	Service *UserService
}

func (h UserHandler) HelloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, infra-lab")
}
