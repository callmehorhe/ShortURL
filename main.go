package main

import (
	"log"
	"net"
	"os"

	"github.com/callmehorhe/shorturl/api/pkg/handler"
	"github.com/callmehorhe/shorturl/api/pkg/repository"
	"github.com/callmehorhe/shorturl/api/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	srv := handler.NewHandler(serv)

	handler.RegisterURLServer(s, srv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Listen port 8080")
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
