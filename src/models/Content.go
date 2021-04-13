package models

import (
	"fmt"
	"go_rest_api/src/database"
	"log"
)

type Content struct {
	Content_id int    `json:"content_id"`
	User_id    int    `json:"user_id"`
	User_name  string `json:"user_name"`
	Title      string `json:"title"`
	Context    string `json:"context"`
	Created_at string `json:"created_at"`
}

func AddContent(title string, context string) (Content, bool) {
	db := database.GetConnection()

	var content_id int

	var pg_query = "INSERT INTO CONTENT (user_id,title,context,created_at,update_at) values ($1, $2, $3,now(),null) RETURNING content_id,title,context"

	db.QueryRow(pg_query, 1, title, context).Scan(&content_id, &title, &context)

	if content_id == 0 {
		return Content{}, false
	}

	return Content{content_id, 1, "", title, context, ""}, true
}

func GetContentlist() []Content {
	db := database.GetConnection()
	//전체리스트
	var pg_query = ` SELECT CT.content_id,U.user_id,U.user_name,CT.title,CT.context,CT.created_at FROM CONTENT CT 
		INNER JOIN USERS U on U.user_id = CT.user_id 
		ORDER BY CT.created_at DESC `

	fmt.Printf(pg_query)

	rows, err := db.Query(pg_query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var contents []Content
	for rows.Next() {
		t := Content{}

		var content_id int
		var user_id int
		var user_name string
		var title string
		var context string
		var created_at string

		err := rows.Scan(&content_id, &user_id, &user_name, &title, &context, &created_at)
		if err != nil {
			log.Fatal(err)
		}

		t.Content_id = content_id
		t.User_id = user_id
		t.User_name = user_name
		t.Title = title
		t.Context = context
		t.Created_at = created_at

		contents = append(contents, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return contents
}

func GetContentDetail(id string) (Content, bool) {
	db := database.GetConnection()
	var pg_query = ` SELECT CT.content_id,U.user_id,U.user_name,CT.title,CT.context,CT.created_at 
	FROM CONTENT CT INNER JOIN USERS U on U.user_id = CT.user_id WHERE CT.content_id = $1`

	fmt.Println(pg_query)
	fmt.Println("id===>" + id)

	row := db.QueryRow(pg_query, id)

	var content_id int
	var user_id int
	var user_name string
	var title string
	var context string
	var created_at string

	err := row.Scan(&content_id, &user_id, &user_name, &title, &context, &created_at)
	if err != nil {
		return Content{}, false
	}

	return Content{content_id, user_id, user_name, title, context, created_at}, true
}

func RemoveContent(id string) (Content, bool) {
	db := database.GetConnection()

	var content_id int
	db.QueryRow("DELETE FROM content WHERE content_id = $1 RETURNING content_id", id).Scan(&content_id)

	if content_id == 0 {
		return Content{}, false
	}

	return Content{content_id, 0, "", "", "", ""}, true
}

func EditContent(id string, title string, context string) (Content, bool) {
	db := database.GetConnection()

	var content_id int

	var pg_query = " UPDATE CONTENT SET title = $1, context = $2 WHERE content_id = $3 RETURNING content_id "

	fmt.Println(pg_query)
	fmt.Println(title)

	db.QueryRow(pg_query, title, context, id).Scan(&content_id)
	if content_id == 0 {
		return Content{}, false
	}

	return Content{content_id, 0, "", title, context, ""}, true
}
