package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context, str string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, str)
}

// check if timezone is valid and get string value of timezone location.
// see more information at https://pkg.go.dev/time#LoadLocation.
func timezoneValidator(zone string) (string, error) {
	loc, err := time.LoadLocation(zone)
	return loc.String(), err
}

// set database timezone.
// postgresql timezone.
func SetTimezone(ctx context.Context, db *pgx.Conn, timezone string) error {
	var err error
	fail := func(er error) error {
		if err == nil {
			return nil
		}
		return fmt.Errorf("SetTimezone: %w", er)
	}
	if timezone, err = timezoneValidator(timezone); err != nil {
		return fail(err)
	}
	_, err = db.Exec(ctx, "SET TIMEZONE TO $1", timezone)
	return fail(err)
}

// run this acter connect to database to get uuid extension on database.
func GetUUIDLib(ctx context.Context, db *pgx.Conn) error {
	var s, result string
	// fail, create an wraped error.
	fail := func(err error, reson string) error {
		return fmt.Errorf("GetUUIDLib, %s:\n%w", reson, err)
	}
	s = "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\""
	tx, err := db.Begin(ctx)
	if err != nil {
		return fail(err, "cannot begin transaction")
	}
	// defer call rollback if commit has been called, rollback no-op
	defer tx.Rollback(ctx)
	if _, err = tx.Exec(ctx, s); err != nil {
		return fail(err, "cannot Create extension \"uuid-ossp\"")
	}
	// check if uuid created.
	s = "SELECT uuid_generate_v4()"
	if err = tx.QueryRow(ctx, s).Scan(&result); err != nil {
		return fail(err, "fail to scan select result")
	}
	if err = tx.Commit(ctx); err != nil {
		return fail(err, "cannot commit transaction")
	}
	return nil
}