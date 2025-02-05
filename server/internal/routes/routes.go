package routes

import "net"

type CustomersController interface {
	GetUsersHandler(conn net.Conn)
	GetUserByIdHandler(conn net.Conn)
	PostUsersHandler(conn net.Conn)
	PutUsersHandler(conn net.Conn)
	DeleteUsersHandler(conn net.Conn)
}

type CustomersManager struct{}

func (cm *CustomersManager) GetUsersHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) GetUserByIdHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) PostUsersHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) PutUsersHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) DeleteUsersHandler(conn net.Conn) {
	panic("Not Implemented")
}
