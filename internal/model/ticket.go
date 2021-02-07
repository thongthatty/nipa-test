package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type TicketStatus string

var (
	TicketStatusPending  TicketStatus = "PENDING"
	TicketStatusAccepted TicketStatus = "ACCEPTED"
	TicketStatusResolved TicketStatus = "RESOLVED"
	TicketStatusRejected TicketStatus = "REJECTED"
)

var stringToTicketStatusMap = map[string]TicketStatus{
	"PENDING":  TicketStatusPending,
	"ACCEPTED": TicketStatusAccepted,
	"RESOLVED": TicketStatusResolved,
	"REJECTED": TicketStatusRejected,
}

func (t TicketStatus) String() string {
	return string(t)
}

// MapString -
func MapTicketStatusString(ticketStatusStr string) TicketStatus {
	return stringToTicketStatusMap[ticketStatusStr]
}

func (t TicketStatus) Valid() error {
	if stringToTicketStatusMap[t.String()] == "" {
		return fmt.Errorf("%s", "Ticket status does not matched")
	}
	return nil
}

// Ticket infomation of ticket table in db (MODEL)
type Ticket struct {
	ID          uint         `gorm:"primary_key" json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"desc"`
	ContactInfo string       `json:"contactInfo"`
	Status      TicketStatus `json:"status"`
	CreatedAt   time.Time    `json:"createAt"`
	UpdatedAt   time.Time    `json:"updateAt"`
}

type RangeTimeFilter struct {
	From *time.Time
	To   *time.Time
}

// TicketGetRequest request of get ticket
type TicketGetRequest struct{}

func (tcr *TicketGetRequest) Bind(c echo.Context) (int, int, *TicketStatus, RangeTimeFilter, error) {
	var (
		err             error
		pageNum         int
		pageSizeNum     int
		ticketStatus    *TicketStatus
		rangeTimeFilter RangeTimeFilter
	)

	page := c.QueryParam("page")
	if page != "" {
		pageNum, err = strconv.Atoi(page)
		if err != nil {
			return 0, 0, nil, rangeTimeFilter, fmt.Errorf("Param [%s] type should be number", "page")
		}
	}

	pageSize := c.QueryParam("page_size")
	if pageSize != "" {
		pageSizeNum, err = strconv.Atoi(pageSize)
		if err != nil {
			return 0, 0, nil, rangeTimeFilter, fmt.Errorf("Param [%s] type should be number", "page_size")
		}
	}

	status := c.QueryParam("status")
	if status != "" {
		tks := MapTicketStatusString(status)
		ticketStatus = &tks
		err = ticketStatus.Valid()
		if ticketStatus.Valid() != nil {
			return 0, 0, nil, rangeTimeFilter, err
		}
	}

	from := c.QueryParam("from")
	to := c.QueryParam("to")
	if from != "" && to != "" {
		fromUnix, err := strconv.ParseInt(from, 10, 64)
		if err != nil {
			return 0, 0, nil, rangeTimeFilter, fmt.Errorf("Param [%s] type should be unix time format", "from")
		}
		fromT := time.Unix(fromUnix, 0)

		rangeTimeFilter.From = &fromT
	}
	if to != "" {
		toUnix, err := strconv.ParseInt(to, 10, 64)
		if err != nil {
			return 0, 0, nil, rangeTimeFilter, fmt.Errorf("Param [%s] type should be unix time format", "to")
		}
		toT := time.Unix(toUnix, 0)
		rangeTimeFilter.To = &toT
	}

	return pageNum, pageSizeNum, ticketStatus, rangeTimeFilter, nil
}

// TicketCreateRequest request of create ticket
type TicketCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"desc"`
	ContactInfo string `json:"contactInfo"`
}

// Bind -
func (tcr *TicketCreateRequest) Bind(c echo.Context, tk *Ticket) error {
	if err := c.Bind(tcr); err != nil {
		return err
	}
	if err := c.Validate(tcr); err != nil {
		return err
	}
	tk.Name = tcr.Name
	tk.Status = TicketStatusPending
	tk.CreatedAt = time.Now()
	tk.Description = tcr.Description
	tk.ContactInfo = tcr.ContactInfo
	return nil
}

// TicketStatusUpdateRequest request of update status of ticket by id
type TicketStatusUpdateRequest struct {
	Status TicketStatus `json:"status" validate:"required"`
}

// Bind -
func (tcr *TicketStatusUpdateRequest) Bind(c echo.Context, tk *Ticket) error {
	id := c.Param("id")
	if id == "" {
		return fmt.Errorf("%s", "Param [id] is required")
	}
	idi, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("%s", "Param [id] type should be number")
	}

	if err := c.Bind(tcr); err != nil {
		return err
	}
	if err := c.Validate(tcr); err != nil {
		return err
	}

	if err := tcr.Status.Valid(); err != nil {
		return err
	}

	tk.ID = uint(idi)
	tk.Status = tcr.Status
	tk.UpdatedAt = time.Now()
	return nil
}
