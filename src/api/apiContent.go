package api

import (
	"encoding/json"
	"fmt"
	"go_rest_api/src/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Content struct {
	Success bool             `json:"success"`
	Data    []models.Content `json:"data"`
	Errors  []string         `json:"errors"`
}

func CreateContent(w http.ResponseWriter, req *http.Request) {
	var contents models.Content

	err := json.NewDecoder(req.Body).Decode(&contents)

	if err != nil {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Content = Content{Errors: make([]string, 0)}
	contents.Title = strings.TrimSpace(contents.Title)
	contents.Context = strings.TrimSpace(contents.Context)

	content, success := models.AddContent(contents.Title, contents.Context)

	if success != true {
		data.Errors = append(data.Errors, "could not create content")
	}

	data.Success = success
	data.Data = append(data.Data, content)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

func GetContents(w http.ResponseWriter, req *http.Request) {
	var contents []models.Content = models.GetContentlist()

	var data = Content{true, contents, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetContent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Content

	//var content models.Content
	//var success bool
	content, success := models.GetContentDetail(id)

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

func DeleteContent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Content = Content{Errors: make([]string, 0)}

	content, success := models.RemoveContent(id)
	if success != true {
		data.Errors = append(data.Errors, "could not delete content")
	}

	data.Success = true
	data.Data = append(data.Data, content)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateContent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	content_id := vars["id"]

	var contents models.Content
	err := json.NewDecoder(req.Body).Decode(&contents)
	if err != nil {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Content = Content{Errors: make([]string, 0)}
	contents.Title = strings.TrimSpace(contents.Title)
	contents.Context = strings.TrimSpace(contents.Context)

	fmt.Println("content_id ===>" + content_id)
	fmt.Println(contents.Title)
	fmt.Println(contents.Context)

	content, success := models.EditContent(content_id, contents.Title, contents.Context)
	if success != true {
		data.Errors = append(data.Errors, "could not update todo")
	}

	data.Success = success
	data.Data = append(data.Data, content)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}
