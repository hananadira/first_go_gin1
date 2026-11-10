package repository

import (
	"context"
	"errors"

	"github.com/hananadira/ecommarce-api/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrEmailTaken = errors.New("email sudah terdaftar")
	ErrUserNotFound = errors.New("user tidak ditemukan")
) 

type UserRepository struct {
	db *pgxpool.Pool
}


// NewUserRepository adalah construktor. go tidak punya constructor bawaan jadi polanya fungsi berawalan New
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	query := `INSERT INTO users (name, email, password)
			VALUES ($1, $2, $3)
			RETURNING id, created_at`
	
	err := r.db.QueryRow(ctx, query, u.Name, u.Email, u.PasswordHash).
	//Scan(&u.ID, ...) memasukan hasil nya ke stuct, fungsi & memberi alamat variabel agar bisa diisi 
		Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrEmailTaken
		}
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, name, email, password, created_at
			FROM users
			WHERE email = $1`

	var u model.User
	err := r.db.QueryRow(ctx, query, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}