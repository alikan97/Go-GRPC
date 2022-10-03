package services

import (
	"context"
	"log"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserRepository
}

type UConfig struct {
	UseRepository model.UserRepository
}

func NewUserService(c *UConfig) model.UserService {
	return &UserService{
		UserRepository: c.UseRepository,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, error := s.UserRepository.FindById(ctx, uid)

	return u, error
}

func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)

	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", u.Email)
		return model.NewInternal()
	}

	u.Password = pw

	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}

	return nil
}
