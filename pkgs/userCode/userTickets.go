package userCode

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (ut *PgxUserTicketTransaction) SelectUserTicket(ctx context.Context, uEmail string) (user *User, tickets []*Ticket, err error) {
	// user sequence select method to get a user, and tickets which that user own.
	// use join in this situaltion (* to many) would be little slower because of duplicate main table data.
	// Likewise, if user hold more than one ticket,
	// when use join there will select multiple user data along with ticket data row.
	/* join statement example.
		SELECT u.id, u.first_name, u.last_name, u.email, u.contact, 
		t.id, t.ticket_name, t.qrcode, t.bought_at, t.used_at, 
		FROM users as u 
		JOIN tickets as t 
		ON u.id = t.holder 
		WHERE u.email = $1
	*/
	fail := func(e error) (*User, []*Ticket, error) {
		return nil, nil, fmt.Errorf("SelectUserTicket, %w\n", e)
	}
	tx, err := ut.db.Begin(ctx)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback(ctx)
	if user, err = ut.selectUserByEmail(ctx, uEmail, tx); err != nil {
		return fail(err)
	}

	tickets, err = ut.selectTicketByUserID(ctx, user.ID, tx)
	if err != nil {
		return fail(err)
	}
	if err = tx.Commit(ctx); err != nil {
		return fail(err)
	}
	return
}

// this function is called by SelectUserTicket function.
// use transaction for sequence select.
func (uc *PgxUserTicketTransaction) selectTicketByUserID(ctx context.Context, uId string, tx pgx.Tx) ([]*Ticket, error) {
	s := `SELECT id, ticket_name, holder, bought_at, used_at, qrcode FROM tickets WHERE holder = $1`
	fail := func(e error) ([]*Ticket, error) {
		return nil, fmt.Errorf("SelectTicketByUserID: %w\n", e)
	}
	rows, err := tx.Query(ctx, s, uId)
	if err != nil {
		return fail(err)
	}
	var tickets []*Ticket
	for rows.Next() {
		var t *Ticket
		if err = rows.Scan(&t.ID, &t.TicketName, &t.Holder, &t.BoughtAt, &t.CheckedAt, &t.qrcode); err != nil {
			return fail(err)
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

// this function is called by SelectUserTicket function.
// use transaction for sequence select.
func (uc *PgxUserTicketTransaction) selectUserByEmail(ctx context.Context, uEmail string, tx pgx.Tx) (*User, error) {
	s := `SELECT id, first_name, last_name, email, contact, created_at FROM users WHERE email = $1`
	// use QueryRow because email are unique, so one row will return when search row with email.
	var u *User
	err := tx.QueryRow(ctx, s, uEmail).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Contact, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("SelectuserByEmail: %w\n", err)
	}
	return u, nil
}