package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func main() {

	r := mux.NewRouter().StrictSlash(false)
	mainRoutes := r.PathPrefix("/api/v1/users").Subrouter()

}
