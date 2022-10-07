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
	pw, err := HashPassword(u.Password)

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

func (s *UserService) Signin(ctx context.Context, u *model.User) error {
	existingUser, err := s.UserRepository.FindByEmail(ctx, u.Email)

	if err != nil {
		return model.NewAuthorization("Invalid email or password combination")
	}

	match, err := ComparePasswords(existingUser.Password, u.Password)

	if err != nil {
		return model.NewInternal()
	}

	if !match {
		return model.NewAuthorization("Invalid password")
	}

	*u = *existingUser
	return nil
}

func (s *UserService) UpdateDetails(ctx context.Context, u *model.User) error {
	err := s.UserRepository.Update(ctx, u)

	if err != nil {
		return err
	}

	return nil
}
