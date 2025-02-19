package app

import (
	"bufio"
	interfaces "client/internal/domain/interfaces/userservice"
	"client/internal/domain/models"
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
)

type App struct {
	log             *slog.Logger
	userservice     interfaces.UserService
	port            int
	expiration_time time.Duration
}

func New(log *slog.Logger, userservice interfaces.UserService, port int, expiration_time time.Duration) *App {
	return &App{
		log:             log,
		userservice:     userservice,
		port:            port,
		expiration_time: expiration_time,
	}
}

func (a *App) Start() {
	const op = "app.Start"
	var choise string = ""

	for {
		fmt.Println("1. Get users")
		fmt.Println("2. Get user by id")
		fmt.Println("3. Get user by email")
		fmt.Println("4. Insert")
		fmt.Println("5. Update")
		fmt.Println("6. Delete")
		fmt.Println("7. Exit")

		fmt.Scanf("%d", &choise)
		switch choise {
		case "1":
			fmt.Println("Get users")
			context, cancel := context.WithDeadline(context.Background(), time.Now().Add(a.expiration_time))
			defer cancel()
			users, err := a.userservice.GetUsers(context)
			if err != nil {
				// TODO: not implemented

				fmt.Println("Error fetching users")
			}

			fmt.Println("Users:")
			for _, user := range users {
				fmt.Println(user)
			}

			fmt.Println("Press Enrer to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "2":
			fmt.Println("Get users by id")
			var uuid_id string
			fmt.Scanf("%s", &uuid_id)
			parsedUUID, err := uuid.Parse(uuid_id)
			if err != nil {
				// TODO: not implemented

				break
			}

			context, cancel := context.WithDeadline(context.Background(), time.Now().Add(a.expiration_time))
			defer cancel()
			user_by_id, err := a.userservice.GetUserById(context, parsedUUID)
			if err != nil {
				// TODO: not implemented

				break
			}

			fmt.Println("User by id: " + uuid_id + ":")
			fmt.Println(user_by_id)

			fmt.Println("Press Enrer to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "3":
			fmt.Println("Get user by email")
			var email string
			fmt.Scanf("%s", &email)

			context, cancel := context.WithDeadline(context.Background(), time.Now().Add(a.expiration_time))
			defer cancel()
			user, err := a.userservice.GetUserByEmail(context, email)
			if err != nil {
				// TODO: not implemented

				break
			}

			fmt.Println("User by email: " + email + ":")
			fmt.Println(user)

			fmt.Println("Press Enrer to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "4":
			fmt.Println("Insert")
			user_for_insert := models.NewUser()

			context, cancel := context.WithDeadline(context.Background(), time.Now().Add(a.expiration_time))
			defer cancel()

			err := a.userservice.Insert(context, *user_for_insert)
			if err != nil {
				// TODO: not implemented

				break
			}

			fmt.Println("User inserted successfully")
			fmt.Println("Press Enrer to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "5":
			fmt.Println("Update")

			user_for_update := models.NewUser()

			context, cancel := context.WithDeadline(context.Background(), time.Now().Add(a.expiration_time))
			defer cancel()

			err := a.userservice.Update(context, user_for_update.Id, *user_for_update)
			if err != nil {
				// TODO: not implemented

				break
			}

			fmt.Println("User updated successfully")
			fmt.Println("Press Enrer to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "6":
			fmt.Println("Delete")
			var uuid_id string
			fmt.Scanf("%s", &uuid_id)
			id, err := uuid.Parse(uuid_id)
			if err != nil {
				// TODO: not implemented

				break
			}

			context, cancel := context.WithDeadline(context.Background(), time.Now().Add(a.expiration_time))
			defer cancel()

			user, err := a.userservice.Delete(context, id)
			if err != nil {
				// TODO: not implemented

				break
			}

			fmt.Println("User deleted successfully")
			fmt.Println(user)

			fmt.Println("Press Enrer to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "7":
			fmt.Println("Exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			return

		default:
			fmt.Println("Press Enrer to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			return
		}
	}
}
