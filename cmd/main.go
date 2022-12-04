package main

import (
	"EchoSkeleton/api/routes"
	"fmt"

	"EchoSkeleton/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// var Db *DB

func main() {
	// 配置日志
	var defaultLog = `${time_rfc3339} ${level} file=${long_file} line=${line}` + "\n>"
	log.SetHeader(defaultLog)
	// 载入配置
	utils.GetConfig()
	utils.InitDb()
	e := echo.New()
	routes.Run(e)
	fmt.Println("web is http://127.0.0.1:" + utils.Config.WebPort)
	e.Logger.Fatal(e.Start(":" + utils.Config.WebPort))
}

/*func DbConnect(cof DbConf) {

}*/
