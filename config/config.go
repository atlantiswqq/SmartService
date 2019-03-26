package config

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"path/filepath"
	"time"
)


func logWriter() io.Writer{
	BasePath,err:=os.Getwd()
	if err !=nil{
		panic(err)
	}
	logPath:= filepath.Join(BasePath,"log","info.log")
	file,err:=os.OpenFile(logPath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err !=nil{
		fmt.Println(err)
	}
	return file
}

func InitConfig(e *echo.Echo){
	// 开启debug模式
	e.Debug = true
	// 设置log路径和级别
	logW:=logWriter()
	e.Logger.SetOutput(logW)
	e.Logger.SetLevel(log.DEBUG)
	// 隐藏echo横幅
	e.HideBanner=false
	// 设置超时时间为60秒
	e.Server.ReadTimeout=60*time.Second
	e.Static("/static","static")
}