package privateLog

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)


type QscLog struct{
	BasePath string
	ProjectName string
}

func (q *QscLog) SetBasePath(path string) *QscLog{
	q.BasePath = path
	return q
}

func (q *QscLog) SetProjectName(name string) *QscLog{
	q.ProjectName= name
	return q
}
func (q *QscLog) GetOsFile(filename string)(*os.File,error){
	fileAbsPath:=filepath.Join(q.BasePath,"log",q.ProjectName+"-"+filename)
	file,err:=os.OpenFile(fileAbsPath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	if err !=nil{
		return nil,err
	}
	return file,nil
}

func (q *QscLog) Fatal(s string){
	logFile,err:=q.GetOsFile("Fatal")
	if err !=nil{
		panic(errors.New("Fatal log init error."))
	}
	_,err=io.WriteString(logFile,s+"\n")
	if err !=nil{
		panic(errors.New("write into fatal log error."))
	}
}

func (q *QscLog) Error(s string){
	logFile,err:=q.GetOsFile("Error")
	if err !=nil{
		panic(errors.New("Error log init error."))
	}
	_,err=io.WriteString(logFile,s+"\n")
	if err !=nil{
		panic(errors.New("write into Error log error."))
	}
}
func (q *QscLog) Warn(s string){
	logFile,err:=q.GetOsFile("Warn")
	if err !=nil{
		panic(errors.New("Warn error init error."))
	}
	_,err=io.WriteString(logFile,s+"\n")
	if err !=nil{
		panic(errors.New("write into Warn log error."))
	}
}
func (q *QscLog) Info(s string){
	logFile,err:=q.GetOsFile("Info")
	if err !=nil{
		panic(errors.New("Info error init error."))
	}
	_,err=io.WriteString(logFile,s+"\n")
	if err !=nil{
		panic(errors.New("write into Info log error."))
	}
}
func (q *QscLog) Debug(s string){
	logFile,err:=q.GetOsFile("Debug")
	if err !=nil{
		panic(errors.New("Debug error init error."))
	}
	_,err=io.WriteString(logFile,s+"\n")
	if err !=nil{
		panic(errors.New("write into Debug log error."))
	}
}
func (q *QscLog) Trace(s string){
	logFile,err:=q.GetOsFile("Trace")
	if err !=nil{
		panic(errors.New("Trace error init error."))
	}
	_,err=io.WriteString(logFile,s+"\n")
	if err !=nil{
		panic(errors.New("write into Trace log error."))
	}
}

func (q *QscLog) All(s string){
	logFile,err:=q.GetOsFile("All")
	if err !=nil{
		panic(errors.New("All error init error."))
	}
	_,err=io.WriteString(logFile,s+"\n")
	if err !=nil{
		panic(errors.New("write into All log error."))
	}
}
