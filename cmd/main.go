package main

import (
	"fibonacci"
	"log"
	"net"

	"fibonacci/api/proto"
	handler "fibonacci/pkg/handler/http"
	"fibonacci/pkg/repository"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "",
		DB:       0,
	})

	repos := repository.NewRepository(rdb)
	handler := handler.NewHandler(repos)

	srv := new(fibonacci.HTTPServer)

	s := grpc.NewServer()
	grpcSrv := proto.NewGRPCServer(repos)
	proto.RegisterFibonacciServer(s, grpcSrv)

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
