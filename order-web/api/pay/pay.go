package pay

import (
	"context"
	"mx-shop-api/order-web/global"
	"mx-shop-api/order-web/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
)

func Notify(ctx *gin.Context) {
	// 1. 确认支付宝身份
	s := zap.S()
	appID := global.ServerConfig.AlipayInfo.AppID
	privateKey := global.ServerConfig.AlipayInfo.PrivateKey
	aliPublicKey := global.ServerConfig.AlipayInfo.AliPublicKey

	client, err := alipay.New(appID, privateKey, false)
	if err != nil {
		s.Errorln("【Create】生成支付宝link失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "支付宝回调失败"})
		return
	}

	err = client.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		s.Errorln("【Create】生成支付宝link失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "支付宝回调失败"})
		return
	}

	notify, err := client.GetTradeNotification(ctx.Request)
	if err != nil {
		s.Errorln("【Notify】获取支付宝回调通知失败", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的支付宝回调"})
		return
	}

	// 2. 更新订单状态
	_, err = global.OrderSrvClient.UpdateOrderStatus(context.Background(), &proto.OrderStatus{
		OrderSn: notify.OutTradeNo,
		Status:  string(notify.TradeStatus),
	})

	if err != nil {
		s.Errorln("【Notify】更新订单状态失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新订单状态失败"})
		return
	}

	// 3. 通知支付宝已成功完成逻辑处理
	client.ACKNotification(ctx.Writer)
}
