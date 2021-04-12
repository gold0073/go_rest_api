package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"goWeb/go-rest-api/src/models"

	"goWeb/go-rest-api/src/helpers"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool             `json:"sucess"`
	Data    []models.Content `json:"data"`
	Errors  []string         `json:"errors"`
}

func CreateContent(w http.ResponseWriter, req *http.Request) {

	bodyContent, sucess := helpers.DecodeBody(req)

	if sucess != true {
		http.Error(w, " could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyContent.Content = strings.TrimSpace(bodyContent.Content)

	if !helpers.IsValidDescription(bodyContent.content) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid content")

		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return

	}

	content, success := models.addContent(bodyContent.title, bodyContent.content)

	if success != true {
		data.Errors = append(data.Errors, "could not create content")
	}

	data.Success = true
	data.Data = append(data.Data, content)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return

}

func GetContent(w http.ResponseWriter, req *http.Request) {
	var todos []models.Todo = models.contentlist()

	var data = Data{true, todos, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetContentDetail(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["content_id"]

	var data Data

	var todo models.Content
	var success bool
	todo, success = models.contentlist_detail(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "todo not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
