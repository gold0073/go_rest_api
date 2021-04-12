package main

import (
	"fmt"
	"net/http"

	"goWeb/go-rest-api/src/api"

	"github.com/gorilla/mux"
)

func main() {
	//db := database.GetConnection()
	//fmt.Printf("Hello World")

	var port string = "15432"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/ms-addContent", api.CreateContent).Methods("POST")

	//Search list
	apiRouter.HandleFunc("/ms-contentlist", api.GetContent).Methods("GET")
	//Search Id
	apiRouter.HandleFunc("/ms-contentlist/{content_id}", api.GetContentDetail).Methods("GET")

	fmt.Printf("Server runing at port %s", port)
	http.ListenAndServe(":"+port, router)
}
