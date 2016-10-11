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
	clearScreen()
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
	fmt.Println("Connected to: " + gs.connHost)
	fmt.Println()
	fmt.Println("Use arrow keys")
	fmt.Println("Press Esc to exit")
	fmt.Println()
	fmt.Println()
}

func updateScreen() {

	for {

		printGameInfo()

		for _, h := range gc.playField {
			for _, cell := range h {
				fmt.Printf(codeToBlock(cell))
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

	switch i {
	case 1:
		return "█"
	case 2:
		return "@"
	case 3:
		return "="
	}
	return "░"
}
