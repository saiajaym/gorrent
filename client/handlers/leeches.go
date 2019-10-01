package handlers

import (
	"fmt"
	"net"
)

//PeerHandles Handles leech connections
func PeerHandles(con net.Conn) {
	fmt.Println("New Leech Connected: " + con.RemoteAddr().String())
}
