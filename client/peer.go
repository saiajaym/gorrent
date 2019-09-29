package main

import (
	"fmt"
	"net"
	"os"
)

const port = "localhost:"
const server = "localhost:7777"

func tracker() {
	net.Dial("tcp", server)
}

func peerHandle(con net.Conn) {
	fmt.Println("peer_handle new connection " + con.RemoteAddr().String())

}

func peerManager(handle net.Listener) {
	fmt.Println("peer_manager starting ....")
	for {
		con, err := handle.Accept()
		if err != nil {
			fmt.Println("peer_manager error new client " + err.Error())
		}

		go peerHandle(con)
	}

}
func main() {
	myPort := port + string(os.Args[1])
	handle, err := net.Listen("tcp", myPort)
	if err != nil {
		fmt.Println("Error Listening: " + err.Error())
	}

	go tracker()
	go peerManager(handle)
}
