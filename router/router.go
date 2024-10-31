package router

import (
  "github.com/gin-gonic/gin"
)

func Initialize() {
  // Inicitalize Router
  r := gin.Default()

  // Initialize Routes
  initializeRoutes(r)

  // Run the server
	r.Run()
}
