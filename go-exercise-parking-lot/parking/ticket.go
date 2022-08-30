package parking

type Ticket struct {
	Id string
}

func NewTicket(id string) *Ticket{
	return &Ticket{Id: id}
}

func GetTicket(ticket *Ticket) string {
	return ticket.Id
}