package main

import (
	"time"
)

func main() {
	game := &Game{}
	go game.Run()

	time.Sleep(1 * time.Hour)
}
