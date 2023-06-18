package services

import (
	"context"
	"gardener/pkg/cursor"
	"gardener/services/users/internal/models/user"
	"gardener/services/users/internal/models/user/profile"

	"github.com/google/uuid"
)

type UserService interface {
	WriteUserService
	ReadUserService
}

type WriteUserService interface {
	CreateUser(ctx context.Context, users ...user.User) ([]user.User, error)
	UpdateUserProfile(ctx context.Context, userId uuid.UUID, data profile.Profile) (user.User, error)
	RemoveUser(ctx context.Context, userId uuid.UUID) (user.User, error)
}

type ReadUserService interface {
	GetUserByID(ctx context.Context, userId uuid.UUID) (user.User, error)
	ListUsers(ctx context.Context, paging cursor.PaginationCursor) ([]user.User, cursor.PaginationCursor, error)
}
