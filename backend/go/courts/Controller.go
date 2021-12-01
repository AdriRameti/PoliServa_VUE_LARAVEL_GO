package courts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllCourts(c *gin.Context) {
	courts := GetAllCourtsDB(c)

	serializer := CourtsSerializer{c, courts}

	c.JSON(http.StatusOK, serializer.Response())
}

func CreateCourt(c *gin.Context) {

	courtModelValidator := NewModelValidator()

	if err := courtModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	if errC := CreateCourtDB(&courtModelValidator.courtModel, c); errC != nil {
		c.JSON(http.StatusUnprocessableEntity, errC)
		return
	}

	serializer := CourtSerializer{c, courtModelValidator.courtModel}

	c.JSON(http.StatusCreated, serializer.Response())

}