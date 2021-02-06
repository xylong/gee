package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/gee"
	"net/http"
)

func NewHome() *Home {
	return &Home{}
}

type Home struct{}

func (h *Home) Index(ctx *gin.Context) gee.View {
	return "home"
}

func (h *Home) Build(gee *gee.Gee) {
	gee.Handle(http.MethodGet, "/", h.Index)
}
