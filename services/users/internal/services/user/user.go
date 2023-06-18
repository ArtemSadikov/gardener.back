package user

import (
	"context"
	"fmt"
	"gardener/pkg/cursor"
	"gardener/services/users/internal/models/user"
	"gardener/services/users/internal/models/user/profile"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Service) CreateUser(ctx context.Context, users ...user.User) ([]user.User, error) {
	if err := s.repo.Save(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) UpdateUserProfile(ctx context.Context, userId uuid.UUID, data profile.Profile) (user.User, error) {
	var model user.User
	if err := s.repo.First(&model, &user.User{ID: userId}).Error; err != nil {
		return user.User{}, err
	}

	model.Profile = &data

	if err := s.repo.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Save(&model).Error; err != nil {
		return user.User{}, err
	}
	return model, nil
}

func (s *Service) RemoveUser(ctx context.Context, userId uuid.UUID) (user.User, error) {
	user := user.User{ID: userId}
	if err := s.repo.Delete(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, userId uuid.UUID) (user.User, error) {
	user := user.User{ID: userId}
	if err := s.repo.Preload(clause.Associations).Find(&user).Error; err != nil {
		return user, err
	}
	fmt.Println(user)

	return user, nil
}

func (s *Service) ListUsers(ctx context.Context, paging cursor.PaginationCursor) ([]user.User, cursor.PaginationCursor, error) {
	var (
		users []user.User
		total int64
	)

	err := s.repo.Model(&user.User{}).
		Count(&total).
		Limit(int(paging.Limit)).
		Offset(int(paging.Offset)).
		Find(&users).Error
	if err != nil{
		return nil, cursor.PaginationCursor{}, err
	}

	return users, paging.Next(total), nil
}
