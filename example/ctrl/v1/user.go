package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/gee"
	"github.com/xylong/gee/example/model"
	"net/http"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (user *User) login(ctx *gin.Context) string {
	return "version 1"
}

func (user *User) register(ctx *gin.Context) gee.Object {
	return &model.User{
		ID:     1,
		Name:   "静静",
		Age:    18,
		Gender: 0,
	}
}

func (user *User) friends(ctx *gin.Context) gee.Objects {
	users := []*model.User{
		&model.User{
			ID:     2,
			Name:   "小明",
			Age:    20,
			Gender: 1,
		},
	}
	return gee.MakeObjects(users)
}

func (user *User) Build(gee *gee.Gee) {
	gee.Handle(http.MethodPost, "/login", user.login)
	gee.Handle(http.MethodPost, "/register", user.register)
	gee.Handle(http.MethodGet, "/friends", user.friends)
}
