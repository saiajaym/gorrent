package main

import (
	handlers "PFS/client/handlers"
	"fmt"
	"net"
	"os"

	"github.com/boltdb/bolt"
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

func peerManager(handle net.Listener, db *bolt.DB) {
	fmt.Println("peer_manager starting ....")
	for {
		con, err := handle.Accept()
		if err != nil {
			fmt.Println("peer_manager error new client " + err.Error())
		}
		go handlers.LeechHandle(con, db)
	}

}
func main() {
	myPort := port + string(os.Args[1])
	fmt.Println("Listening on my port: " + myPort)
	handle, err := net.Listen("tcp", myPort)
	if err != nil {
		fmt.Println("Error Listening: " + err.Error())
		os.Exit(0)
	}
	handlers.Server = server
	if !handlers.CheckTracker() {
		os.Exit(1)
	}
	db, _ := handlers.DBOpen(myPort + ".db")
	handlers.DB = db
	go peerManager(handle, db)

	handlers.Console(db)

}
