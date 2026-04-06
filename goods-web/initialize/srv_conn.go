package initialize

import (
	"fmt"
	"mx-shop-api/goods-web/global"
	"mx-shop-api/goods-web/proto"

	"mx-shop-api/goods-web/utils/otgrpc"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	s := zap.S()
	consulInfo := global.ServerConfig.ConsulInfo
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvConf.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		s.Fatal("【InitSrvConn】商品服务连接失败")
	}

	goodsClient := proto.NewGoodsClient(goodsConn)
	global.GoodSrvClient = goodsClient
}
