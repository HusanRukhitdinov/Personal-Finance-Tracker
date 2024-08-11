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
		fmt.Printf("Server is listening on port %s\n", config.Load().USER_SERVICE)
		auth := api.NewServer(authHandler)
		router := auth.NewRouter()
		if err := router.Run(config.Load().USER_SERVICE); err != nil {
			log.Error("server error", "Error while running HTTP server: %v", err)
		}
	}()

	// gRPC serverni TCP portda tinglash
	go func() {

		liss, err := net.Listen("tcp", "localhost"+config.Load().HTTP_PORT)
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
	// Goroutine'larni kutish
	// wg.Wait()
}

// 	go func() {
// 		fmt.Printf("Server is listening on port %s\n", config.Load().Server.USER_PORT)
// 		auth := api.NewServer(authHandler)
// 		router := auth.NewRouter()
// 		if err := router.Run("localhost:8080"); err != nil {
// 			log.Error("server error", "Error while running HTTP server: %v", err)
// 		}
// 	}()

// 	go func() {

// 		liss, err := net.Listen("tcp", "localhost"+config.Load().Postgres.DB_PORT)
// 		if err != nil {
// 			log.Error("...", "Error while listening on TCP: %v", err)
// 			return
// 		}

// 		if err := s.Serve(liss); err != nil {
// 			log.Error("service not listening", "Failed to serve: %v", err)
// 		}
// 	}()

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// }
