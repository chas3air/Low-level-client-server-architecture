package usersmanager

import (
	"context"
	"server/internal/domain/interfaces"
	"server/internal/domain/profiles"

	umv1 "github.com/chas3air/protos/gen/go/usersManager"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	umv1.UnimplementedUsersManagerServer
	usersManager interfaces.UsersManager
}

func New(usersManager interfaces.UsersManager) *serverAPI {
	return &serverAPI{
		usersManager: usersManager,
	}
}

func Register(grpc *grpc.Server, usersManager interfaces.UsersManager) {
	umv1.RegisterUsersManagerServer(grpc, &serverAPI{usersManager: usersManager})
}

func (s *serverAPI) GetUsers(ctx context.Context, in *umv1.GetUsersRequest) (*umv1.GetUsersResponse, error) {
	users, err := s.usersManager.GetUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve users")
	}

	usersForResp := make([]*umv1.User, len(users))
	for i, user := range users {
		profileUser, err := profiles.UsrToProroUsr(user)
		if err != nil {
			continue
		}
		usersForResp[i] = profileUser
	}

	return &umv1.GetUsersResponse{
		Users: usersForResp,
	}, nil
}

func (s *serverAPI) GetUserById(ctx context.Context, in *umv1.GetUserByIdRequest) (*umv1.GetUserByIdResponse, error) {
	userID := in.GetId()
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id must be uuid")
	}

	user, err := s.usersManager.GetUserById(ctx, parsedUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve user by id")
	}

	userForResp, err := profiles.UsrToProroUsr(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to convert user")
	}

	return &umv1.GetUserByIdResponse{
		User: userForResp,
	}, nil
}

func (s *serverAPI) GetUserByEmail(ctx context.Context, in *umv1.GetUserByEmailRequest) (*umv1.GetUserByEmailResponse, error) {
	email := in.GetEmail()
	if email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	user, err := s.usersManager.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve user by email")
	}

	userForResp, err := profiles.UsrToProroUsr(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to convert user")
	}

	return &umv1.GetUserByEmailResponse{
		User: userForResp,
	}, nil
}

func (s *serverAPI) Insert(ctx context.Context, in *umv1.InsertRequest) (*umv1.InsertResponse, error) {
	protoUser := in.GetUser()
	if protoUser == nil {
		return nil, status.Error(codes.InvalidArgument, "user is required")
	}

	user, err := profiles.ProtoUsrToUsr(protoUser)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to convert user")
	}

	err = s.usersManager.Insert(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to insert user")
	}

	return nil, nil
}

func (s *serverAPI) Update(ctx context.Context, in *umv1.UpdateRequest) (*umv1.UpdateResponse, error) {
	userID := in.GetId()
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id must be uuid")
	}

	protoUser := in.GetUser()
	if protoUser == nil {
		return nil, status.Error(codes.InvalidArgument, "user is required")
	}

	user, err := profiles.ProtoUsrToUsr(protoUser)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to convert user")
	}

	err = s.usersManager.Update(ctx, parsedUUID, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return nil, nil
}

func (s *serverAPI) Delete(ctx context.Context, in *umv1.DeleteRequest) (*umv1.DeleteResponse, error) {
	userID := in.GetId()
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user_id must be uuid")
	}

	user, err := s.usersManager.Delete(ctx, parsedUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	userForResp, err := profiles.UsrToProroUsr(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to convert user")
	}

	return &umv1.DeleteResponse{
		User: userForResp,
	}, nil
}
