package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prajnapras19/attacher/api"
	"github.com/prajnapras19/attacher/client/mysql"
	"github.com/prajnapras19/attacher/config"
	"github.com/prajnapras19/attacher/user"
)

func main() {
	cfg := config.Get()
	dbmysql := mysql.NewService(cfg.MySQLConfig)

	// repositories
	userRepository := user.NewRepository(cfg, dbmysql.GetDB())

	// services
	userService := user.NewService(cfg, userRepository)

	// handlers
	handler := api.NewHandler(
		cfg,
		userService,
	)

	// routes
	router := gin.Default()
	if cfg.AllowCORS {
		router.Use(api.CORSMiddleware())
	}

	router.GET("/_health", handler.HealthCheck)

	//apiV1 := router.Group("/api/v1")

	router.Run(fmt.Sprintf(":%d", cfg.RESTPort))
}
