package blogsapi

import (
	"net/http"
	"strconv"

	blogsdb "github.com/Rajendro1/AntinoGoLang/blogs/db"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
    posts, postsErr := blogsdb.GetPostsFromDB()
    if postsErr != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "Sorry! We encountered a problem while fetching posts",
            "error":   postsErr.Error(),
        })
        return
    }
    if len(posts) == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "Sorry! We don't have any posts",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Posts retrieved successfully",
        "data":    posts,
    })
}


func GetPost(c *gin.Context) {
	id := c.Request.FormValue("id")
	intId, err := strconv.Atoi(id)

	if err != nil || intId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please provide a valid input",
		})
		return
	}

	validateId := blogsdb.ValidatePostId(intId)
	if !validateId {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Sorry! This ID is not present in our database",
		})
		return
	}

	post, postErr := blogsdb.GetPostByIdDFromDB(intId)
	if postErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! We encountered a problem while fetching your post",
			"error":   postErr.Error(),
		})
		return
	}

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Sorry! This ID does not have any posts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Post retrieved successfully",
		"data":    post,
	})
}

func CreatePost(c *gin.Context) {
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	author := c.Request.FormValue("author")

	createPostID, createPostErr := blogsdb.CreatePostToDB(title, content, author)
	if createPostErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! we could add your post",
			"error":   createPostErr.Error(),
		})
		return
	}
	getPost, getPostErr := blogsdb.GetPostByIdDFromDB(createPostID)
	if getPostErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! we have problem in our server while fetching your post",
			"error":   getPostErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "post are getting successfully",
		"data":    getPost,
	})

}

func UpdatePost(c *gin.Context) {
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	author := c.Request.FormValue("author")
	id := c.Request.FormValue("id")
	intId, _ := strconv.Atoi(id)
	validateId := blogsdb.ValidatePostId(intId)
	if !validateId {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.NotFound,
			"message": "Sorry! this id is not present in our database",
		})
		return
	}
	updatePost, updatePostErr := blogsdb.UpdatePost(intId, title, content, author)
	if updatePostErr != nil || !updatePost {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! we have problem in our server while updating this post",
			"error":   updatePostErr.Error(),
		})
		return
	}
	if updatePost {
		getPost, getPostErr := blogsdb.GetPostByIdDFromDB(intId)
		if getPostErr != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Sorry! we have problem in our server while fetching your post",
				"error":   getPostErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "post are updated successfully",
			"data":    getPost,
		})
	}
}

func DeletePost(c *gin.Context) {
	id := c.Request.FormValue("id")

	intId, err := strconv.Atoi(id)

	if err != nil || intId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please provide a valid input",
		})
		return
	}

	validateId := blogsdb.ValidatePostId(intId)
	if !validateId {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Sorry! This ID is not present in our database",
		})
		return
	}

	deletePost, deletePostErr := blogsdb.DeletePost(intId)
	if deletePostErr != nil || !deletePost {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! We encountered a problem while deleting this post",
			"error":   deletePostErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Post deleted successfully",
	})
}
