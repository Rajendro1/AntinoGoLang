package blogsdb

import (
	"database/sql"
	"log"

	blogsStruct "github.com/Rajendro1/AntinoGoLang/blogs"
	mysqldb "github.com/Rajendro1/AntinoGoLang/db/mysql"
)

func GetPostByIdDFromDB(id int) (blogsStruct.Post, error) {
	var post blogsStruct.Post
	if rowErr := mysqldb.DB.QueryRow("SELECT id, title, content, author, created_at, updated_at FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.UpdatedAt); rowErr != nil {
		if rowErr == sql.ErrNoRows {
			return post, nil
		}
		log.Println("GetPostByIdDFromDB rowErr: ", rowErr.Error())
		return post, rowErr
	}
	return post, nil
}
func CreatePostToDB(title, content, author string) (int, error) {
	res, resErr := mysqldb.DB.Exec("INSERT INTO posts (title, content, author) VALUES (?, ?, ?)",
		title, content, author)
	if resErr != nil {
		log.Println("CreatePostToDB resErr: ", resErr.Error())
		return 0, resErr
	}

	lastId, lastIdErr := res.LastInsertId()
	if lastIdErr != nil {
		log.Println("CreatePostToDB lastIdErr: ", lastIdErr.Error())
		return 0, lastIdErr
	}

	return int(lastId), nil
}
func GetPostsFromDB() ([]blogsStruct.Post, error) {
	var post blogsStruct.Post
	var posts []blogsStruct.Post

	rows, rowsErr := mysqldb.DB.Query("SELECT * FROM posts")
	if rowsErr != nil {
		log.Println("GetPostsFromDB rowsErr: ", rowsErr.Error())
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		scanErr := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.UpdatedAt)
		if scanErr != nil {
			log.Println("GetPostsFromDB scanErr: ", scanErr.Error())
			return nil, scanErr
		}
		posts = append(posts, post)
	}

	return posts, nil
}
func UpdatePost(id int, title, content, author string) (bool, error) {
	if _, err := mysqldb.DB.Exec("UPDATE posts SET title = ?, content = ?, author = ? WHERE id = ?",
		title, content, author, id); err != nil {
		log.Println("UpdatePost err: ", err.Error())
		return false, err
	}
	return true, nil
}

func DeletePost(id int) (bool, error) {
	if _, err := mysqldb.DB.Exec("DELETE FROM posts WHERE id = ?", id); err != nil {
		log.Println("DeletePost err: ", err.Error())
		return false, err
	}
	return true, nil
}
func ValidatePostId(id int) bool {
	var isHave bool
	if err := mysqldb.DB.QueryRow("SELECT TRUE FROM posts WHERE id = ?", id).Scan(&isHave); err != nil {
		log.Println("ValidatePostId err: ", err.Error())
		return isHave
	}
	return isHave
}
