package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prajnapras19/attacher/api"
	"github.com/prajnapras19/attacher/attachment"
	"github.com/prajnapras19/attacher/client/mysql"
	"github.com/prajnapras19/attacher/config"
	"github.com/prajnapras19/attacher/user"
)

func main() {
	cfg := config.Get()
	dbmysql := mysql.NewService(cfg.MySQLConfig)

	// repositories
	userRepository := user.NewRepository(cfg, dbmysql.GetDB())
	attachmentRepository := attachment.NewRepository(cfg, dbmysql.GetDB())

	// services
	userService := user.NewService(cfg, userRepository)
	attachmentService := attachment.NewService(cfg, attachmentRepository)

	// handlers
	handler := api.NewHandler(
		cfg,
		userService,
		attachmentService,
	)

	// routes
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	if cfg.AllowCORS {
		router.Use(api.CORSMiddleware())
	}

	router.GET("/_health", handler.HealthCheck)
	router.GET("/login", handler.GetLoginPage)
	router.POST("/login", handler.DoLogin)

	adminRouter := router.Group("/admin")
	adminRouter.Use(api.JWTSystenTokenMiddleware(userService))
	adminRouter.GET("/upsert-user", handler.GetUpsertUserWithFilePage)
	adminRouter.POST("/upsert-user", handler.UpsertUserWithFile)

	router.Use(api.JWTTokenMiddleware(userService))
	router.GET("", handler.ListActiveFiles)
	router.GET("/attachments/:serial", handler.DownloadAttachment)

	router.Run(fmt.Sprintf(":%d", cfg.RESTPort))
}
