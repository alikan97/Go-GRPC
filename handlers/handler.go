package handlers

import (
	"net/http"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

// Config will hold sercvices that will eventually be injkected into this handler layer on handler initialization
type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

// NewHandler intializes the handler required inkected services along with the http routes\
// Does not return as it deals directly with a reference to the gin engine
func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group("api/")

	g.GET("/allassets", h.GetAllAssets)
	g.GET("/account", h.account)
	g.POST("/signup", h.Signup)
}

func (h *Handler) GetAllAssets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hlllo": "All assets",
	})
}
