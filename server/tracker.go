package main

import (
	"fmt"
	"net"
)

const port = "localhost:7777"

func handleClient(con net.Conn) {
	fmt.Println("recieved req")
}

func main() {
	//TO-DO - Init DB
	handle, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening: " + err.Error())
	}

	for {
		con, err := handle.Accept()
		if err != nil {
			fmt.Println("Error new cleint: " + err.Error())
		}
		go handleClient(con)

	}
}
