package helpers

import (
	"encoding/json"
	"go_rest_api/src/models"
	"net/http"
	"strings"
)

func DecodeBody(req *http.Request) (models.Content, bool) {
	var content models.Content
	err := json.NewDecoder(req.Body).Decode(&content)
	if err != nil {
		return models.Content{}, false
	}

	return content, true
}

func IsValidDescription(description string) bool {
	desc := strings.TrimSpace(description)
	if len(desc) == 0 {
		return false
	}

	return true
}
