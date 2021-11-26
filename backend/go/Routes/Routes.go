package Routes

import (
	"poliserva/sports"
	"poliserva/courts"
	"poliserva/reservations"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api");

	sports.SportsRouting(api.Group("/sports"))
	courts.CourtsRouting(api.Group("/courts"))
	reservations.ReservationRouting(api.Group("/reservations"))

	return r
}