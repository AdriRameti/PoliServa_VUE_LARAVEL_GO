package courts

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"poliserva/Config"
)

func GetAllCourts(c *gin.Context) {
	var courts []CourtModel

	if err := Config.DB.Find(&courts).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	c.JSON(http.StatusOK, courts)
}