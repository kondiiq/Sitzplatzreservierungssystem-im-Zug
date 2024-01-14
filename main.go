package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	docs "zugSystem/docs"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//trackRouter := router.Group("/route")
	//{
	//	trackRouter.GET("/")
	//	trackRouter.POST("/")
	//	trackRouter.GET("/:id")
	//	trackRouter.PATCH("/:id")
	//	trackRouter.DELETE("/:id")
	//}
	//
	//trainRouter := router.Group("/train")
	//{
	//	trainRouter.GET("/")
	//	trainRouter.POST("/")
	//	trainRouter.GET("/:id")
	//	trainRouter.PATCH("/:id")
	//	trainRouter.DELETE("/:id")
	//}
	//
	//ticketRouter := router.Group("/ticket")
	//{
	//	ticketRouter.GET("/")
	//	ticketRouter.POST("/")
	//	ticketRouter.GET("/:id")
	//	ticketRouter.PATCH("/:id")
	//	ticketRouter.DELETE("/:id")
	//}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/")
		userRouter.POST("/")
		userRouter.GET("/:id")
		userRouter.PATCH("/:id")
		userRouter.DELETE("/:id")
	}

	router.Run(":8080")
}
