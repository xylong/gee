package model

type User struct {
	ID     int    `json:"id" uri:"id" binding:"required,gt=0"`
	Name   string `json:"name"`
	Age    uint8  `json:"age"`
	Gender int    `json:"gender"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) String() string {
	return ""
}
