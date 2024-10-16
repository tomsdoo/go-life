package main

import (
	"fmt"
	"go-life/life"
	"time"
)

func main() {
	board := life.NewBoard(20,10)
	board.Seed()
	for i := 0; i < 300; i++ {
		fmt.Print("\x0c", board.String())
		time.Sleep(time.Second / 30)
		changed := board.GoToNext()
		if changed == false {
			break
		}
	}
}

