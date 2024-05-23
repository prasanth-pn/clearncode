package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/prasanth-pn/cleancode/pkg/config"
)

func ConnectPGDB(cnf config.Config) *sql.DB {
	fmt.Println(cnf)
	postgresURL := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v", cnf.PGDBmsName, cnf.PGPassword, cnf.PGPassword, cnf.PGHost, cnf.PgPort, cnf.PGDBName, cnf.PgSSLMode)
	db, err := sql.Open(cnf.PgDriverName, postgresURL)
	if err != nil {
		log.Fatal(err, err.Error(), "driver name", cnf.PgDriverName, "postgres url", postgresURL)
	}

	/*

		maxOpenConnections, err := strconv.Atoi(cnf.PgServerMaxOpenConnections)
		if err != nil {
			log.Fatal("unable to set max open connections: ", err.Error())
		}
		db.SetMaxOpenConns(maxOpenConnections)

		maxIdleConnections, err := strconv.Atoi(cnf.PgServerMaxIdleConnections)
		if err != nil {
			log.Fatal("unable to set max idle connections: ", err.Error())
		}
		db.SetMaxIdleConns(maxIdleConnections)

	*/

	err = db.Ping()
	if err != nil {
		log.Fatal("not connected to postgres db: ", err.Error())
	}

	log.Println("connected to postgres db successfully!")

	return db
}
