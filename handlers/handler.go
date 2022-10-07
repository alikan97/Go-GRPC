package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alikan97/Go-GRPC/handlers/middleware"
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
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	BaseURL         string
	TimeoutDuration time.Duration
}

// NewHandler intializes the handler required inkected services along with the http routes\
// Does not return as it deals directly with a reference to the gin engine
func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group(c.BaseURL)

	if gin.Mode() != gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, model.NewServiceUnavailable()))
		fmt.Println("Timeout middleware active...")
		g.GET("/account", middleware.AuthUser(h.TokenService), h.account)
		g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
		g.PUT("/updatedetails", middleware.AuthUser(h.TokenService), h.UpdateDetails)
	} else {
		g.PUT("/updatedetails", h.UpdateDetails)
		g.GET("/account", h.account)
		g.POST("/tokens", h.Tokens)
	}
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.GET("/allassets", h.GetAllAssets)
}

func (h *Handler) GetAllAssets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hlllo": "All assets",
	})
}
