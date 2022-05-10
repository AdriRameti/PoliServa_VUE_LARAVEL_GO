package courts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poliserva/Common"
)

func GetAllCourts(c *gin.Context) {
	courts := GetAllCourtsDB(c)

	serializer := CourtsSerializer{c, courts}

	c.JSON(http.StatusOK, serializer.Response())
}
func GetDateReservations(c *gin.Context){
	println("Entra en la función")
	Date:= c.Query("date")
	hini:= c.Query("hini")
	hfin:= c.Query("hfin")
	slug:= c.Query("slug")
	if len(Date)<1||  len(hini)<1 || len(hfin)<1 || len(slug)<1 {
		return
	}
	validDate:= common.ValidateDate(Date)
	validHini:= common.ValidateHour(hini)
	validHfin:= common.ValidateHour(hfin)
	if validDate == false || validHini == false || validHfin == false{
		c.AbortWithStatus(http.StatusNotFound);
	}
	MyArray:= []string {Date,hini,hfin,slug}
	courts:=GetDateOneReservation(MyArray,c)
	serializer:= CourtsSerializer{c, courts}
	c.JSON(http.StatusCreated, serializer.Response())
}
func GetOneCourt(c *gin.Context) {
	var id string

	id = c.Query("id")

	court := GetOneCourtDB(id, c)

	serializer := CourtSerializer{c, court}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetCourtsBySport(c *gin.Context) {
	var slug string

	slug = c.Query("slug")

	courts := GetCourtsBySportDB(slug, c)

	serializer := CourtsSerializer{c, courts}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetCourtsCarousel(c *gin.Context) {
	
	courts := GetCourtsCarouselDB(c)

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

func UpdateCourt(c *gin.Context) {

	var id string
	var court CourtModel

	id = c.Query("id")
	
	courtModelValidator := NewModelValidator()

	if err := courtModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	court, errU := UpdateCourtDB(id, &courtModelValidator.courtModel, c)

	if errU != nil {
		c.JSON(http.StatusUnprocessableEntity, errU)
		return
	}

	serializer := CourtSerializer{c, court}

	c.JSON(http.StatusOK, serializer.Response())

}

func DeleteCourt(c *gin.Context) {
	var id string

	id = c.Query("id")

	DeleteCourtDB(id, c)
}