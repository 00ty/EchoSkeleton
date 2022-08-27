package utils

import (
	"io/ioutil"
	"os"

	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
)

type conf struct {
	WebPort    string `yaml:"web_port"`
	DbHost     string `yaml:"host"`
	DbPort     string `yaml:"port"`
	DbUser     string `yaml:"username"`
	DbPassWord string `yaml:"password"`
	DbName     string `yaml:"dbname"`
}

var Config conf

func GetConfig() {

	content, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Error("配置文件读取错误: %v", err)
		os.Exit(3)
	}
	if yaml.Unmarshal(content, &Config) != nil {
		log.Error("解析配置文件出错: %v", err)
		os.Exit(3)
	}
	//  = config
	/* 	WebPort = config.WebPort
	   	DbHost = config.Host
	   	DbPort = config.Port
	   	DbUser = config.Username
	   	DbPassWord = config.Password
	   	DbName = config.DBname */
}
