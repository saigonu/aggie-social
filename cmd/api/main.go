package main

import (
	"log"

	"github.com/saigonu/aggie-social/internal/env"
	"github.com/saigonu/aggie-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
