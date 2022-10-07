package services

import (
	"context"
	"crypto/rsa"
	"log"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/google/uuid"
)

type TokenService struct {
	TokenRepository      model.TokenRepository
	PrivKey              *rsa.PrivateKey
	Pubkey               *rsa.PublicKey
	RefreshSecret        string
	IDExpirationSec      int64
	RefreshExpirationSec int64
}

type TSConfig struct {
	TokenRepository      model.TokenRepository
	PrivKey              *rsa.PrivateKey
	Pubkey               *rsa.PublicKey
	RefreshSecret        string
	IDExpirationSec      int64
	RefreshExpirationSec int64
}

func NewTokenService(c *TSConfig) model.TokenService {
	return &TokenService{
		TokenRepository:      c.TokenRepository,
		PrivKey:              c.PrivKey,
		Pubkey:               c.Pubkey,
		RefreshSecret:        c.RefreshSecret,
		IDExpirationSec:      c.IDExpirationSec,
		RefreshExpirationSec: c.RefreshExpirationSec,
	}
}

func (s *TokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	if prevTokenID != "" {
		if err := s.TokenRepository.DeleteRefreshToken(ctx, u.UID.String(), prevTokenID); err != nil {
			log.Printf("Could not delete previous refreshToken for uid: %v, tokenID: %v\n", u.UID.String(), prevTokenID)
			return nil, err
		}
	}

	idToken, err := generateIDToken(u, s.PrivKey, s.IDExpirationSec)

	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v", u.UID, err.Error())
		return nil, model.NewInternal()
	}

	RefreshToken, err := generateRefreshToken(u.UID, s.RefreshSecret, s.RefreshExpirationSec)

	if err != nil {
		log.Printf("Error generating refreshToken for uid: %v, Error: %v\n", u.UID, err.Error())
		return nil, model.NewInternal()
	}

	// Store the token into redis
	if err := s.TokenRepository.SetRefreshToken(ctx, u.UID.String(), RefreshToken.ID.String(), RefreshToken.ExpiresIn); err != nil {
		log.Printf("Error storing tokenID for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, model.NewInternal()
	}

	return &model.TokenPair{
		IDToken:      model.IDToken{SS: idToken},
		RefreshToken: model.RefreshToken{SS: RefreshToken.SS, ID: RefreshToken.ID, UID: u.UID},
	}, nil
}

func (s *TokenService) ValidateIDToken(tokenString string) (*model.User, error) {
	claims, err := validateIDToken(tokenString, s.Pubkey)

	if err != nil {
		log.Printf("Unable to validate or parse jwt token - Error: %v\n", err)
		return nil, model.NewAuthorization("Unable to verify user from id token")
	}

	return claims.User, nil
}

func (s *TokenService) ValidateRefreshToken(refreshTokenString string) (*model.RefreshToken, error) {
	claims, err := validateRefreshToken(refreshTokenString, s.RefreshSecret)

	// For now we'll just return the error and handle logging in service level
	if err != nil {
		log.Printf("Unable to validate or parse refreshToken for string: %s\n%v\n", refreshTokenString, err)
		return nil, model.NewAuthorization("Unable to verify user from refresh token")
	}

	tokenUUID, err := uuid.Parse(claims.Id)

	if err != nil {
		log.Printf("Claims ID could not be parsed as UUID: %s\n%v\n", claims.Id, err)
		return nil, model.NewAuthorization("Unable to verify user from refresh token")
	}

	return &model.RefreshToken{
		SS:  refreshTokenString,
		ID:  tokenUUID,
		UID: claims.UID,
	}, nil
}

func (s *TokenService) Signout(ctx context.Context, uid uuid.UUID) error {
	return s.TokenRepository.DeleteUserRefreshToken(ctx, uid.String())
}
