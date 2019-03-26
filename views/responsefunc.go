package views

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)


// service test
func Ping(c echo.Context) error{
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
	result := new(Sentence)
	result.Lang = strings.ToUpper(s.Lang)
	return c.JSON(http.StatusOK,result)
}