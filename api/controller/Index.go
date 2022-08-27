package controller

import (
	"EchoSkeleton/api/dao"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type IndexController struct{}

func (this IndexController) Index(c echo.Context) error {
	log.Infof("我日")
	return c.String(http.StatusOK, "Hello, World!")
}
func (this IndexController) Echo(c echo.Context) error {

	log.Infof("测试")
	fmt.Println(dao.NewUserDao().Find(11111))
	return c.String(http.StatusOK, "初露锋芒")
}
