package helpers

import (
	"encoding/json"
	"go_rest_api/src/models"

	"net/http"
	"strings"
)

func DecodeBody(req *http.Request) (models.Todo, bool) {
	var todo models.Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		return models.Todo{}, false
	}

	return todo, true
}

func IsValidDescription(description string) bool {
	desc := strings.TrimSpace(description)
	if len(desc) == 0 {
		return false
	}

	return true
}
