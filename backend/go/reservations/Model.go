package reservations

type ReservationModel struct {
	Id	uint `json:"id"`
	Id_user	string `json:"slug"`
	Id_court	string	`json:"name"`
	Date	string	`json:"date"`
	Hini	string	`json:"hini"`
	Hfin	string	`json:"hfin"`
	Total_price int `json:"total_price"`
}

func (b *ReservationModel) TableName() string {
	return "reservations"
}
