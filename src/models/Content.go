package models

import (
	"time"

	"goWeb/go-rest-api/src/database"
)

type Content struct {
	content_id	int			`json:"content_id"`
	user_id		int     	`json:"user_id"`
	title  		string  	`json:"title"`
	context     string  	`json:"context"`
	CreatedAt   time.Time  	`json:"CreatedAt"`
}

func addContent(title string ,context string) (Content,bool)
{
	db:= database.GetConnection()

	var sec_id int 
	var pg_query = 
    ` INSERT INTO CONTENT (user_id,title,context,created_at,update_at) values ($1, $2, $3,now(),null)`;
	
	db.QueryRow(pg_query,{1 ,title , context}).Scan(&sec_id)

	if sec_id ==0 {
		return Content{},false
	}

	return Content{sec_id,""} true
}

func contentlist() (Content,bool){
	db:= database.GetConnection()

	const pg_query = 
    ` SELECT * FROM CONTENT CT 
        INNER JOIN USERS U on U.user_id = CT.user_id 
        ORDER BY CT.created_at DESC
    `;

	row:= db.QueryRow(pg_query)

	var content_id int

	err:= row.Scan(&content_id)

	if err != nil{
		return Content{}, false
	}

	return Content{
		content_id,
		user_id,
		title,
		context,
		CreatedAt,
	} , true
}

func contentlist_detail(search_id) (Content,bool){
	db:= database.GetConnection()

	var pg_query = 
	` SELECT * FROM CONTENT CT 
	INNER JOIN USERS U on U.user_id = CT.user_id 
	WHERE CT.content_id = $1`;

	row:= db.QueryRow(pg_query,search_id)

	var content_id int

	err:= row.Scan(&content_id)

	if err != nil{
		return Content{}, false
	}

	return Content{
		content_id,
		user_id,
		title,
		context,
		CreatedAt,
	} , true
}