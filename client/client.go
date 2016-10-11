package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/nsf/termbox-go"
	"time"
)

var screenRefreshRate = 1 * time.Second
var gs = GetGameServer()
var gc = GetGameClient()

func main() {

	initGame()

	//Get playfield from server

	go updatePlayfield()
	//Redraw Screen
	go updateScreen()
	//Send input to server
	go handleInput()
	time.Sleep(10 * time.Second)
}

func initGame() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	defer termbox.Close()
}

func handleInput() {
	//Read input
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				fmt.Println("esc pressed")
				gs.sendString("escape")
			case termbox.KeyArrowLeft:
				gs.sendString("left")
			case termbox.KeyArrowRight:
				gs.sendString("right")
			case termbox.KeyArrowDown:
				gs.sendString("down")
			case termbox.KeyArrowUp:
				gs.sendString("Up")
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func updateScreen() {

	for {
		clearScreen()

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				fmt.Printf(codeToBlock(playField[i][j]))
			}
			fmt.Printf("\n")
		}
		time.Sleep(screenRefreshRate)
	}
}

func updatePlayfield() {
	for {
		//Read from channel
		playField
	}
}

func clearScreen() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
}

func codeToBlock(i int) string {
	if i == 1 {
		return "██"
	}
	return "░░"
}
