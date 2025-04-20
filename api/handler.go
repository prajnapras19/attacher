package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prajnapras19/attacher/config"
	"github.com/prajnapras19/attacher/constants"
	"github.com/prajnapras19/attacher/user"
)

type Handler interface {
	HealthCheck(*gin.Context)
	GetLoginPage(*gin.Context)
	DoLogin(*gin.Context)
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

func (h *handler) GetLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, constants.LoginPage, nil)
}

func (h *handler) DoLogin(c *gin.Context) {
	loginResponse, err := h.userService.Login(
		&user.LoginRequest{
			Username: c.PostForm(constants.Username),
			Password: c.PostForm(constants.Password),
		},
	)
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.LoginPage, gin.H{
			constants.Error: err.Error(),
		})
		return
	}
	c.SetCookie(constants.Token, loginResponse.Token, 3600, "", h.cfg.CookieDomain, false, false)
	c.Redirect(http.StatusFound, "")
}
