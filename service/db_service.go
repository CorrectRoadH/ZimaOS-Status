package service

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/CorrectRoadH/ZimaOS-Status/codegen"
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

func convertTimestamp(ts string) (int64, error) {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return 0, err
	}
	// 转换为Go的time.Time，然后格式化为SQLite可接受的格式

	return i, nil
}

func (s *DBService) LatestCPUUsage() (codegen.CpuInfo, error) {
	cpu := codegen.CpuInfo{}

	sqlStmt := `SELECT * FROM CPUData ORDER BY timestamp DESC LIMIT 1`
	rows, err := s.db.Query(sqlStmt)
	if err != nil {
		fmt.Println(err)
		return cpu, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var timestamp string
		var percent float64
		err = rows.Scan(&id, &timestamp, &percent)
		if err != nil {
			fmt.Println(err)
			return cpu, err
		}
		cpu = codegen.CpuInfo{
			Percent:   float32(percent),
			Timestamp: timestamp,
		}
	}
	return cpu, nil
}

func (s *DBService) LatestMemUsage() (codegen.MemoryInfo, error) {
	mem := codegen.MemoryInfo{}
	sqlStmt := `SELECT * FROM MemData ORDER BY timestamp DESC LIMIT 1`
	rows, err := s.db.Query(sqlStmt)
	if err != nil {
		fmt.Println(err)
		return mem, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var timestamp string
		var percent float64
		err = rows.Scan(&id, &timestamp, &percent)
		if err != nil {
			fmt.Println(err)
			return mem, err
		}
		mem = codegen.MemoryInfo{
			Percent:   float32(percent),
			Timestamp: timestamp,
		}
	}
	return mem, nil
}

func (s *DBService) QueryCPUUsageHistory(start string, end string) ([]codegen.CpuInfo, error) {
	startTime, err := convertTimestamp(start)
	if err != nil {
		return nil, err
	}
	endTime, err := convertTimestamp(end)
	if err != nil {
		return nil, err
	}

	startTimeSQL := time.Unix(startTime, 0).UTC().Format("2006-01-02 15:04:05")
	if err != nil {
		return nil, err
	}
	endTimeSQL := time.Unix(endTime, 0).UTC().Format("2006-01-02 15:04:05")
	if err != nil {
		return nil, err
	}

	sqlStmt := `SELECT * FROM CPUData WHERE timestamp BETWEEN ? AND ?`
	rows, err := s.db.Query(sqlStmt, startTimeSQL, endTimeSQL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	history := []codegen.CpuInfo{}
	for rows.Next() {
		var id int
		var timestamp string
		var percent float64
		err = rows.Scan(&id, &timestamp, &percent)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		history = append(history, codegen.CpuInfo{
			Percent:   float32(percent),
			Timestamp: timestamp,
		})
	}
	return history, nil
}

func (s *DBService) QueryMemUsageHistory(start string, end string) ([]codegen.MemoryInfo, error) {
	startTime, err := convertTimestamp(start)
	if err != nil {
		return nil, err
	}
	endTime, err := convertTimestamp(end)
	if err != nil {
		return nil, err
	}

	startTimeSQL := time.Unix(startTime, 0).UTC().Format("2006-01-02 15:04:05")
	if err != nil {
		return nil, err
	}
	endTimeSQL := time.Unix(endTime, 0).UTC().Format("2006-01-02 15:04:05")
	if err != nil {
		return nil, err
	}

	sqlStmt := `SELECT * FROM MemData WHERE timestamp BETWEEN ? AND ?`
	rows, err := s.db.Query(sqlStmt, startTimeSQL, endTimeSQL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	history := []codegen.MemoryInfo{}
	for rows.Next() {
		var id int
		var timestamp string
		var percent float64
		err = rows.Scan(&id, &timestamp, &percent)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		history = append(history, codegen.MemoryInfo{
			Percent:   float32(percent),
			Timestamp: timestamp,
		})
	}
	return history, nil
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
