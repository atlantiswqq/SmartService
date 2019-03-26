package main

import (
	"SmartService/config"
	"SmartService/views"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const IP string="0.0.0.0:8001"



func main(){
	e:=echo.New()
	// 初始化配置文件
	config.InitConfig(e)
	// ping is for test request.
	e.GET("/ping",views.Ping)
	//qi bao group is for all method by qi bao robot.
	d:=e.Group("/qibao")
	d.Use(middleware.Logger())
	d.POST("/dialog", views.Dialog)
	e.Logger.Fatal(e.Start(IP))
}