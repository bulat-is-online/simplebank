package main

import (
	"database/sql"
	"log"

	db "github.com/bulat-is-online/simplebank/db/sqlc"
	"github.com/bulat-is-online/simplebank/db/util"
	"github.com/bulat-is-online/simplebank/ft/docker/api"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannott start server:", err)
	}
}
