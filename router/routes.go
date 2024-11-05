package router

import (
	"net/http"

	"github.com/WesleyYuri/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
  v1 := router.Group("/api/v1")
  {
    v1.GET("ping", func(ctx *gin.Context){
      ctx.JSON(http.StatusOK, gin.H{
        "message" : "pong",
      })
    })

    v1.GET("/opening", handler.ShowOpeningHandler)
    v1.POST("/opening", handler.CreateOpeningHandler)
    v1.DELETE("/opening", handler.DeleteOpeningHandler)
    v1.PUT("/opening", handler.UpdateOpeningHandler)
    v1.GET("/openings", handler.ListOpeningsHandler)
  }
}

