package blogsroute

import (
	blogsapi "github.com/Rajendro1/AntinoGoLang/blogs/api"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	post := r.Group("/post")
	{
		post.GET("", blogsapi.GetPost)
		post.POST("", blogsapi.CreatePost)
		post.PUT("", blogsapi.UpdatePost)
		post.DELETE("", blogsapi.DeletePost)
	}
	r.GET("/posts", blogsapi.GetPosts)
}
