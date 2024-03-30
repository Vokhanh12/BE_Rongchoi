package entity

type User struct {
	ID          string "json:id"
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	NickName    string `json:"nick_name"`
	NumberPhone string `json:"number_phone"`
	DayOfBirth  string `json:"day_of_birth"`
	Address     string `json:"address"`
	Role        string `json:"role"`
}
