package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/alikan97/Go-GRPC/handlers"
	"github.com/alikan97/Go-GRPC/repository"
	"github.com/alikan97/Go-GRPC/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func inject(d *repository.DataSources) (*gin.Engine, error) {
	envFile, _ := godotenv.Read(".env")
	userRepo := repository.NewUserRepository(d.DB)
	tokenRepo := repository.NewTokenRepository(d.RedisClient)

	userService := services.NewUserService(&services.UConfig{
		UseRepository: userRepo,
	})

	//load rsa key file
	privKeyFile := envFile["PRIV_KEY_FILE"]
	priv, err := ioutil.ReadFile(privKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not ready private key pem file %w", err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key, %w", err)
	}

	pubKeyFile := envFile["PUB_KEY_FILE"]
	pub, err := ioutil.ReadFile(pubKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not ready public key pem file %w", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key, %w", err)
	}
	refreshSecret := envFile["REFRESH_SECRET"]

	idTokenExp := envFile["ID_TOKEN_EXP"]
	refreshTokenExp := envFile["REFRESH_TOKEN_EXP"]

	tokenExp, err := strconv.ParseInt(idTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("Could not parse ID Token EXP as integer, %w", err)
	}

	refreshExp, err := strconv.ParseInt(refreshTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("Could not parse Refresh Token EXP as integer, %w", err)
	}

	tokenService := services.NewTokenService(&services.TSConfig{
		TokenRepository:      tokenRepo,
		PrivKey:              privKey,
		Pubkey:               pubKey,
		RefreshSecret:        refreshSecret,
		IDExpirationSec:      tokenExp,
		RefreshExpirationSec: refreshExp,
	})

	router := gin.Default()

	baseURL := envFile["BASE_URL"]
	ht := envFile["HANDLER_TIMEOUT"]
	handlerTimeout, err := strconv.ParseInt(ht, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("Could not parse handler timeout as integer, %w", err)
	}

	handlers.NewHandler(&handlers.Config{
		R:               router,
		UserService:     userService,
		TokenService:    tokenService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})

	return router, nil
}
