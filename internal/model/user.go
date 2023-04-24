package model

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserWithId struct {
	User
	Id int `json:"id"`
}

type UserWithRole struct {
	UserWithId
	Role string `json:"role"`
}
