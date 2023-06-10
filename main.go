package main

import (
	"net/http"

	"github.com/shutt90/goggins-spam/routes"
)

func main() {
	router := routes.Init()

	http.ListenAndServe(":8080", router)
}
