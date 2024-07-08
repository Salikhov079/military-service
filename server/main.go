package main

import (
	"log"
	"net"

	"github.com/Military/military-service/config"
	pb "github.com/Military/military-service/genprotos/militaries"
	"github.com/Military/military-service/service"
	"github.com/Military/military-service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	db, err := postgres.DbConn()
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterBulletServiceServer(server, service.NewBulletService(db))
	pb.RegisterFuelServiceServer(server, service.NewFuelService(db))
	pb.RegisterTechniqueServiceServer(server, service.NewTechniqueService(db))

	log.Printf("Server started on port: %v", cfg.HTTPPort)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
