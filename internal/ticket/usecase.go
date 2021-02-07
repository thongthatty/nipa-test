package ticket

import (
	"net/http"
	"nipatest/main/internal/model"
	"nipatest/main/pkg/helper"

	"github.com/labstack/echo/v4"
)

// Get usecase for get infomation of ticket
// @Router / [get]
// @Query ?page ?pageSize ?status
// e.g. http://localhost:1323/api/ticket?page=1&pageSize=10&status=REJECTED
func (tr *Ticket) Get(c echo.Context) error {
	var (
		err error
	)

	tcr := &model.TicketGetRequest{}
	page, pageSize, statusFilter, rangeTimeFilter, err := tcr.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.NewError(err))
	}

	results, err := tr.ticketRepo.Get(page, pageSize, statusFilter, rangeTimeFilter)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.NewError(err))
	}

	return c.JSON(http.StatusOK, results)
}

// Create usecase for create infomation of ticket
// @Router / [post]
func (tr *Ticket) Create(c echo.Context) error {
	var (
		tk   model.Ticket
		resp *model.Ticket
		err  error
	)
	tcr := &model.TicketCreateRequest{}
	if err = tcr.Bind(c, &tk); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.NewError(err))
	}

	if resp, err = tr.ticketRepo.Create(&tk); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.NewError(err))
	}
	return c.JSON(http.StatusCreated, resp)
}

// Update usecase for update infomation of ticket
// @Router /:id [put]
func (tr *Ticket) Update(c echo.Context) error {
	var (
		tk  model.Ticket
		err error
	)
	tcr := &model.TicketStatusUpdateRequest{}
	if err = tcr.Bind(c, &tk); err != nil {
		// fmt.Sprintf("%s", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, helper.NewError(err))
	}

	if err = tr.ticketRepo.Update(&tk); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.NewError(err))
	}
	return c.JSON(http.StatusOK, "Update ticket successfully")
}
