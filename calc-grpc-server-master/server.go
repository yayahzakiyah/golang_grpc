package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"enigmacamp.com/calculator-grpc-server/api"
	"enigmacamp.com/calculator-grpc-server/calc"
	"google.golang.org/grpc"
)

func main() {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")

	serverInfo := fmt.Sprintf("%s:%s", host, port)
	
	listen, err := net.Listen("tcp", serverInfo)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}

	s := calc.Server{}
	grpcServer :=  grpc.NewServer()
	
	api.RegisterCalculatorServer(grpcServer, &s)
	log.Println("Server run on", serverInfo)
	
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln("Failed to server..", err)
	}

}