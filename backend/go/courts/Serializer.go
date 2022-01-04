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
	Sport SportModel `json:"Sport"`
}

type CourtSerializer struct {
	C *gin.Context
	Court CourtModel
}

type CourtsSerializer struct {
	C *gin.Context
	Courts []CourtModel
}

func (cs *CourtSerializer) Response() CourtResponse {
	response := CourtResponse {
		Id: cs.Court.Id,
		Id_sport: cs.Court.Id_sport,
		Price_h: cs.Court.Price_h,
		Sector: cs.Court.Sector,
		Img: cs.Court.Img,
		Sport: cs.Court.Sports,
	}

	return response
}

func (cs *CourtsSerializer) Response() []CourtResponse {
	response := []CourtResponse{}

	for _, court := range cs.Courts {
		serializer := CourtSerializer{cs.C, court}
		response = append(response, serializer.Response())
	}

	return response
}