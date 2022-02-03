package main

import (
	"file-information/api"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8081", srv)
}
