package reservations

import (
	"github.com/gin-gonic/gin"
)

/**SERIALIZERS CREATE RESERVATION**/

type ReservationSerializer struct {
	C *gin.Context
	ReservationModel
}

type ReservationResponse struct{
	Id          uint   `json:"id"`
	Id_user     uint   `json:"id_user"`
	Id_court    uint   `json:"id_court"` 
	Date        string `json:"date"`
	Hini        string `json:"hini"`
	Hfin        string `json:"hfin"`
	Total_price int    `json:"total_price"`
}

func (r *ReservationSerializer) Response() ReservationResponse {
	response := ReservationResponse{
		Id: r.Id,
		Id_user: r.Id_user,    
		Id_court: r.Id_court,
		Date: r.Date,       
		Hini: r.Hini,       
		Hfin: r.Hfin,        
		Total_price: r.Total_price, 
	}

	return response
}

/**SERIALIZERS GET RESERVATIONS**/
type ReservationsSerializer struct {
	C *gin.Context
	Reservation []ReservationModel
}

func (R *ReservationsSerializer) Response() []ReservationResponse {
	response := []ReservationResponse{}
	for _, reservation := range R.Reservation {
		serializer := ReservationSerializer{R.C, reservation}
		response = append(response, serializer.Response())
	}
	return response
}