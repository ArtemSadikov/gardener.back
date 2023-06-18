package user

import "gorm.io/gorm"

type Service struct {
	repo *gorm.DB
}

func New(repo *gorm.DB) *Service {
	return &Service{repo: repo}
}
