package main

import (
	"context"
	"log"
	"os"
	"user-manages/config"
	userRepo "user-manages/modules/user/userRepository"
	"user-manages/pkg/database"
	"user-manages/server"
)

func main() {
	ctx := context.Background()

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	// Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	userRepo := userRepo.NewUserRepository(db)

	go userRepo.LogUserCount(ctx)

	server.Start(ctx, &cfg, db)
}
