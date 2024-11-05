package main

import (
	"github.com/WesleyYuri/gopportunities/config"
	"github.com/WesleyYuri/gopportunities/router"
)

var (
  logger *config.Logger
)

func main () {
  logger = config.GetLogger("main")
  // Initialize Configs
  err := config.Init()
  if err != nil {
    logger.Errf("config initialization error: %v", err)
    return
  }

  // Initialize Router
  router.Initialize()
}
