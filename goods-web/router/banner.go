package router

import (
	"mx-shop-api/goods-web/api/banner"
	"mx-shop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerGroup := router.Group("banner").Use(middlewares.Trace())
	{
		BannerGroup.POST("/create", banner.CreateBanner)
		BannerGroup.DELETE("/delete/:id", banner.DeleteBanner)
		BannerGroup.PUT("/update/:id", banner.UpdateBanner)
		BannerGroup.GET("/list", banner.BannerList)
	}
}
