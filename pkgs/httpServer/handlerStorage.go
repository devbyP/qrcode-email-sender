package httpServer

import "github.com/devbyP/qrcode-email-sender/pkgs/userCode"

var UserDB userCode.UserStorage
var TicketDB userCode.TicketStorage
var TxDB userCode.UserTicketTransactionStorage
