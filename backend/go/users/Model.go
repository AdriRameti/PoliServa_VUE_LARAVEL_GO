package users

type UserModel struct {
	ID uint `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Surnames string `json:"surnames"`
	Mail string `json:"mail"`
	Pass string `json:"pass"`
	Img string `json:"img"`
	Google2fa_secret string `json:"google2fa_secret"`
	Role string `json:"role"`
	Is_blocked bool `json:"is_blocked"`
}

func (b *UserModel) TableName() string {
	return "users"
}