package entity

type User struct {
	Id         int64  `json:"user_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Patronymic string `json:"patronymic"`
	Role       string `json:"role"`
}
