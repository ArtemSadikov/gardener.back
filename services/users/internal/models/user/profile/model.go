package profile

import "github.com/google/uuid"

type Profile struct {
	UserID    uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Username  string    `gorm:"type:character varying(255);"`
	FirstName string
	LastName  string
}

func (p Profile) TableName() string {
	return "user_profile"
}

func New(
	userId uuid.UUID,
	username string,
	firstName string,
	lastName string,
) Profile {
	return Profile{
		UserID:    userId,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
	}
}
