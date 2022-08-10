// business logic
package userCode

type CodeAdaptor interface {
	RegisterUser(fName, lName, email, contact string) (string, error)
	BuyTicket(uId, tname string) (string, error)
	GetUsersTicket(email string) (*User, []*Ticket, error)
}

type CodeChecker interface {
	GetUserTicketInfo(tId string) (*User, *Ticket, error)
	CheckInTicket(tId string) error
}
