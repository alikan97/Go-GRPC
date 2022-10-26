package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/alikan97/Go-GRPC/handlers"
	pb "github.com/alikan97/Go-GRPC/proto"
	"github.com/alikan97/Go-GRPC/repository"
	"github.com/alikan97/Go-GRPC/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func inject(d *repository.DataSources) (*gin.Engine, error) {
	// envFile, errs := godotenv.Read(".env")

	// if errs != nil {
	// 	fmt.Printf("Error reading in env file, %v", errs)
	// }

	userRepo := repository.NewUserRepository(d.DB)
	tokenRepo := repository.NewTokenRepository(d.RedisClient)

	userService := services.NewUserService(&services.UConfig{
		UseRepository: userRepo,
	})

	//load rsa key file
	privKeyFile := os.Getenv("PRIV_KEY_FILE")
	priv, err := ioutil.ReadFile(privKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not ready private key pem file %w", err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key, %w", err)
	}

	pubKeyFile := os.Getenv("PUB_KEY_FILE")
	pub, err := ioutil.ReadFile(pubKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not ready public key pem file %w", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key, %w", err)
	}
	refreshSecret := os.Getenv("REFRESH_SECRET")

	idTokenExp := os.Getenv("ID_TOKEN_EXP")
	refreshTokenExp := os.Getenv("REFRESH_TOKEN_EXP")

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

	baseURL := os.Getenv("BASE_URL")
	ht := os.Getenv("HANDLER_TIMEOUT")
	handlerTimeout, err := strconv.ParseInt(ht, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("Could not parse handler timeout as integer, %w", err)
	}

	grpcAddress := os.Getenv("GRPC_SERVER_ADDRESS")

	serverAddress := flag.String("address", grpcAddress, "the server address")

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewCryptoClient(conn)

	handlers.NewHandler(&handlers.Config{
		Client:          c,
		R:               router,
		UserService:     userService,
		TokenService:    tokenService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})

	return router, nil
}
