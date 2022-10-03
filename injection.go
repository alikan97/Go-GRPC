package main

import (
	"fmt"
	"io/ioutil"

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
	tokenService := services.NewTokenService(&services.TSConfig{
		PrivKey:       privKey,
		Pubkey:        pubKey,
		RefreshSecret: refreshSecret,
	})

	router := gin.Default()

	handlers.NewHandler(&handlers.Config{
		R:            router,
		UserService:  userService,
		TokenService: tokenService,
	})

	return router, nil
}
