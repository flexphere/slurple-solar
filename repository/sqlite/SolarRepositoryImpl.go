package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/flexphere/slurple-solar/repository"
)

type SolarRepositoryImpl struct {
	db *sql.DB
}

func (r SolarRepositoryImpl) init() {
	sqlStmt := "CREATE TABLE IF NOT EXISTS result (`ts` integer PRIMARY KEY,`year` integer NOT NULL,`month` integer NOT NULL,`day` integer NOT NULL,`hour` integer NOT NULL,`generation` int  NOT NULL,`consumption` int  NOT NULL,`selling` int  NOT NULL,`buying` int  NOT NULL);"
	_, err := r.db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
}

func (r *SolarRepositoryImpl) Connect() {
	dbFile := os.Getenv("DATABASE")
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	r.db = db
	r.init()
}

func (r *SolarRepositoryImpl) Disconnect() {
	if r.db == nil {
		return
	}
	r.db.Close()
}

func (r *SolarRepositoryImpl) SaveRecords(records []repository.SolarRecord) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	stmt, err := tx.Prepare(
		"INSERT INTO result (`ts`,`year`,`month`,`day`,`hour`,`generation`,`consumption`,`selling`,`buying`) " +
			"VALUES (?,?,?,?,?,?,?,?,?) " +
			"ON CONFLICT (`ts`) DO UPDATE SET " +
			"generation=?,consumption=?,selling=?,buying=?",
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer stmt.Close()

	for i, record := range records {
		_, err = stmt.Exec(
			record.TS,
			record.Year,
			record.Month,
			record.Day,
			i,
			record.Generation,
			record.Consumption,
			record.Selling,
			record.Buying,
			record.Generation,
			record.Consumption,
			record.Selling,
			record.Buying,
		)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("Insert data for %d-%d-%d\n", records[0].Year, records[0].Month, records[0].Day)
}
