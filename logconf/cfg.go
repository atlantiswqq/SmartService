package privateLog

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"path/filepath"
)

type QscLog struct {
	output      *io.Writer
	basePath    string
	projectName string
}

func (q *QscLog) SetBasePath(path string){
	q.basePath = path
}

func (q *QscLog) SetProjectName(name string){
	q.projectName=name
}

func (q *QscLog) Fatal(v ...interface{}){
	fileName:= filepath.Join(q.basePath,"log",q.projectName+"-Fatal.log")
	fatalWriter,err:=os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Println(err)
	}
	l:=log.New("-")
	l.SetOutput(fatalWriter)
	l.Fatal(v)
}

func (q *QscLog) Warn(v ...interface{}){
	fileName:= filepath.Join(q.basePath,"log",q.projectName+"-Warn.log")
	warnWriter,err:=os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Println(err)
	}
	l:=log.New("-")
	l.SetOutput(warnWriter)
	l.Warn(v)
}

func (q *QscLog) Error(v ...interface{}){
	fileName:= filepath.Join(q.basePath,"log",q.projectName+"-Error.log")
	errorWriter,err:=os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Println(err)
	}
	l:=log.New("-")
	l.SetOutput(errorWriter)
	l.Error(v)
}

func (q *QscLog) Info(v ...interface{}){
	fileName:= filepath.Join(q.basePath,"log",q.projectName+"-Info.log")
	infoWriter,err:=os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Println(err)
	}
	l:=log.New("-")
	l.SetOutput(infoWriter)
	l.Info(v)
}

func (q *QscLog) Debug(v ...interface{}){
	fileName:= filepath.Join(q.basePath,"log",q.projectName+"-Debug.log")
	debugWriter,err:=os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Println(err)
	}
	l:=log.New("-")
	l.SetOutput(debugWriter)
	l.Debug(v)
	l.SetLevel(0)
}

func (q *QscLog) All(v ...interface{}){
	fileName:= filepath.Join(q.basePath,"log",q.projectName+"-All.log")
	allWriter,err:=os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Println(err)
	}
	l:=log.New("-")
	l.SetOutput(allWriter)
	l.Info(v)
}



