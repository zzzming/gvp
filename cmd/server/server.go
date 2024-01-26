package server

import (
	"context"
	"log"
	"net"

	pb "github.com/zzzming/gvp/pkg/pinecone/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement your service.
type server struct {
	pb.UnimplementedVectorServiceServer
}

// Implement your service methods here.
func (s *server) Upsert(ctx context.Context, in *pb.UpsertRequest) (*pb.UpsertResponse, error) {
	return &pb.UpsertResponse{UpsertedCount: 1}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}

func (s *server) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchResponse, error) {
	return &pb.FetchResponse{}, nil
}

func (s *server) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	return &pb.QueryResponse{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return &pb.UpdateResponse{}, nil
}

func (s *server) DescribeIndexStats(ctx context.Context, in *pb.DescribeIndexStatsRequest) (*pb.DescribeIndexStatsResponse, error) {
	return &pb.DescribeIndexStatsResponse{}, nil
}

func GRPCServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterVectorServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}