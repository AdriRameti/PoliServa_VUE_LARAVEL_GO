package courts

import (
	"github.com/gin-gonic/gin"
)

type CourtResponse struct {
	Id	uint `json:"id"`
	Id_sport	uint	`json:"id_sport"`
	Price_h	int	`json:"price_h"`
	Sector	int	`json:"sector"`
	Img string `json:"img"`
}

type CourtSerializer struct {
	C *gin.Context
	court CourtModel
}

type CourtsSerializer struct {
	C *gin.Context
	courts []CourtModel
}

func (cs *CourtSerializer) Response() CourtResponse {
	response := CourtResponse {
		Id: cs.court.Id,
		Id_sport: cs.court.Id_sport,
		Price_h: cs.court.Price_h,
		Sector: cs.court.Sector,
		Img: cs.court.Img,
	}

	return response
}

func (cs *CourtsSerializer) Response() []CourtResponse {
	response := []CourtResponse{}

	for _, court := range cs.courts {
		serializer := CourtSerializer{cs.C, court}
		response = append(response, serializer.Response())
	}

	return response
}