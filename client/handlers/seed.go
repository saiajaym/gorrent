package handlers

import (
	"fmt"
	"net"
)

//SeedHandle Handles incoming seed requests
func SeedHandle(con net.Conn) {
	fmt.Println("SeedHandles")
}
