package main

import (
	"nipatest/main/db"
	"nipatest/main/internal/repository"
	"nipatest/main/internal/ticket"
	"nipatest/main/router"
)

func main() {
	r := router.New()

	v1 := r.Group("/api")

	d := db.Initialize()
	db.AutoMigrate(d)

	tr := repository.NewTicketRepo(d)
	tk := ticket.NewTicket(tr)

	tk.Register(v1)
	r.Logger.Fatal(r.Start(":1323"))
}
