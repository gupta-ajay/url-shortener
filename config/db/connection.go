package db

import (
	"database/sql"
	"log"
	"urlshortner/config/dotenv"
	"urlshortner/utils/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func ConnectDB() *bun.DB {
	// pgx.ParseConfig()
	PgURL := "postgres://" +
		dotenv.Global.PgUser +
		":" + dotenv.Global.PgPassword +
		"@" + dotenv.Global.PgHost +
		":" + dotenv.Global.PgPort +
		"/" + dotenv.Global.PgDB

	isVerbose := false
	if dotenv.Global.GoEnv != "production" && dotenv.Global.GoEnv != "test" {
		isVerbose = true
		PgURL += "?sslmode=disable"
	} else if dotenv.Global.GoEnv == "test" {
		isVerbose = true
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(PgURL)))
	db := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())

	db.AddQueryHook(NewJSONQueryHook(logger.Log, true, isVerbose))

	if err := db.Ping(); err != nil {
		log.Fatalf("⛔ database not running: %s", err.Error())
	}

	return db
}
