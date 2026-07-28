package main

import (
	"fmt"
	"net/http"
)

func HelloFromGitOpsLab(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World From GitOps Labbbbbb")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloFromGitOpsLab)

	http.ListenAndServe(":8080", mux)
}
