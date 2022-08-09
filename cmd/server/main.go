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
)

func main() {
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
		log.Println(err)
	}
	// migrate table to database.
	userCode.CreateUserDBPgx(setupCtx, pdb, &userCode.User{}, &userCode.Ticket{})

	// create store db objects and assign to handlers.
	// handler will have access to these db globally inside their package.
	httpServer.UserDB = userCode.NewPgxUserStore(pdb)
	httpServer.TicketDB = userCode.NewPgxTicketStore(pdb)
	httpServer.TxDB = userCode.NewPgxTxStore(pdb)

	// get port and set run server
	if port, err = strconv.Atoi(os.Getenv("PORT")); err != nil {
		log.Println(err)
	}
	httpServer.RunServer(port)
}
