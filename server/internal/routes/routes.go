package routes

import "net"

type CustomersController interface {
	GetUsersHandler(conn net.Conn)
	GetUserByIdHandler(conn net.Conn)
	InsertUserHandler(conn net.Conn)
	UpdateUserHandler(conn net.Conn)
	DeleteUserHandler(conn net.Conn)
}

type CustomersManager struct{}

func (cm *CustomersManager) GetUsersHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) GetUserByIdHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) InsertUserHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) UpdateUserHandler(conn net.Conn) {
	panic("Not Implemented")
}

func (cm *CustomersManager) DeleteUserHandler(conn net.Conn) {
	panic("Not Implemented")
}
