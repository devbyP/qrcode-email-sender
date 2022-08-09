package userCode

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserStorage interface {
	SelectUser(ctx context.Context, id string) (*User, error)
	InsertUser(ctx context.Context, fn, ln, e, c string) (string, error)
}

type TicketStorage interface {
	InsertTicket(ctx context.Context, uId, tname, qrcode string) (id string, err error)
	TicketCheck(ctx context.Context, tId string) error
}

type UserTicketTransactionStorage interface {
	SelectUserTicket(ctx context.Context, email string) (*User, []*Ticket, error)
}

type PgxUserStorage struct {
	db *pgx.Conn
}
func NewPgxUserStore(d *pgx.Conn) *PgxUserStorage {
	return &PgxUserStorage{db: d}
}

type PgxTicketStorage struct {
	db *pgx.Conn
}
func NewPgxTicketStore(d *pgx.Conn) *PgxTicketStorage {
	return &PgxTicketStorage{db: d}
}

type PgxUserTicketTransaction struct {
	db *pgx.Conn
}

func NewPgxTxStore(d *pgx.Conn) *PgxUserTicketTransaction {
	return &PgxUserTicketTransaction{db: d}
}
