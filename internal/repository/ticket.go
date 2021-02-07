package repository

import (
	"nipatest/main/db"
	"nipatest/main/internal/model"

	"github.com/jinzhu/gorm"
)

// TicketRepo query db for ticket package
type TicketRepo struct {
	db *gorm.DB
}

// NewTicketRepo constuctor for TicketRepo
func NewTicketRepo(db *gorm.DB) *TicketRepo {
	return &TicketRepo{
		db: db,
	}
}

// Get -
func (tr *TicketRepo) Get(page, pageSize int, statusFilter *model.TicketStatus, rangeTimeFilter model.RangeTimeFilter) ([]model.Ticket, error) {
	var tks []model.Ticket
	query := tr.db.Scopes(db.Paginate(page, pageSize))
	if statusFilter != nil {
		query = query.Where("status = ?", statusFilter.String())
	}
	if rangeTimeFilter.From != nil && rangeTimeFilter.To != nil {
		query = query.Where("created_at BETWEEN ? AND ?", rangeTimeFilter.From, rangeTimeFilter.To)
	} else if rangeTimeFilter.From != nil && rangeTimeFilter.To == nil {
		query = query.Where("created_at > ?", rangeTimeFilter.From)
	} else if rangeTimeFilter.To != nil && rangeTimeFilter.From == nil {
		query = query.Where("created_at < ?", rangeTimeFilter.To)
	}
	if err := query.Find(&tks).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return tks, nil
}

// Create repo for create ticket
func (tr *TicketRepo) Create(tk *model.Ticket) (*model.Ticket, error) {
	if err := tr.db.Create(&tk).Error; err != nil {
		return nil, err
	}
	return tk, nil
}

// Update repo for update ticket
func (tr *TicketRepo) Update(tk *model.Ticket) error {
	return tr.db.Model(tk).Update(tk).Error
}
