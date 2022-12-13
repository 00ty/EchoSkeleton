package routes

import (
	c "EchoSkeleton/api/controller"

	"github.com/labstack/echo/v4"
)

func Run(e *echo.Echo) {
	// 关闭条幅
	e.HideBanner = true
	// 配置路由
	e.GET("/", c.IndexController{}.Index)
	e.GET("/echo", c.IndexController{}.Echo)
	e.POST("/user", c.UserController{}.Register)
}
