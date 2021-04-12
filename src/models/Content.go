package models

import (
	"go_rest_api/src/database"
)

type Content struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

/*
type Content struct {
	content_id int       `json:"content_id"`
	user_id    int       `json:"user_id"`
	title      string    `json:"title"`
	context    string    `json:"context"`
	CreatedAt  time.Time `json:"CreatedAt"`
}
*/

func addContent(Description string) (Content, bool) {
	db := database.GetConnection()

	var sec_id int
	var pg_query = "INSERT INTO CONTENT (user_id,title,context,created_at,update_at) values ($1, $2, $3,now(),null)"

	db.QueryRow(pg_query, 1, Description).Scan(&sec_id)

	if sec_id == 0 {
		return Content{}, false
	}

	return Content{}, true
}

func contentlist() (Content, bool) {
	db := database.GetConnection()

	const pg_query = "SELECT * FROM CONTENT CT INNER JOIN USERS U on U.user_id = CT.user_id ORDER BY CT.created_at DESC"

	row := db.QueryRow(pg_query)

	var content_id int

	err := row.Scan(&content_id)

	if err != nil {
		return Content{}, false
	}

	return Content{}, true
}

func contentlist_detail(id string) (Content, bool) {
	db := database.GetConnection()

	var pg_query = " SELECT * FROM CONTENT CT INNER JOIN USERS U on U.user_id = CT.user_id  WHERE CT.content_id = $1"

	row := db.QueryRow(pg_query, id)

	var content_id int

	err := row.Scan(&content_id)

	if err != nil {
		return Content{}, false
	}

	return Content{}, true
}
