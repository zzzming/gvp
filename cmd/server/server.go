package main

import (
	"context"
	"net"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	pb "github.com/zzzming/gvp/pkg/pinecone/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	PineCodeAPIKey = "api-key"
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

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Extract authentication metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing credentials")
	}
	apiKey := md.Get(PineCodeAPIKey)
	if len(apiKey) == 0 || apiKey[0] == "" {
		return nil, status.Errorf(codes.Unauthenticated, "missing api-key")
	}

	// Continue with the handler if authentication is successful
	return handler(ctx, req)
}

func GRPCServer(port string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	pb.RegisterVectorServiceServer(s, &server{})
	reflection.Register(s)
	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	GRPCServer("50051")
}
