package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

var screenRefreshRate = 1 * time.Second
var gs = GetGameServer()
var gc = GetGameClient()

func main() {

	fmt.Println("test")
	initGame()

	//Get playfield from server

	go updatePlayfield()
	//Redraw Screen
	go updateScreen()
	//Send input to server
	go handleInput()

	for {
	}
}

func initGame() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
}

func handleInput() {
	//Read input
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				gs.sendString("escape")
				os.Exit(0)
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

func printGameInfo() {
	fmt.Println("ASCII SNAKE")
	fmt.Println("")
	fmt.Println("Use arrow keys")
	fmt.Println("Press Esc to exit")
}

func updateScreen() {

	for {

		printGameInfo()
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				fmt.Printf(codeToBlock(playField[i][j]))
			}
			fmt.Printf("\n")
		}
		time.Sleep(screenRefreshRate)
		clearScreen()
	}
}

func updatePlayfield() {
	for {

		//Read from channel
		//playField
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
