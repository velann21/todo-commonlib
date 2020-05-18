package routes

import (
	"github.com/gorilla/mux"
	"github.com/velann21/todo-commonlib/controller"
)

func Initialize(indexRoute *mux.Router) {
	indexRoute.HandleFunc("/signin", controller.RegisterSchema).Methods("POST")
}
