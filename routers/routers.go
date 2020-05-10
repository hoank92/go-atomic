package routers

import (
	"github.com/gin-gonic/gin"
	"go-atomic/main/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("products", controllers.GetAllProduct)
		v1.GET("products/10", controllers.GetProductById)
		v1.GET("updateProduct", controllers.UpdateProduct)
	}

	return r
}
