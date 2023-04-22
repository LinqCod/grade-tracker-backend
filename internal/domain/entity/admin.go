package entity

type Admin struct {
	User
}

type AdminDTO struct {
	Id         int64  `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Patronymic string `json:"patronymic"`
}

type AdminLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
