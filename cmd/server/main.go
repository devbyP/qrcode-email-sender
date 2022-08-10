package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/devbyP/qrcode-email-sender/pkgs/db"
	"github.com/devbyP/qrcode-email-sender/pkgs/httpServer"
	"github.com/devbyP/qrcode-email-sender/pkgs/userCode"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables from .env file, if global env variables not exists.
	godotenv.Load()

	var dbCon string
	var port int
	var err error
	var pdb *pgx.Conn

	// set context timeout for seting up server.
	setupCtx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	// connect to 
	dbCon = os.Getenv("DATABASE_URL")
	if pdb, err = db.Connect(setupCtx, dbCon); err != nil {
		log.Fatal(err)
	}

	// database use uuid as unique id
	err = db.GetUUIDLib(setupCtx, pdb)
	if err != nil {
		log.Fatal(err)
	}

	// migrate table to database.
	err = userCode.CreateUserDBPgx(setupCtx, pdb, &userCode.User{}, &userCode.Ticket{})
	if err != nil {
		log.Fatal(err)
	}

	// create store db objects and assign to handlers.
	// handler will have access to these db globally inside their package.
	httpServer.UserDB = userCode.NewPgxUserStore(pdb)
	httpServer.TicketDB = userCode.NewPgxTicketStore(pdb)
	httpServer.TxDB = userCode.NewPgxTxStore(pdb)

	// get port and set run server
	if port, err = strconv.Atoi(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
	httpServer.RunServer(port)
}
