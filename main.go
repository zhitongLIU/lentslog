package main

import (
	"github.com/zhitongLIU/lentslog/routes"
	"net/http"
	"os"
)

func main() {
	r := routes.NewRouter()

	http.Handle("/", r)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
