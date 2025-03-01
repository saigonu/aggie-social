package main

import (
	"log"

	"github.com/saigonu/aggie-social/internal/db"
	"github.com/saigonu/aggie-social/internal/env"
	"github.com/saigonu/aggie-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
