package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hzy/util"

	ct "github.com/daviddengcn/go-colortext"
)

const checkInter int64 = 1000

var (
	logFile *log.Logger
	create  int64
	nextCk  int64
	fPath   string
	fName   string
	fSize   int32 = 512 * 1024 * 1024
)

// Launcher 启动日志
func Launcher(path string, name string) bool {
	err := os.MkdirAll(path, 666)
	if err != nil {
		log.Printf("Launcher | %v", err)
		return false
	}
	t := time.Now()
	f, err1 := os.Create(fmt.Sprintf("%s/%s-%4d-%2d-%2d_%2d_%2d_%2d.log", path, name,
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
	if err1 != nil {
		log.Printf("Launcher | %v", err1)
		return false
	}
	logFile = log.New(f, "", log.Ldate|log.Lmicroseconds)
	initArgs(path, name)

	log.SetFlags(log.Lmicroseconds)
	return true
}

// Info info print and write file
func Info(format string, args ...interface{}) {
	printLog(fmt.Sprintf(format, args...), ct.Green)
}

// Warning warn print and write file
func Warning(format string, args ...interface{}) {
	printLog(fmt.Sprintf(format, args...), ct.Yellow)
}

// Error err print and write file
func Error(format string, args ...interface{}) {
	printLog(fmt.Sprintf(format, args...), ct.Red)
}

/////////////////////////////////////////////
func writeFile(msg string) {
	logFile.Println(msg)
	if nextCk > util.GetTs() {
		checkFile()
	}
}

func checkFile() {
	if !util.IsToday(create) {
		Launcher(fPath, fName)
	}
}

func initArgs(path string, name string) {
	create, nextCk, fPath, fName = util.GetTs(), create+checkInter, path, name
}

func printLog(msg string, clr ct.Color) {
	ct.Background(clr, true)
	log.Print(msg)
	ct.ResetColor()
	writeFile(msg)
}
