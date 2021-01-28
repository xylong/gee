package middleware

import "github.com/gin-gonic/gin"

// Authorize 授权
type Authorize struct {
}

func NewAuthorize() *Authorize {
	return &Authorize{}
}

func (u *Authorize) OnRequest(ctx *gin.Context) error {
	//fmt.Println("验证授权", ctx.Query("id"))
	//return fmt.Errorf("if error")
	return nil
}
