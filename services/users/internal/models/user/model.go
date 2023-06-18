package user

import (
	"gardener/services/users/internal/models/user/profile"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey;default:uuid_generate_v4();type:uuid;"`
	Email     string    `gorm:"unique;type:character varying(255);not null;"`
	Password  string
	Profile   *profile.Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:id;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func New(
	id uuid.UUID,
	email string,
	password string,
	profile *profile.Profile,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) User {
	result := User{
		ID:        id,
		Email:     email,
		Password:  password,
		Profile:   profile,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if !deletedAt.IsZero() {
		result.DeletedAt = gorm.DeletedAt{Time: deletedAt}
	}

	return result
}

func NewWithoutId(
	email string,
	password string,
	profile *profile.Profile,
) User {
	return User{
		Email:    email,
		Password: password,
		Profile:  profile,
	}
}
