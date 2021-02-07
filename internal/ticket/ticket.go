package ticket

import (
	"nipatest/main/internal/model"
)

// TicketRepo -
type TicketRepo interface {
	Get(page, pageSize int, statusFilter *model.TicketStatus, rangeTimeFilter model.RangeTimeFilter) ([]model.Ticket, error)
	Create(tk *model.Ticket) (*model.Ticket, error)
	Update(*model.Ticket) error
}

// Ticket -
type Ticket struct {
	ticketRepo TicketRepo
}

// NewTicket constuctor for Ticket
func NewTicket(tr TicketRepo) *Ticket {
	return &Ticket{
		ticketRepo: tr,
	}
}
