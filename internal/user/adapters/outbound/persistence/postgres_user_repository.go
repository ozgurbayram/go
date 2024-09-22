package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mygogo/internal/user/domain"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type PostgresUserRepository struct {
	db *pgx.Conn
}

func NewPostgresUserRepository(db *pgx.Conn) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	query := `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`
	
	row := r.db.QueryRow(ctx, query, id)
	user := &domain.User{}
	
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(http.StatusText(http.StatusNotFound))
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return user, nil
}

func (r *PostgresUserRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	query := `
		SELECT id,name,email FROM users
	`
	
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, errors.New(http.StatusText(http.StatusNoContent))
	}

	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User

		err = rows.Scan(&user.Id, &user.Name, &user.Email)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, name, email)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query, user.Id, user.Name, user.Email)

	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
