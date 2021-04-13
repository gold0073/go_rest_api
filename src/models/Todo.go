package models

import (
	"go_rest_api/src/database"
	"log"
)

//Todo ...
type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

//Insert ...
func TodoInsert(description string) (Todo, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("INSERT INTO todos(description) VALUES($1) RETURNING id", description).Scan(&todo_id)

	if todo_id == 0 {
		return Todo{}, false
	}

	return Todo{todo_id, ""}, true
}

func TodoGet(id string) (Todo, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	var ID int
	var description string
	err := row.Scan(&ID, &description)
	if err != nil {
		return Todo{}, false
	}

	return Todo{ID, description}, true
}

func TodoDelete(id string) (Todo, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("DELETE FROM todos WHERE id = $1 RETURNING id", id).Scan(&todo_id)
	if todo_id == 0 {
		return Todo{}, false
	}

	return Todo{todo_id, ""}, true
}

func TodoUpdate(id string, description string) (Todo, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("UPDATE todos SET description = $1 WHERE id = $2 RETURNING id", description, id).Scan(&todo_id)
	if todo_id == 0 {
		return Todo{}, false
	}

	return Todo{todo_id, description}, true
}

func TodoGetAll() []Todo {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM todos ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		t := Todo{}

		var ID int
		var description string

		err := rows.Scan(&ID, &description)
		if err != nil {
			log.Fatal(err)
		}

		t.ID = ID
		t.Description = description

		todos = append(todos, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return todos
}
