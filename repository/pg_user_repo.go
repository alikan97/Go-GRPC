package repository

import (
	"context"
	"log"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PGUserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) model.UserRepository {
	return &PGUserRepository{
		DB: db,
	}
}

func (r *PGUserRepository) Create(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *"

	if err := r.DB.Get(u, query, u.Email, u.Password); err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create user with email: %v Reason: %v\n", u.Email, err.Code.Name())
			return model.NewConflict("email", u.Email)
		}

		log.Printf("Could not create user with email: %v Reason: %v\n", u.Email, err)
		return model.NewInternal()
	}
	return nil
}

func (r *PGUserRepository) FindById(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}
	query := "SELECT * FROM users WHERE uid=$1"

	if err := r.DB.Get(user, query, uid); err != nil {
		return user, model.NewNotFound("uid", uid.String())
	}

	return user, nil
}
