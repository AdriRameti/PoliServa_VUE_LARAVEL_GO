package courts

import (
	"github.com/gin-gonic/gin"
)

func CourtsRouting(router *gin.RouterGroup) {

	router.GET("/", GetAllCourts)
	
}