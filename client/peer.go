package main

import (
	"fmt"
	"net"
	"os"
)

const port = "localhost:"
const server = "localhost:7777"

func tracker() {
	con, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Connection to tracker failed, exiting program: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Connection to tracker Successful: " + con.RemoteAddr().String())

}



func peerManager(handle net.Listener) {
	fmt.Println("peer_manager starting ....")
	for {
		con, err := handle.Accept()
		if err != nil {
			fmt.Println("peer_manager error new client " + err.Error())
		}

		go PeerHandle(con)
	}

}
func main() {
	myPort := port + string(os.Args[1])
	fmt.Println("Listening on my port: " + myPort)
	handle, err := net.Listen("tcp", myPort)
	if err != nil {
		fmt.Println("Error Listening: " + err.Error())
	}

	go tracker()
	go peerManager(handle)
	for {

	}
}
