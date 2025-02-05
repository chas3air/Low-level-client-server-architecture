package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"server/internal/routes"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	var controller routes.CustomersController = &routes.CustomersManager{}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go Handler(controller, conn)
	}
}

// TODO: нужно описать логику соответствия строки запроса
// и проверить запрос, в случае нужного заголовка вызвать его
// функцию и соответствующий метод
func Handler(cc routes.CustomersController, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req_line := scanner.Text()
		req_params := strings.Fields(req_line)

		if req_params[0] == http.MethodGet {
			if strings.Split(req_params[1], "/")[1] != "" {
				cc.GetUsersHandler(conn)
			} else {
				cc.GetUserByIdHandler(conn)
			}
		} else if req_params[0] == http.MethodPost {
			cc.PostUsersHandler(conn)
		} else if req_params[0] == http.MethodPut {
			cc.PutUsersHandler(conn)
		} else if req_params[0] == http.MethodDelete {
			cc.DeleteUsersHandler(conn)
		} else {
			log.Println("Method Not Allowed")
			break
		}
	}
	defer conn.Close()
}
