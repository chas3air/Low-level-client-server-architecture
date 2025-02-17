package usersmanager

import (
	"context"
	"log/slog"
	"server/internal/domain/interfaces"

	umv1 "github.com/chas3air/protos/gen/go/usersManager"
	"google.golang.org/grpc"
)

type serverAPI struct {
	umv1.UnimplementedUsersManagerServer

	log          *slog.Logger
	usersManager interfaces.UsersManager
}

func New(log *slog.Logger, usersManager interfaces.UsersManager) *serverAPI {
	return &serverAPI{
		log:          log,
		usersManager: usersManager,
	}
}

func Register(grpc *grpc.Server, usersManager interfaces.UsersManager) {
	umv1.RegisterUsersManagerServer(grpc, &serverAPI{usersManager: usersManager})
}

func (s *serverAPI) GetUsers(ctx context.Context, in *umv1.GetUsersRequest) (*umv1.GetUsersResponse, error) {
	return nil, nil
}

func (s *serverAPI) GetUserById(ctx context.Context, in *umv1.GetUserByIdRequest) (*umv1.GetUserByIdResponse, error) {
	return nil, nil
}

func (s *serverAPI) GetUserByEmail(ctx context.Context, in *umv1.GetUserByEmailRequest) (*umv1.GetUserByEmailResponse, error) {
	return nil, nil
}

func (s *serverAPI) Insert(ctx context.Context, in *umv1.InsertRequest) (*umv1.InsertResponse, error) {
	return nil, nil
}

func (s *serverAPI) Update(ctx context.Context, in *umv1.UpdateRequest) (*umv1.UpdateResponse, error) {
	return nil, nil
}

func (s *serverAPI) Delete(ctx context.Context, in *umv1.DeleteRequest) (*umv1.DeleteResponse, error) {
	return nil, nil
}
