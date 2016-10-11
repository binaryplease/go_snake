package main

import (
	_ "fmt"
	_ "io"
	"net"
)

type gameServer struct {
	connToServer   net.Conn
	connFromServer net.Conn
}

//GetGameServer constructor for Gameserver
func GetGameServer() gameServer {
	gs := gameServer{}
	gs.connect()
	return gs
}

func (gs gameServer) connect() {
	var connHost = "localhost"
	var connPort = "9000"
	var connType = "tcp"

	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		panic(err)
	}

	defer l.Close()

	//for {
	//// Listen for an incoming connection.
	//gs.connFromServer, err = l.Accept()
	//if err != nil {
	//panic(err)
	//}
	//// Handle connections in a new goroutine.

	//defer gs.connFromServer.Close()
	//}

}

func (gs gameServer) getPlayfield() [][]int {
	return playField
}

func (gs gameServer) sendString(input string) {
	//io.WriteString(gs.connToServer, fmt.Sprint(input))
}
