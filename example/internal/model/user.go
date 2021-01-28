package model

type User struct {
	ID     int    `gorm:"id" xorm:"id" json:"id" uri:"id" binding:"required,gt=0"`
	Name   string `gorm:"name" xorm:"name" json:"name"`
	Age    uint8  `gorm:"age" xorm:"age" json:"age"`
	Gender int    `gorm:"gender" xorm:"gender" json:"gender"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) String() string {
	return ""
}
