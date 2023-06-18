package api

import (
	"context"
	"gardener/pkg/cursor"
	"gardener/services/users/internal/api/grpc/converter"
	pb "gardener/services/users/internal/api/grpc/interface"
	"gardener/services/users/internal/models/user"
	"gardener/services/users/internal/models/user/profile"

	"github.com/google/uuid"
)

func (s *Server) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	cur, err := cursor.DecodePaginationCursor(in.Cursor)
	if err != nil {
		return nil, err
	}

	res, pagination, err := s.userService.ListUsers(ctx, cur)
	if err != nil {
		return nil, err
	}

	pagingInfo := cursor.MakePaginationInfo(pagination, cur)

	resp := &pb.ListUsersResponse{
		NextCursor: pagingInfo.Next,
		PrevCursor: pagingInfo.Prev,
	}

	for _, user := range res {
		resp.Users = append(resp.Users, converter.FromModel(user))
	}

	return resp, nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	res, err := s.userService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.FromModel(res), nil
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	user := user.NewWithoutId(in.GetEmail(), in.GetPassword(), nil)
	res, err := s.userService.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return converter.FromModel(res[0]), nil
}

func (s *Server) UpdateUserProfile(ctx context.Context, in *pb.UpdateUserProfileRequest) (*pb.User, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}
	profile := profile.New(id, in.Username, in.FirstName, in.LastName)
	res, err := s.userService.UpdateUserProfile(ctx, id, profile)
	if err != nil {
		return nil, err
	}

	return converter.FromModel(res), nil
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.User, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}
	res, err := s.userService.RemoveUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.FromModel(res), nil
}
