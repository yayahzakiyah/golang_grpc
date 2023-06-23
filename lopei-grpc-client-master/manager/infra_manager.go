package manager

import (
	"fmt"
	"log"
	"lopei-grpc-client/config"
	"lopei-grpc-client/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InfraManager interface {
	LogicClientConn() service.LopeiPaymentClient
}

type infraManager struct {
	lopeiClient service.LopeiPaymentClient
	cfg         config.Config
}

func (i *infraManager) LogicClientConn() service.LopeiPaymentClient {
	return i.lopeiClient
}

func (i *infraManager) initGrpcClient() {
	dial, err := grpc.Dial(i.cfg.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect...", err)
	}

	client := service.NewLopeiPaymentClient(dial)
	i.lopeiClient = client
	fmt.Println("GRPC client connected....")
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initGrpcClient()
	return &infra
}
