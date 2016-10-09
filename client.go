package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/nsf/termbox-go"
	"io"
	"net"
	"time"
)

var playField = [10][10]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}

var screenRefreshRate = 1 * time.Second

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	defer termbox.Close()

	conn := connectToServer

	//Get playfield from server

	//go updatePlayfield()
	//Redraw Screen
	go updateScreen()
	//Send input to server
	//go handleInput(conn)
	time.Sleep(10 * time.Second)
}

func connectToServer() net.Conn {

	var connHost = "localhost"
	var connPort = "9000"
	var connType = "tcp"

	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		panic(err)
	}

	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// Handle connections in a new goroutine.

		defer conn.Close()
		return conn
	}
}

func handleInput(conn net.Conn) {
	//Read input
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				sendInput("escape", conn)
			case termbox.KeyArrowLeft:
				sendInput("left", conn)
			case termbox.KeyArrowRight:
				sendInput("right", conn)
			case termbox.KeyArrowDown:
				sendInput("down", conn)
			case termbox.KeyArrowUp:
				sendInput("Up", conn)
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

func sendInput(input string, conn net.Conn) {
	io.WriteString(conn, fmt.Sprint(input))
}
