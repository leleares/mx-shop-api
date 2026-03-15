package order

import (
	"context"
	"mx-shop-api/order-web/api"
	"mx-shop-api/order-web/forms"
	"mx-shop-api/order-web/global"
	"mx-shop-api/order-web/models"
	"mx-shop-api/order-web/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(ctx *gin.Context) {
	s := zap.S()
	var form forms.CreateOrderForm
	uid, _ := ctx.Get("userId")
	err := ctx.ShouldBind(&form)
	if err != nil {
		s.Errorln("【Create】表单验证失败", err)
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	_, err = global.OrderSrvClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  int32(uid.(uint)),
		Address: form.Address,
		Name:    form.Name,
		Mobile:  form.Mobile,
		Post:    form.Post,
	})
	if err != nil {
		s.Errorln("【Create】创建订单失败", err)
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	// TODO：用户下单完成一般需要跳转到支付宝的支付页面，所以这里需要构造跳转到支付宝页面的link
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func Detail(ctx *gin.Context) {
	s := zap.S()
	idStr := ctx.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	uid, _ := ctx.Get("userId")

	resp, err := global.OrderSrvClient.OrderDetail(context.Background(), &proto.OrderRequest{
		Id:     int32(idInt),
		UserId: int32(uid.(uint)),
	})
	if err != nil {
		s.Errorln("【Detail】获取订单详情失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	models.ToStringLog(resp)

	/*
		{
			msg:"ok",
			"order_info":{
				"id":       order.Id,
				"user_id":  order.UserId,
				"order_sn": order.OrderSn,
				"pay_type": order.PayType,
				"status":   order.Status,
				"post":     order.Post,
				"total":    order.Total,
				"address":  order.Address,
				"name":     order.Name,
				"mobile":   order.Mobile,
				"add_time": order.AddTime,
			},
			good_list:[
				{
					id:xx,
					order_id: xx,
					goods_id:xx,
					goods_name:xx,
					goods_img:xx,
					goods_price:xx,
					num:xx,
				}
			]
		}
	*/

	orderInfo := make(map[string]interface{})
	goodList := make([]interface{}, 0)

	orderInfo["id"] = resp.OrderInfo.Id
	orderInfo["user_id"] = resp.OrderInfo.UserId
	orderInfo["order_sn"] = resp.OrderInfo.OrderSn
	orderInfo["pay_type"] = resp.OrderInfo.PayType
	orderInfo["status"] = resp.OrderInfo.Status
	orderInfo["post"] = resp.OrderInfo.Post
	orderInfo["total"] = resp.OrderInfo.Total
	orderInfo["address"] = resp.OrderInfo.Address
	orderInfo["name"] = resp.OrderInfo.Name
	orderInfo["mobile"] = resp.OrderInfo.Mobile
	orderInfo["add_time"] = resp.OrderInfo.AddTime

	for _, orderGood := range resp.Goods {
		goodList = append(goodList, map[string]interface{}{
			"id":          orderGood.Id,
			"order_id":    orderGood.OrderId,
			"goods_id":    orderGood.GoodsId,
			"goods_name":  orderGood.GoodsName,
			"goods_img":   orderGood.GoodsImage,
			"goods_price": orderGood.GoodsPrice,
			"num":         orderGood.Nums,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":        "ok",
		"order_info": orderInfo,
		"good_list":  goodList,
	})
}

func List(ctx *gin.Context) {
	uid, _ := ctx.Get("userId")
	pnStr := ctx.DefaultQuery("pn", "1")
	pn, _ := strconv.Atoi(pnStr)
	pSizeStr := ctx.DefaultQuery("pSize", "10")
	pSize, _ := strconv.Atoi(pSizeStr)

	resp, err := global.OrderSrvClient.OrderList(context.Background(), &proto.OrderFilterRequest{
		UserId:      int32(uid.(uint)),
		Pages:       int32(pn),
		PagePerNums: int32(pSize),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	finalDataSlice := make([]interface{}, 0)

	for _, order := range resp.Data {
		finalDataSlice = append(finalDataSlice, map[string]interface{}{
			"id":       order.Id,
			"user_id":  order.UserId,
			"order_sn": order.OrderSn,
			"pay_type": order.PayType,
			"status":   order.Status,
			"post":     order.Post,
			"total":    order.Total,
			"address":  order.Address,
			"name":     order.Name,
			"mobile":   order.Mobile,
			"add_time": order.AddTime,
		})
	}

	/*
		resp:
		{
			total:1,
			data:[
				{
					id:1,
					user_id:1,
					order_sn:1,
					pay_type:x,
					status:x,
					post:xx,
					total:1.1,
					address:xx,
					name:xx,
					mobile:11,
					add_time:11
				}
			]
		}
	*/

	ctx.JSON(http.StatusOK, gin.H{
		"total": resp.Total,
		"data":  finalDataSlice,
	})
}
