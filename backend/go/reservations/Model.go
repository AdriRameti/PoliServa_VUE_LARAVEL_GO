package reservations

type ReservationModel struct {
	Id	uint `json:"id"`
	Id_user	uint `json:"id_user"`
	Id_court	uint	`json:"id_court"`
	Date	string `json:"date"`
	Hini	string	`json:"hini"`
	Hfin	string	`json:"hfin"`
	Total_price int `json:"total_price"`
}

func (b *ReservationModel) TableName() string {
	return "reservations"
}