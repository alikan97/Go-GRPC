package services

import (
	"context"
	"crypto/rsa"
	"log"

	"github.com/alikan97/Go-GRPC/model"
)

type TokenService struct {
	PrivKey       *rsa.PrivateKey
	Pubkey        *rsa.PublicKey
	RefreshSecret string
}

type TSConfig struct {
	PrivKey       *rsa.PrivateKey
	Pubkey        *rsa.PublicKey
	RefreshSecret string
}

func NewTokenService(c *TSConfig) model.TokenService {
	return &TokenService{
		PrivKey:       c.PrivKey,
		Pubkey:        c.Pubkey,
		RefreshSecret: c.RefreshSecret,
	}
}

func (s *TokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	idToken, err := generateIDToken(u, s.PrivKey)

	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v", u.UID, err.Error())
		return nil, model.NewInternal()
	}

	RefreshToken, err := generateRefreshToken(u.UID, s.RefreshSecret)

	if err != nil {
		log.Printf("Error generating refreshToken for uid: %v, Error: %v\n", u.UID, err.Error())
		return nil, model.NewInternal()
	}

	return &model.TokenPair{
		IDToken:      idToken,
		RefreshToken: RefreshToken.SS,
	}, nil
}
