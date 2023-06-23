package main

import (
	"context"
	"flag"
	"log"

	"enigmacamp.com/calculator-grpc-client/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//buat akses menggunakan flog
	number1 := flag.Int("num1", 0, "Number 1")
	number2 := flag.Int("num2", 1, "Number 1")
	operator := flag.String("opr", "+", "Calculator Operator (+,-,*,/)")
	serverHost := flag.String("srv", "localhost:8888", "Server Host")

	flag.Parse()
	var conn *grpc.ClientConn
	//DEPRECATED: grpc.WithInsecure -> maksudnya tanpa SSL
	conn, err := grpc.Dial(*serverHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect...", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	c := api.NewCalculatorClient(conn)

	response, err := c.DoCalc(context.Background(), &api.CalculatorInputMessage{
		Number1:  int32(*number1),
		Number2:  int32(*number2),
		Operator: *operator,
	})

	if err != nil {
		log.Fatalln("error when calling...", err)
	}
	log.Printf("%v", response.ResNumber)
}
