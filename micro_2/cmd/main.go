package main

import (
	"golang-application/internal/config"
	"golang-application/internal/handler"
	"golang-application/internal/service"
	userpb "golang-application/proto"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	config.LoadEnv()
	router := gin.Default()
	userconn, err := grpc.NewClient(os.Getenv("USER_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to doctor service: %v", err)
	}
	userClient := userpb.NewUserServiceClient(userconn)

	service := service.NewUserService(userClient)
	handler := handler.NewUserHandler(service)

	user := router.Group("/user")
	config.UserRouter(user, handler)
	port := os.Getenv("PORT")
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
