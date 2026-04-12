package initialize

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"go.uber.org/zap"
)

const resName = "goods-web-goods-list"

func InitSentinel() {
	// 初始化sentinel
	conf := config.NewDefaultConfig()
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		zap.S().Fatalf("初始化sentinel失败: %v", err)
	}

	// 配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               resName,     // 资源名，即规则的作用目标
			TokenCalculateStrategy: flow.Direct, // 当前流量控制器的Token计算策略，Direct表示直接使用字段 Threshold 作为阈值；WarmUp表示使用预热方式计算Token的阈值。
			ControlBehavior:        flow.Reject, // 超过阈值直接拒绝
			Threshold:              10,          // 流控阈值
			StatIntervalInMs:       1000,        // 规则对应的流量控制器的独立统计结构的统计周期。如果StatlntervallnMs是1000，也就是统计QPS。
		},
	})
	if err != nil {
		zap.S().Fatalf("配置限流规则失败: %v", err)
		return
	}
}
