package main

import (
	"context"
	"log"
	"net"

	api "github.com/danik-progger/go_mini_projects/grpc_server/api"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

type TodoServer struct {
	api.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *api.NewTodo) (*api.Todo, error) {
	log.Printf("Received: %v", in.GetName())
	todo := &api.Todo{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	return todo, nil

}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	api.RegisterTodoServiceServer(s, &TodoServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
