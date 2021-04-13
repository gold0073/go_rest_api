package models

import (
	"go_rest_api/src/database"
)

type Content struct {
	Content_id int    `json:"content_id"`
	User_id    int    `json:"user_id"`
	Title      string `json:"title"`
	Context    string `json:"context"`
	Created_at string `json:"created_at"`
}

func addContent(title string, context string) (Content, bool) {
	db := database.GetConnection()

	var content_id int

	var pg_query = "INSERT INTO CONTENT (user_id,title,context,created_at,update_at) values ($1, $2, $3,now(),null)"

	db.QueryRow(pg_query, 1, title, context).Scan(&content_id)

	if content_id == 0 {
		return Content{}, false
	}

	return Content{content_id, 1, "", "", ""}, true
}
