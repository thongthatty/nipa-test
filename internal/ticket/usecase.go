package ticket

import (
	"net/http"
	"nipatest/main/internal/model"
	"nipatest/main/pkg/helper"

	"github.com/labstack/echo/v4"
)

// Get godoc
// @Summary GET TICKETS
// @Description GET TICKETS
// @Tags GET TICKETS
// @Param from query int false "Start date (Unix time)"
// @Param to query int false "End date (Unix time)"
// @Param status query string false "Filter by ticket status" Enums(PENDING, ACCEPTED, RESOLVED, REJECTED)
// @Param page query int false "page of pagination" minimum(1)
// @Param page_size query int false "total record to show" minimum(100)
// @Success 200 {array} model.Ticket
// @Failure 400 {object} helper.Error
// @Failure 422 {object} helper.Error
// @Router /ticket [get]
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

// Create godoc
// @Summary CREATE TICKET
// @Description CREATE TICKET
// @Tags CREATE TICKETS
// @Accept  json
// @Produce  json
// @Param ticketInfo body model.TicketCreateRequest true "Body of ticket"
// @Success 200 {object} model.Ticket
// @Failure 400 {object} helper.Error
// @Failure 422 {object} helper.Error
// @Router /ticket [post]
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

// Update godoc
// @Summary UPDATE TICKET
// @Description UPDATE TICKET
// @Tags UPDATE TICKETS
// @Accept  json
// @Produce  json
// @Param ticketInfo body model.TicketStatusUpdateRequest true "Body of ticket"
// @Success 200 {string} string "Update ticket successfully"
// @Failure 400 {object} helper.Error
// @Failure 422 {object} helper.Error
// @Router /ticket [put]
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
