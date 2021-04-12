package api

import (
	"encoding/json"
	"go_rest_api/src/helpers"
	"go_rest_api/src/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool             `json:"sucess"`
	Data    []models.Content `json:"data"`
	Errors  []string         `json:"errors"`
}

func CreateContent(w http.ResponseWriter, req *http.Request) {

	bodyContent, success := helpers.DecodeBody(req)

	if success != true {
		http.Error(w, " could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyContent.Description = strings.TrimSpace(bodyContent.Description)

	if !helpers.IsValidDescription(bodyContent.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid content")

		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return

	}

	content, success := models.(bodyContent.Description)

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
	var contents []models.Content = models.contentlist()

	var data = Data{true, contents, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetContentDetail(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["content_id"]

	var data Data

	var content models.Content
	var success bool
	content, success = models.contentlist_detail(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "content not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, content)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
