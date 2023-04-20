package entity

type Student struct {
	User
	Group Group `json:"group"`
}

type StudentDTO struct {
	Id         int64  `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Patronymic string `json:"patronymic"`
	Group      Group  `json:"group"`
}

type StudentRegistrationDTO struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Patronymic string `json:"patronymic"`
	Group      Group  `json:"group"`
}

type StudentLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
