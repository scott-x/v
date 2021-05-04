package main

import (
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"log"
	"os"
	"path"
)

var TEMPLATES string

func init(){
	TEMPLATES = os.Getenv("GOPATH")+"/src/github.com/scott-x/v/templates"
}

func main() {
	option := cmd.AddTask("please select the task:", 7, "Add vue.config.js", "Add component", "Add View(Page)")
	switch option {
	case "1":
		//do something
		task1()
	//anycode here ...
	case "2":
		task2()
		//do something
	case "3":
		task3()
		//do something
	default:
		//do something
		log.Println("input error, program is exiting....")
	}
}

func task2() {
	log.Println("task2")
}

func task3() {
	log.Println("task3")
}

func task1(){
	flag:=fs.IsExist("./package.json")
	if !flag{
		cmd.Warning("You should move to the project root folder, then re-try!")
		return
	}
	configfile:= path.Join(TEMPLATES,"vue.config.js")
	_,err:=fs.CopyFile(configfile,"./vue.config.js")
	if err!=nil{
		cmd.Warning("internal error")
		return
	}
	cmd.Info("generate vue.config.js successfully!")
}