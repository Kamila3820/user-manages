package server

import (
	"log"
	userHandler "user-manages/modules/user/userHandler"
	userPb "user-manages/modules/user/userPb"
	userRepository "user-manages/modules/user/userRepository"
	userUsecase "user-manages/modules/user/userUsecase"
	"user-manages/pkg/grpccon"
)

func (s *server) userService() {
	repo := userRepository.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repo)
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.UserUrl)

		userPb.RegisterUserGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("User gRPC server listening on %s", s.cfg.Grpc.UserUrl)
		grpcServer.Serve(lis)
	}()

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)

	user.POST("/user/register", httpHandler.CreateUser)
	user.GET("/user/:user_id", httpHandler.FindOneUserProfile)
}
