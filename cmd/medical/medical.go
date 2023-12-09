package main

import (
	"awesomeProject/internal/repository"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"awesomeProject/internal/api/medical"
	medicalGrpc "awesomeProject/pkg/generated/medical"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("Start app")

	db, err := sqlx.Open("postgres", "user=postgres password=password dbname=test sslmode=disable")

	if err != nil {
		log.Fatalf("fail to connecto to db: %v", err)
	}

	medicalRepo := repository.NewMedicalRepo(db)

	medicalApi := medical.NewImplementation(medicalRepo)

	lis, err := net.Listen("tcp", ":8082")

	if err != nil {
		log.Fatalf("fail to listen port: %v", err)
	}

	var medicalServer = grpc.NewServer()

	reflection.Register(medicalServer)
	medicalGrpc.RegisterMedicalServiceServer(medicalServer, medicalApi)

	log.Info("App started")

	if err = medicalServer.Serve(lis); err != nil {
		log.Fatalf("serving error %v", err)
	}
}
