package mylog

import (
	"log"
	"os"
)

func InitLog(path string) {
	f, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Create log file faied!")
		return
	}

	log.SetOutput(f)
}