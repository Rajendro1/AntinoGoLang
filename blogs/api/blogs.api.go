package blogsapi

import (
	"net/http"

	blogsStruct "github.com/Rajendro1/AntinoGoLang/blogs"
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
	var input blogsStruct.Input

	// Bind JSON body to Post struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	if input.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please provide a valid input",
		})
		return
	}

	validateId := blogsdb.ValidatePostId(input.ID)
	if !validateId {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Sorry! This ID is not present in our database",
		})
		return
	}

	post, postErr := blogsdb.GetPostByIdDFromDB(input.ID)
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
	var newPost blogsStruct.Input
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	createPostID, createPostErr := blogsdb.CreatePostToDB(newPost.Title, newPost.Content, newPost.Author)
	if createPostErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! We couldn't add your post",
			"error":   createPostErr.Error(),
		})
		return
	}

	getPost, getPostErr := blogsdb.GetPostByIdDFromDB(createPostID)
	if getPostErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! We encountered a problem while fetching your post",
			"error":   getPostErr.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Post created successfully",
		"data":    getPost,
	})
}

func UpdatePost(c *gin.Context) {
	var input blogsStruct.Input

	// Bind JSON body to Post struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	// Validate the post ID
	validateId := blogsdb.ValidatePostId(input.ID)
	if !validateId {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": "Sorry! this id is not present in our database",
		})
		return
	}

	// Attempt to update the post
	updatePost, updatePostErr := blogsdb.UpdatePost(input.ID, input.Title, input.Content, input.Author)
	if updatePostErr != nil || !updatePost {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! we have a problem in our server while updating this post",
			"error":   updatePostErr.Error(),
		})
		return
	}

	// Retrieve the updated post
	getPost, getPostErr := blogsdb.GetPostByIdDFromDB(input.ID)
	if getPostErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! we have a problem in our server while fetching your post",
			"error":   getPostErr.Error(),
		})
		return
	}

	// Respond with the updated post
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Post updated successfully",
		"data":    getPost,
	})
}

func DeletePost(c *gin.Context) {
	var input blogsStruct.Input

	// Bind JSON body to Post struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	if input.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please provide a valid input",
		})
		return
	}

	validateId := blogsdb.ValidatePostId(input.ID)
	if !validateId {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Sorry! This ID is not present in our database",
		})
		return
	}

	deletePost, deletePostErr := blogsdb.DeletePost(input.ID)
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
