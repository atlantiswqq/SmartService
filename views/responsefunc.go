package views

import (
	"SmartService/logconf"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strings"
)

var pLog *privateLog.QscLog

func init() {
	basePath,err:= os.Getwd()
	if err !=nil{
		panic(err)
	}
	pLog = new(privateLog.QscLog)
	//parentPath:= filepath.Dir(basePath)
	pLog.SetBasePath(basePath)
	pLog.SetProjectName("Smart")
	pLog.All("log init success")
}
// service test
func Ping(c echo.Context) error{
	pLog.Info("Test GET request.")
	pLog.All("Test GET request")
	return c.String(http.StatusOK,"service is ok.")
}

// qi bao smart service
type Sentence struct{
	Lang string `json:"lang"`
}
func Dialog(c echo.Context) error{
	s:=new(Sentence)
	if err:=c.Bind(s);err !=nil{
		return err
	}
	pLog.Debug(fmt.Sprintf("%+v\n",*s))
	result := new(Sentence)
	result.Lang = strings.ToUpper(s.Lang)
	pLog.All(fmt.Sprintf("%+v",*result))
	return c.JSON(http.StatusOK,result)
}