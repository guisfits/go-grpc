package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/guisfits/go-grpc/internal/database"
	"github.com/guisfits/go-grpc/internal/pb"
	"github.com/guisfits/go-grpc/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
