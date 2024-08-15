package main

import (
	"auth/api"
	"auth/api/handler"
	"auth/config"
	"auth/genproto/user"
	"auth/logger"
	"auth/service"
	"auth/storage/postgres"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log := logger.NewLogger()

	userRepo := postgres.NewUserRepository(db)
	authService := service.NewUserService(userRepo, log)
	authHandler := handler.NewAuthenticaionHandler(authService,log)

	s := grpc.NewServer()
	user.RegisterUsersServer(s, authService)

	go func() {
		// defer wg.Done()
		fmt.Printf("Server is listening on port %s\n", config.Load().HTTP_PORT)
		auth := api.NewServer(authHandler)
		router := auth.NewRouter()
		if err := router.Run(config.Load().HTTP_PORT); err != nil {
			log.Error("server error", "Error while running HTTP server: %v", err)
		}
	}()

	// gRPC serverni TCP portda tinglash
	go func() {

		liss, err := net.Listen("tcp", "auth-services1"+config.Load().USER_SERVICE)
		if err != nil {
			log.Error("...", "Error while listening on TCP: %v", err)
			return
		}

		// gRPC serverni ishga tushurish
		if err := s.Serve(liss); err != nil {
			log.Error("service not listening", "Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
}

