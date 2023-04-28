package main

import (
	"database/sql"
	"log"

	"github.com/bulat-is-online/simplebank/api"
	db "github.com/bulat-is-online/simplebank/db/sqlc"
	"github.com/bulat-is-online/simplebank/db/util"
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
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can't start server:", err)
	}
}
