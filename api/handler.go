package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/prajnapras19/attacher/attachment"
	"github.com/prajnapras19/attacher/config"
	"github.com/prajnapras19/attacher/constants"
	"github.com/prajnapras19/attacher/lib"
	"github.com/prajnapras19/attacher/user"
)

type Handler interface {
	HealthCheck(*gin.Context)
	GetLoginPage(*gin.Context)
	DoLogin(*gin.Context)
	ListActiveFiles(*gin.Context)
	DownloadAttachment(*gin.Context)
}

type handler struct {
	cfg               *config.Config
	userService       user.Service
	attachmentService attachment.Service
}

func (h *handler) HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func NewHandler(
	cfg *config.Config,
	userService user.Service,
	attachmentService attachment.Service,
) Handler {
	return &handler{
		cfg:               cfg,
		userService:       userService,
		attachmentService: attachmentService,
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

func (h *handler) ListActiveFiles(c *gin.Context) {
	jwtClaims, err := lib.GetJWTClaimsFromContext(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, constants.ListFilePage, gin.H{
			constants.Error: err.Error(),
		})
		return
	}

	res, err := h.attachmentService.GetAllActiveAttachmentsByUserID(jwtClaims.ID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, constants.LoginPage, gin.H{
			constants.Error: err.Error(),
		})
		return
	}

	var data []attachment.AttachmentResponse
	for _, d := range res {
		data = append(data, d.ToAttachmentResponse())
	}

	c.HTML(http.StatusOK, constants.ListFilePage, gin.H{
		constants.Attachments: data,
	})
}

func (h *handler) DownloadAttachment(c *gin.Context) {
	jwtClaims, err := lib.GetJWTClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			constants.Error: err.Error(),
		})
		return
	}

	serial := c.Param(constants.Serial)
	res, err := h.attachmentService.GetActiveAttachmentByUserIDAndSerial(jwtClaims.ID, serial)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			constants.Error: err.Error(),
		})
		return
	}

	content, err := os.ReadFile(res.Path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			constants.Error: err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(res.Path))
	c.Data(http.StatusOK, "application/octet-stream", content)
}
