package userCode

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Contact string `json:"contact"`
	CreatedAt time.Time `json:"createdOn"`
}

func (uc *PgxUserStorage) SelectUser(ctx context.Context, id string) (*User, error) {
	var u *User
	s := `SELECT id, first_name, last_name, email, contact FROM users WHERE id = $1`
	if err := uc.db.QueryRow(ctx, s, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Contact); err != nil {
		return nil, fmt.Errorf("SelectUser: %w\n", err)
	}
	return u, nil
}

func (uc *PgxUserStorage) InsertUser(ctx context.Context, fn, ln, e, c string) (string, error) {
	var id, s string
	s = `INSERT INTO users(first_name, last_name, email, contact) VALUES($1, $2, $3, $4) RETURNING id`
	if err := uc.db.QueryRow(ctx, s, fn, ln, e, c).Scan(&id); err != nil {
		return "", fmt.Errorf("InsertUser: %w\n", err)
	}
	return id, nil
}
