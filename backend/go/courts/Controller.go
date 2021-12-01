package courts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllCourts(c *gin.Context) {
	GetAllCourtsDB(c)
}

func CreateCourt(c *gin.Context) {

	courtModelValidator := NewModelValidator()

	if err := courtModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	CreateCourtDB(&courtModelValidator.courtModel, c)

	// c.JSON(http.StatusOK, courtModelValidator.courtModel)

	// var court CourtModel

	// c.BindJSON(&court)

	// if err := ValidateCourt(&court); err == true {
	// 	CreateCourtDB(&court, c)
	// } else {
	// 	c.JSON(http.StatusOK, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	// }

}