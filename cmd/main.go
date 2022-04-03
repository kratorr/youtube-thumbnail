package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"youtube-thumbnail/internal/handler"
	"youtube-thumbnail/internal/repository"
	"youtube-thumbnail/internal/service"

	"google.golang.org/grpc"
)

type Configuration struct {
	Ip   string
	Port int
}

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	addr := net.JoinHostPort(configuration.Ip, strconv.Itoa(configuration.Port))

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicln("Failed to listen port", err)
	}

	server := grpc.NewServer()
	mu := &sync.RWMutex{}
	cache := make(map[string][]byte)
	repos := repository.NewRepository(cache, mu)
	service := service.NewService(repos)
	handler.NewHandler(server, service)
	log.Println("Start server on %s", addr)
	err = server.Serve(listen)
	if err != nil {
		log.Println("Failed to start server", err)
	}
}
