package routes

import (
	"github.com/gorilla/mux"
	controller "github.com/shutt90/goggins-spam/controllers"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/getquote", controller.GetQuote)

	return router
}
