package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/horacioskrp/kitoo/cmd/api"
	"github.com/horacioskrp/kitoo/config"
	"github.com/horacioskrp/kitoo/db"
)

func main() {
	db, err := db.NewMysSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPAssword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBNAme,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.RUN(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected")

}
