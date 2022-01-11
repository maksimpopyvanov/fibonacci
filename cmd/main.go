package main

import (
	"fibonacci"
	"log"

	handler "fibonacci/pkg/handler/http"
	"fibonacci/pkg/repository"
	"fibonacci/pkg/service"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
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
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(fibonacci.Server)
	if err := srv.Run(viper.GetString("port"), handler.Routes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
