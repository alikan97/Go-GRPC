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

	if err := r.DB.GetContext(ctx, u, query, u.Email, u.Password); err != nil {
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

	if err := r.DB.GetContext(ctx, user, query, uid); err != nil {
		return user, model.NewNotFound("uid", uid.String())
	}

	return user, nil
}

func (r *PGUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT * FROM users WHERE email=$1"

	if err := r.DB.GetContext(ctx, user, query, email); err != nil {
		log.Printf("Unable to get user with email address: %v, Error: %v\n", email, err)
		return user, model.NewNotFound("Email", email)
	}

	return user, nil
}

func (r *PGUserRepository) Update(ctx context.Context, u *model.User) error {
	query := `UPDATE users SET name=:name, email=:email, website=:website WHERE uid=:uid RETURNING *;`

	nstmt, err := r.DB.PrepareNamedContext(ctx, query)

	if err != nil {
		log.Printf("Unable to prepare user update query, %v\n", err)
		return model.NewInternal()
	}

	if err := nstmt.GetContext(ctx, u, u); err != nil {
		log.Printf("Failed to update details for user: %v\n", u)
		return model.NewInternal()
	}

	return nil
}
