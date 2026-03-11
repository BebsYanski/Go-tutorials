package models

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetPost() Post {
	var post Post

	return post
}

func GetPosts() []Post {
	var posts []Post
	return posts
}
