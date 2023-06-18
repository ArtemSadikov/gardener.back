package converter

import (
	pb "gardener/services/users/internal/api/grpc/interface"
	"gardener/services/users/internal/models/user"
)

func FromModel(data user.User) *pb.User {
	result := &pb.User{
		Id: data.ID.String(),
		Personal: &pb.Personal{
			Email: data.Email,
		},
	}

	if data.Profile != nil {
		result.Profile = &pb.Profile{
			Username:  data.Profile.Username,
			FirstName: data.Profile.FirstName,
			LastName:  data.Profile.LastName,
		}
	}

	return result
}
