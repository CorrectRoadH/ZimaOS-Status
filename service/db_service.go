package service

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DBService struct {
	db *sql.DB
}

func NewDBService() *DBService {
	db, err := sql.Open("sqlite3", "file:mydb.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}

	dbService := &DBService{
		db: db,
	}

	// init
	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='CPUData'").Scan(&tableName)

	if err == sql.ErrNoRows {
		dbService.Init()
	}

	return dbService
}

func (s *DBService) Init() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS CPUData (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME NOT NULL,
		percent REAL NOT NULL
	);`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	sqlStmt = `
	CREATE TABLE IF NOT EXISTS MemData (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME NOT NULL,
		percent REAL NOT NULL
	);`
	_, err = s.db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DBService initialized")
}

func (s *DBService) Query() {

}

func (s *DBService) InsertCPUData(value float64) {
	sqlStmt := `INSERT INTO CPUData (timestamp, percent) VALUES (datetime('now'), ?)`
	_, err := s.db.Exec(sqlStmt, value)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *DBService) InsertMemData(value float64) {
	sqlStmt := `INSERT INTO MemData (timestamp, percent) VALUES (datetime('now'), ?)`
	_, err := s.db.Exec(sqlStmt, value)
	if err != nil {
		fmt.Println(err)
	}
}
