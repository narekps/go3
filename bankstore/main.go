package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/narekps/go3/bankstore/api"
	db "github.com/narekps/go3/bankstore/db/sqlc"
	"github.com/narekps/go3/bankstore/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not read config file", err)
	}

	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}

	defer pool.Close()

	store := db.NewStore(pool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can not create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}
}
