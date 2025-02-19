package server

import (
	"client/internal/domain/models"
	"client/internal/domain/profilers"
	"context"
	"fmt"
	"log/slog"

	umv1 "github.com/chas3air/protos/gen/go/usersManager"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ServerUsersStorage struct {
	log        *slog.Logger
	ServerHost string
	ServerPort int
}

func New(log *slog.Logger, host string, port int) *ServerUsersStorage {
	return &ServerUsersStorage{
		ServerHost: host,
		ServerPort: port,
		log:        log,
	}
}

// GetUsers implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "storage.server.getUsers"
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.ServerHost, s.ServerPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	c := umv1.NewUsersManagerClient(conn)
	res, err := c.GetUsers(ctx, nil)
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var res_users []models.User = make([]models.User, 0, len(res.Users))
	for _, pb_user := range res.Users {
		user, err := profilers.ProtoUsrToUsr(pb_user)
		if err != nil {
			s.log.Error(fmt.Sprintf("%s: failed to convert proto user to model user: %v", op, err))
			continue
		}
		res_users = append(res_users, user)
	}

	return res_users, nil
}

// GetUserById implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) GetUserById(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "storage.server.getUserById"
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.ServerHost, s.ServerPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	c := umv1.NewUsersManagerClient(conn)
	res, err := c.GetUserById(ctx, &umv1.GetUserByIdRequest{
		Id: uid.String(),
	})
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user, err := profilers.ProtoUsrToUsr(res.GetUser())
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: failed to convert proto user to model user: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// GetUserByEmail implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "storage.server.getUserByEmail"
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.ServerHost, s.ServerPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	c := umv1.NewUsersManagerClient(conn)
	res, err := c.GetUserByEmail(ctx, &umv1.GetUserByEmailRequest{
		Email: email,
	})
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user, err := profilers.ProtoUsrToUsr(res.GetUser())
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: failed to convert proto user to model user: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// Insert implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) Insert(ctx context.Context, user models.User) error {
	const op = "storage.server.insert"
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.ServerHost, s.ServerPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	c := umv1.NewUsersManagerClient(conn)
	_, err = c.Insert(ctx, &umv1.InsertRequest{
		User: profilers.UsrToProroUsr(user),
	})
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Update implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) Update(ctx context.Context, uid uuid.UUID, user models.User) error {
	const op = "storage.server.update"
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.ServerHost, s.ServerPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	c := umv1.NewUsersManagerClient(conn)
	_, err = c.Update(ctx, &umv1.UpdateRequest{
		Id:   uid.String(),
		User: profilers.UsrToProroUsr(user),
	})
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Delete implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) Delete(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "storage.server.delete"
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.ServerHost, s.ServerPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	c := umv1.NewUsersManagerClient(conn)
	res, err := c.Delete(ctx, &umv1.DeleteRequest{
		Id: uid.String(),
	})
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user, err := profilers.ProtoUsrToUsr(res.GetUser())
	if err != nil {
		s.log.Error(fmt.Sprintf("%s: failed to convert proto user to model user: %v", op, err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
