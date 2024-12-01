package config

import (
	"fmt"
	"golang-application/internal/handler"
	"golang-application/internal/model"
	"golang-application/internal/repo"
	userpb "golang-application/proto"

	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

}

func DbConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("PSQL_URL")))

	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connect to Psql")

	db.AutoMigrate(&model.User{})
	return db
}
func GrpcSetup() {
	grpcServer := grpc.NewServer()
	db := DbConn()
	repo := repo.NewUserRepository(db)
	handler := handler.NewUserHandler(repo)

	userpb.RegisterUserServiceServer(grpcServer, handler)

	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	log.Printf("Server listening on port %s", port)
	grpcServer.Serve(listener)
}
