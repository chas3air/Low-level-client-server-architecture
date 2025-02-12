package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//var controller routes.CustomersController = &routes.CustomersManager{}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go Handler(conn)
	}
}

func Handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Fprintf(os.Stdout, "%s\n", scanner.Text())
	}
	defer conn.Close()
}
