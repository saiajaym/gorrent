package main

import (
	"os"
)

const port = "localhost:"
const server = "localhost:7777"

func tracker() {

}

func main() {
	my_port := port + string(os.Args[1])
	go tracker()
}
