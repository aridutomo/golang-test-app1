package api

import (
	"testapp1/api/handler"
	"testapp1/config"
	"testapp1/internal/repository"
	"testapp1/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	config.InitDB()

	articleRepo := repository.NewArticleRepository()
	articleUsecase := usecase.NewArticleUsecase(articleRepo)
	articleHandler := handler.NewArticleHandler(articleUsecase)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/articles", articleHandler.GetAll)

	router.GET("/article/:id", articleHandler.GetByID)

	router.Run(":8080")
}
