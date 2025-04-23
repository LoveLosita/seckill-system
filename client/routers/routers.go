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

	itemGroup := h.Group("/items")
	itemGroup.GET("/get-list", api.GetItemList)
	itemGroup.GET("/get-info", api.GetItemInfo)
	itemGroup.POST("/add", api.AddItem)
	itemGroup.POST("/update", api.UpdateItem)
	itemGroup.POST("/delete", api.DeleteItem)
	h.Spin()
}
