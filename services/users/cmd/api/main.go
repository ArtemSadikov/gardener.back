package main

import (
	"fmt"
	api "gardener/services/users/internal/api/grpc"
	user "gardener/services/users/internal/api/grpc/interface"
	"gardener/services/users/internal/infrastructure/container"
	"gardener/services/users/internal/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	container, err := container.New()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(func(userServices services.UserService) *api.Server {
		return api.New(userServices)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(server *api.Server) error {
		var errCh chan error

		listener, err := net.Listen("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
			return err
		}

		gServer := grpc.NewServer()
		user.RegisterUserServiceServer(gServer, server)

		go func() {
			fmt.Println("Listening on 3000")
			errCh <- gServer.Serve(listener)
		}()

		if err := <-errCh; err != nil {
			log.Fatal(err)
		}

		return nil
	})
}
