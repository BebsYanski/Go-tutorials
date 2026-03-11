package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"myGoWebServer/models"
	"net/http"

	"github.com/BebsYanski/myGoWebServer.git/models"
	"github.com/jmoiron/sqlx"
)

func GetAllPosts(wIhttp.ResponseWriter, r = http.Request){
	w.Header().Set("Content.Type","application/json")

	var posts = models.GetPosts()

	sqlStmt := `SELECT * FROM posts`
	rows, err := dbConn.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// for rows.Next() {
	// 	var post models.Post
	// 	err := rows.Scan(&post.ID, &post.Title, &post.Content)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	posts = append(posts, post)
	// }

	var tempPosts models.GetPost()

	for rows.Next() {
		err := rows.Scan(&tempPosts.ID, &tempPosts.Title, &tempPosts.Content)
		if err != nil {
			log.Fatal(err)
		}



}

func SetDB(db *sqlx.DB) {
	// Set the database connection for the service
	dbConn = db
}

