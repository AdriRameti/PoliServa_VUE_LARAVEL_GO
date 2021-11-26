package courts

type CourtModel struct {
	Id	uint `json:"id"`
	Id_sport	uint	`json:"id_sport"`
	Sector	int	`json:"sector"`
	Price_h	int	`json:"price_h"`
}

func (b *CourtModel) TableName() string {
	return "courts"
}