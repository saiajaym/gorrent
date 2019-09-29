package main

import (
	"PFS/server/handles"
	"fmt"
	"net"

	"github.com/boltdb/bolt"
)

const port = "localhost:7777"
const db = "tracker.db"

func handleClient(con net.Conn, db *bolt.DB) {
	fmt.Println("recieved req client" + con.RemoteAddr().String())
	cli := handles.Client{
		Con: con,
		Db:  db}
	
}

func main() {
	//TO-DO - Init DB
	db, err := handles.DBOpen(db)
	handle, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening: " + err.Error())
	}

	for {
		con, err := handle.Accept()
		if err != nil {
			fmt.Println("Error new cleint: " + err.Error())
		}
		go handleClient(con, db)

	}
}
