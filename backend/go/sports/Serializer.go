package sports

import (
	"github.com/gin-gonic/gin"
)

type SportResponse struct {
	Id	uint `json:"id"`
	Slug	string `json:"slug"`
	Name	string	`json:"name"`
	Img	string	`json:"img"`
}

type SportSerializer struct {
	C *gin.Context
	sport SportModel
}

type SportsSerializer struct {
	C *gin.Context
	sports []SportModel
}

func (ss *SportSerializer) Response() SportResponse {
	response := SportResponse {
		Id: ss.sport.ID,
		Slug: ss.sport.Slug,
		Name: ss.sport.Name,
		Img: ss.sport.Img,
	}

	return response
}

func (ss *SportsSerializer) Response() []SportResponse {
	response := []SportResponse{}

	for _, sport := range ss.sports {
		serializer := SportSerializer{ss.C, sport}
		response = append(response, serializer.Response())
	}

	return response
}