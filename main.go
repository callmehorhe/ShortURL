package main

import (
	"log"
	"net"

	"github.com/callmehorhe/shorturl/api/pkg/handler"
	"github.com/callmehorhe/shorturl/api/pkg/repository"
	"github.com/callmehorhe/shorturl/api/pkg/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	s := grpc.NewServer()
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	srv := handler.NewHandler(serv)

	handler.RegisterURLServer(s, srv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
