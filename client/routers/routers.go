package routers

import (
	"client/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRouters() {
	h := server.Default()

	userGroup := h.Group("/user")

	userGroup.GET("/login", api.UserLogin)
	userGroup.POST("/register", api.UserRegister)
	userGroup.GET("/refresh_token", api.RefreshToken)
	h.Spin()
}
