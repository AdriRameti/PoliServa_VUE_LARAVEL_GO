package sports

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"poliserva/Config"
)

func GetAllSports(c *gin.Context) {
	var sports []SportModel

	if err := Config.DB.Find(&sports).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	c.JSON(http.StatusOK, sports)
}