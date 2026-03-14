package initialize

import (
	"fmt"
	"mx-shop-api/order-web/global"
	"mx-shop-api/order-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	s := zap.S()
	consulInfo := global.ServerConfig.ConsulInfo
	orderConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.OrderSrvConf.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		s.Fatal("【InitSrvConn】商品服务连接失败")
	}

	goodsClient := proto.NewOrderClient(orderConn)
	global.OrderSrvClient = goodsClient
}
