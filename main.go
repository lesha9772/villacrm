package main

import (
	"github.com/lesha9772/villacrm/bootstrap"
	"github.com/lesha9772/villacrm/middleware/identity"
	"github.com/lesha9772/villacrm/routes"
)

var app = bootstrap.New("VillaCRM", "edi.ultras@gmai.com",
	identity.Configure,
	routes.Configure,
)

func init() {
	app.Bootstrap()
}

func main() {
	app.Listen(":8080")
}
