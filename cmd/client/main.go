package main

import (
	"context"
	"fibonacci/pkg/api"
	"flag"
	"log"
	"strconv"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalf("not enought arguments")
	}

	start, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	end, err := strconv.ParseInt(flag.Arg(1), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":"+viper.GetString("grpcserver.port"), grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := api.NewFibonacciClient(conn)
	res, err := client.GetSequence(context.Background(), &api.Request{Start: start, End: end})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetResult())

}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
