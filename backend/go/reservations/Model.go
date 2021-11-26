package reservations

type ReservationModel struct {
	Id	uint `json:"id"`
	Id_user	string `json:"slug"`
	Id_court	string	`json:"name"`
	Fini	string	`json:"fini"`
	Ffin	string	`json:ffin`
	Total_price int `json:total_price`
}

func (b *ReservationModel) TableName() string {
	return "reservations"
}