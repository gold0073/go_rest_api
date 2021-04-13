package main

import (
	"fmt"
	"net/http"

	"go_rest_api/src/api"

	"github.com/gorilla/mux"
)

func main() {
	var port string = "8080"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()

	//todos Router
	apiRouter.HandleFunc("/todos", api.GetTodos).Methods("GET")
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")

	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", api.DeleteTodo).Methods("DELETE")
	apiRouter.HandleFunc("/todos/{id}", api.UpdateTodo).Methods("PUT")

	//Content Router
	apiRouter.HandleFunc("/ms-addContent", api.CreateContent).Methods("POST")

	/*
		apiRouter.HandleFunc("/ms-contentlist", api.GetContents).Methods("GET")


		apiRouter.HandleFunc("/ms-contentlist/{id}", api.GetContent).Methods("GET")
		apiRouter.HandleFunc("/ms-deleteContent/{id}", api.DeleteContent).Methods("DELETE")
		apiRouter.HandleFunc("/ms-updateContent/{id}", api.UpdateContent).Methods("PUT")
	*/

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
