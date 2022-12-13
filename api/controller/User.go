package controller

import (
	"EchoSkeleton/api/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func (this UserController) Register(c echo.Context) error {
	rJson := new(dto.Register)
	if err := c.Bind(rJson); err != nil {
		return c.JSON(http.StatusOK, &dto.Response{
		    Code:  400,
		    Msg: "请输入完整",
		})
	}
	if rJson.UserName == "" {
		return c.JSON(http.StatusOK, &dto.Response{
		    Code:  400,
		    Msg: "请输入完整",
		})
	}
	return c.JSON(http.StatusOK, rJson)
}
