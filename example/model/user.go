package model


type User struct {
	ID     int    `gorm:"id" json:"id" uri:"id" binding:"required,gt=0"`
	Name   string `gorm:"name" json:"name"`
	Age    uint8  `gorm:"age" json:"age"`
	Gender int    `gorm:"gender" json:"gender"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) String() string {
	return ""
}
