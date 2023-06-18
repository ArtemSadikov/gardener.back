package api

import "gardener/services/users/internal/services"

type Server struct {
	userService services.UserService
}

func New(userService services.UserService) *Server {
	return &Server{userService: userService}
}
