package handlers

import (
	"fmt"
	"time"

	"github.com/alikan97/Go-GRPC/handlers/middleware"
	"github.com/alikan97/Go-GRPC/model"
	pb "github.com/alikan97/Go-GRPC/proto"
	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
	Client       pb.CryptoClient
}

// Config will hold sercvices that will eventually be injkected into this handler layer on handler initialization
type Config struct {
	Client          pb.CryptoClient
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
		Client:       c.Client,
	}

	g := c.R.Group(c.BaseURL)

	if gin.Mode() != gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, model.NewServiceUnavailable()))
		fmt.Println("Timeout middleware active...")
		g.GET("/account", middleware.AuthUser(h.TokenService), h.account)
		g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
		g.PUT("/updatedetails", middleware.AuthUser(h.TokenService), h.UpdateDetails)
		g.GET("/allassets", middleware.AuthUser(h.TokenService), h.GetAllAssets)
		g.GET("/asset", middleware.AuthUser(h.TokenService), h.getAssetBySymbol)
		g.GET("/recenttrades", middleware.AuthUser(h.TokenService), h.getRecentTrades)
		g.GET("/quote", middleware.AuthUser(h.TokenService), h.getCurrentQuote)
		g.POST("/login", h.Signin)
		g.POST("/register", h.Signup)
		g.POST("/tokens", h.Tokens)
	}
}
