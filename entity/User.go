package entity

type User struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	Password    string `json: "password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	//Books       []Book `json:"books"`
}
