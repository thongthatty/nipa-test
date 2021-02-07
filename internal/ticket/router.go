package ticket

import "github.com/labstack/echo/v4"

// Register create route path for CRUD
func (t *Ticket) Register(v1 *echo.Group) {
	ticketGroup := v1.Group("/ticket")
	ticketGroup.GET("", t.Get)
	ticketGroup.POST("", t.Create)
	ticketGroup.PUT("/:id", t.Update)
}
