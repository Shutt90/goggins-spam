package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/shutt90/goggins-spam/routes"
)

func main() {
	godotenv.Load()
	router := routes.Init()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
