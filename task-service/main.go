package main

import (
	"context"
	"log"
	"net"

	pb "gen/go/proto/task-service/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTaskServiceServer
}

func (s *server) TaskInfo(ctx context.Context, req *pb.TaskInfoRequest) (*pb.TaskInfoResponse, error) {
	if req.TaskId == "" {
		return nil, status.Error(codes.InvalidArgument, "task_id is required")
	}

	if _, err := uuid.Parse(req.TaskId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "task_id must be a valid UUID")
	}

	return &pb.TaskInfoResponse{
		TaskId:      req.TaskId,
		Name:        "Test Task",
		Description: "This is a test task",
		Status:      "IN_PROGRESS",
	}, nil
}

func (s *server) TaskCreate(ctx context.Context, req *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	taskID := uuid.New().String()

	return &pb.TaskCreateResponse{
		TaskId: taskID,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, &server{})

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
