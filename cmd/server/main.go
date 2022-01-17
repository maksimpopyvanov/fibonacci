package main

import (
	"log"
	"net"

	"fibonacci/pkg/api"
	handler "fibonacci/pkg/handler/http"
	"fibonacci/pkg/repository"
	"fibonacci/pkg/server"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	rdb, err := repository.NewRedisClient(repository.Config{
		Addr:     viper.GetString("redis.addr"),
		DB:       0,
		Password: "",
	})

	//Проверка на доступность кеша
	if err != nil {
		log.Fatalf("failed to initialize cache: %s", err.Error())
	}

	repos := repository.NewRepository(rdb) //Слой для работы с кешем
	handler := handler.NewHandler(repos)

	srv := new(server.HTTPServer)

	s := grpc.NewServer()
	grpcSrv := server.NewGRPCServer(repos)
	api.RegisterFibonacciServer(s, grpcSrv)

	go func() {
		l, err := net.Listen("tcp", ":"+viper.GetString("grpcserver.port"))

		if err != nil {
			log.Fatalf("error occured while running grpc server: %s", err.Error())
		}

		if err := s.Serve(l); err != nil {
			log.Fatalf("error occured while running grpc server: %s", err.Error())
		}
	}()

	if err := srv.Run(viper.GetString("httpserver.port"), handler.Routes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
