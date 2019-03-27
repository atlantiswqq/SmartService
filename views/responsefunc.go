package views

import (
	"SmartService/mylogconf"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
)


var Plog *privateLog.QscLog{}
*Plog={os.Getwd(),"Smart"}

// service test
func Ping(c echo.Context) error{
	Plog.Info("Test GET request.")
	Plog.All("Test GET request")
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
	jsonStr,err:=json.Marshal(s)
	if err !=nil{
		log.Fatal(err)
	}
	Plog.Debug(string(jsonStr))
	result := new(Sentence)
	result.Lang = strings.ToUpper(s.Lang)
	return c.JSON(http.StatusOK,result)
}