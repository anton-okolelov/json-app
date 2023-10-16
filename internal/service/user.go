package service

import (
	"context"
	"fmt"

	"github.com/anton-okolelov/json-app/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewService(db *pgxpool.Pool) Service {
	return Service{db: db}
}

type Service struct {
	db *pgxpool.Pool
}

func (s Service) AddUser(u domain.User) (int, error) {
	var id int

	err := s.db.QueryRow(context.Background(), `
			INSERT INTO users 
			(name, age)
			VALUES 
			($1, $2)
			RETURNING id
		`, u.Name, u.Age).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("can't add user: %w", err)
	}

	return id, nil
}

func (s Service) GetUser(id int) (domain.User, error) {
	var u domain.User
	err := s.db.QueryRow(
		context.Background(),
		`SELECT name, age 
			FROM users
			WHERE id = $1`,
		id,
	).Scan(&u.Name, &u.Age)

	if err != nil {
		return u, fmt.Errorf("can't add user: %w", err)
	}

	return u, nil
}

func (s Service) GetAllUsers() ([]domain.User, error) {
	rows, err := s.db.Query(context.Background(), `SELECT name, age FROM users`)
	if err != nil {
		return nil, fmt.Errorf("don't have any users")
	}
	users := make([]domain.User, 0)
	for rows.Next() {
		var u domain.User
		err = rows.Scan(&u.Name, &u.Age)
		if err != nil {
			return nil, fmt.Errorf("error on scan user: %w", err)
		}

		users = append(users, u)
	}

	return users, nil
}
