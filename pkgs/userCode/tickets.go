package userCode

import (
	"context"
	"fmt"
	"time"
)

type Ticket struct {
	ID string `json:"id"`
	// type of ticket
	TicketName string `json:"ticketName"`
	// user id of the ticket's owner
	Holder string `json:"holder"`
	BoughtAt time.Time `json:"boughtAt"`
	// checkin time
	// nil when ticket is not checked
	CheckedAt time.Time `json:"checkedAt"`
	// qrcode filename.
	qrcode string
}

func (uc *PgxTicketStorage) InsertTicket(ctx context.Context, uId, tname, qrcode string) (id string, err error) {
	s := `INSERT INTO tickets(ticket_name, qrcode, holder) VALUES($1, $2, $3) RETURNING id`
	err = uc.db.QueryRow(ctx, s, tname, qrcode, uId).Scan(&id)
	if err != nil {
		err = fmt.Errorf("InsertTicket: %w\n", err)
	}
	return
}

func (uc *PgxTicketStorage) TicketCheck(ctx context.Context, tId string) error {
	s := `UPDATE tickets SET checked_at = NOW() WHERE id = $1`
	result, err := uc.db.Exec(ctx, s, tId)
	if err != nil || result.RowsAffected() < 1 {
		return fmt.Errorf("TicketCheck: %w", err)
	}
	return nil
}
