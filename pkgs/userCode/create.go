package userCode

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TableCreator interface {
	createTable(ctx context.Context, db *pgx.Conn) error
}

func (u User) createTable(ctx context.Context, db *pgx.Conn) error {
	s := `CREATE TABLE IF NOT EXISTS users(
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			email VARCHAR(120) NOT NULL UNIQUE,
			contact VARCHAR(10),
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`
	if _, err := db.Exec(ctx, s); err != nil {
		return fmt.Errorf("Create User Table: %v\n", err)
	}
	return nil
}

func (t Ticket) createTable(ctx context.Context, db *pgx.Conn) error {
	s := `CREATE TABLE IF NOT EXISTS tickets(
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		ticket_name VARCHAR(50) NOT NULL,
		qrcode VARCHAR(50) NOT NULL,
		holder uuid REFERENCES users(id) NOT NULL,
		bought_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		checked_at TIMESTAMPTZ
	)`
	if _, err := db.Exec(ctx, s); err != nil {
		return fmt.Errorf("Create Ticket Table: %v\n", err)
	}
	return nil
}

func CreateUserDBPgx(ctx context.Context, db *pgx.Conn,  ct ...TableCreator) error {
	for _, c := range ct {
		if err := c.createTable(ctx, db); err != nil {
			return err
		}
	}
	return nil
}
