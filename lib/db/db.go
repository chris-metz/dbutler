package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type DbHandler struct {
	db *sql.DB
}

func NewDbHandler() *DbHandler {
	return &DbHandler{
		db: getOrCreateDb(),
	}
}

func (dh *DbHandler) Shutdown() {
	fmt.Println("shutting down ...")
	dh.db.Close()
}

func (dh *DbHandler) ReCreateSchema() {
	fmt.Print("Re-Creating db schema ... ")
	createConnectionTable := `
    DROP TABLE IF EXISTS connection;
    CREATE TABLE IF NOT EXISTS connection (
      id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
      name text,
      conn_string text
    );
  `
	_, err := dh.db.Exec(createConnectionTable)
	if err != nil {
		panic(err)
	}
	fmt.Print("done\n")
}

func (dh *DbHandler) SeedDatabase() {
	fmt.Print("Seeding database ... ")
	insertConn := `INSERT INTO connection(name, conn_string) VALUES(?,?)`
	stmt := dh.createStatement(insertConn)
	defer stmt.Close()
	stmt.Exec("local-postgres", "popstgres://...")
	fmt.Print("done\n")
}

func (dh *DbHandler) createStatement(sql string) *sql.Stmt {
	stmt, err := dh.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	return stmt
}

func getOrCreateDb() *sql.DB {
	fileName := "/home/chris/dbutler.db"
	db, err := sql.Open("sqlite", fileName)
	if err != nil {
		panic(err)
	}
	return db
}
