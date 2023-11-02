package route

import (
	blogsroute "github.com/Rajendro1/AntinoGoLang/blogs/route"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	blogsroute.ApplyRoutes(r)
	r.Run(":8080")
}
