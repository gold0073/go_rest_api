package api

import (
	"encoding/json"
	"go_rest_api/src/models"
	"net/http"
	"strings"
)

type Content struct {
	Success bool             `json:"success"`
	Data    []models.Content `json:"data"`
	Errors  []string         `json:"errors"`
}

func CreateContent(w http.ResponseWriter, req *http.Request) {
	var content models.Content
	err := json.NewDecoder(req.Body).Decode(&content)

	if err != nil {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Content = Content{Errors: make([]string, 0)}
	content.Title = strings.TrimSpace(content.Title)
	content.Context = strings.TrimSpace(content.Context)

	content, success := models.addContent(content.Title, content.Context)

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

/*
func GetContents(w http.ResponseWriter, req *http.Request) {
	var todos []models.Todo = models.GetAll()

	var data = Content{true, todos, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateContent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todo_id := vars["id"]

	bodyTodo, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Content = Content{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	todo, success := models.Update(todo_id, bodyTodo.Description)
	if success != true {
		data.Errors = append(data.Errors, "could not update todo")
	}

	data.Success = success
	data.Content = append(data.Content, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

func GetContent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Content

	var todo models.Todo
	var success bool
	todo, success = models.Get(id)
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
	data.Content = append(data.Content, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteConntent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Content = Content{Errors: make([]string, 0)}

	todo, success := models.Delete(id)
	if success != true {
		data.Errors = append(data.Errors, "could not delete todo")
	}

	data.Success = success
	data.Content = append(data.Content, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
*/
