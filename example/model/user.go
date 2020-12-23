package model

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    uint8  `json:"age"`
	Gender int    `json:"gender"`
}

func (user *User) String() string {
	return ""
}
