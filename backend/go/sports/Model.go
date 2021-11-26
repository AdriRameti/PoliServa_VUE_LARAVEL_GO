package sports

type SportModel struct {
	Id	uint `json:"id"`
	Slug	string `json:"slug"`
	Name	string	`json:"name"`
	Img	string	`json:"img"`
}

func (b *SportModel) TableName() string {
	return "sports"
}