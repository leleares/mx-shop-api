package router

import (
	"mx-shop-api/goods-web/api/category"
	"mx-shop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	GoodsGroup := router.Group("category").Use(middlewares.Trace())
	{
		GoodsGroup.GET("/list", category.CategoryList)
		GoodsGroup.POST("/create", category.CreateCategory) //  middlewares.JWTAuth(), middlewares.IsAdmin(),
		GoodsGroup.DELETE("/delete/:id", category.DeleteCategory)
		GoodsGroup.PUT("/update/:id", category.UpdateCategory)
		GoodsGroup.GET("/subCategory", category.SubCategory)
	}
}
