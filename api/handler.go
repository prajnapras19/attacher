package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prajnapras19/attacher/config"
	"github.com/prajnapras19/attacher/user"
)

type Handler interface {
	HealthCheck(*gin.Context)
}

type handler struct {
	cfg         *config.Config
	userService user.Service
}

func (h *handler) HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func NewHandler(
	cfg *config.Config,
	userService user.Service,
) Handler {
	return &handler{
		cfg:         cfg,
		userService: userService,
	}
}
