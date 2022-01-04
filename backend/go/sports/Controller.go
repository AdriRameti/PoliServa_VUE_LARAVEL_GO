package sports

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSports(c *gin.Context) {
	var sports []SportModel

	sports = GetAllSportsDB(c)

	serializer := SportsSerializer{c, sports}

	c.JSON(http.StatusOK, serializer.Response())
}
// func GetIdSport(c *gin.Context){
// 	var sport SportModel
// 	var id number
// 	id = c.Query("id")
// 	sport = GetIdSportDB(id, c)
// }
func GetOneSport(c *gin.Context) {
	var sport SportModel
	var slug string

	slug = c.Query("slug")

	sport = GetOneSportDB(slug, c)

	serializer := SportSerializer{c, sport}

	c.JSON(http.StatusOK, serializer.Response())
}

func CreateSport(c *gin.Context) {

	var sport SportModel

	sportModelValidator := NewModelValidator()

	if err := sportModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sport, errC := CreateSportDB(&sportModelValidator.sportModel, c)

	if errC != nil {
		c.JSON(http.StatusUnprocessableEntity, errC)
		return
	}

	serializer := SportSerializer{c, sport}

	c.JSON(http.StatusOK, serializer.Response())

}

func UpdateSport(c *gin.Context) {
	var slug string
	var sport SportModel

	slug = c.Query("slug")

	sportModelValidator := NewModelValidator()

	if err := sportModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sport, errU := UpdateSportDB(slug, &sportModelValidator.sportModel, c)

	if errU != nil {
		c.JSON(http.StatusUnprocessableEntity, errU)
		return
	}

	serializer := SportSerializer{c, sport}

	c.JSON(http.StatusOK, serializer.Response())

}

func DeleteSport(c *gin.Context) {
	var slug string

	slug = c.Query("slug")

	DeleteSportDB(slug, c)

} 