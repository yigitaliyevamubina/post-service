package main

import (
	"google.golang.org/grpc"
	"net"
	config "template-post-service/config"
	pb "template-post-service/genproto/post_service"
	"template-post-service/pkg/db"
	"template-post-service/pkg/logger"
	"template-post-service/service"
	grpcCLient "template-post-service/service/grpc_client"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-post-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, _, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sql connection to postgres error", logger.Error(err))
	}

	client, err := grpcCLient.New(cfg)
	if err != nil {
		log.Fatal("error while adding grpc client", logger.Error(err))
	}

	postService := service.NewPostService(connDB, log, client)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("failed to listen to: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server is running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("error while listening: %v", logger.Error(err))
	}
}
