package main

import (
	"log"
	"net"
)

func main() {
	dealer, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Close()

	for {
		// TODO:описать логику
	}
}

func Handler(conn net.Conn) {
	panic("Not Implemented")
}
