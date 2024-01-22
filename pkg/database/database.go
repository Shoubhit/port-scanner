package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type Report struct {
	ID          int       `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	IPAddress   string    `json:"ipAddress"`
	Username    string    `json:"username"`
	PortScan    string    `json:"portScan"`
	Protocol    string    `json:"protocol"`
	DNSAnalysis string    `json:"dnsAnalysis"`
}

var db *sql.DB
var logger *logrus.Logger

func init() {
	initLogger()
	initDB()
}

func initLogger() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(logrus.StandardLogger().Out)
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./security_analysis.db")
	if err != nil {
		logger.WithError(err).Fatal("Failed to open database")
	}
	createTable()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS reports (
			id INTEGER,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			ipAddress TEXT,
			username TEXT,
			portScan TEXT,
			protocol TEXT,
			dnsAnalysis TEXT,
			PRIMARY KEY(id AUTOINCREMENT)
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		logger.WithError(err).Fatal("Failed to create table")
	}
}

func AddReport(report Report) error {
	query := "INSERT INTO reports (ipAddress, username, portScan, protocol, dnsAnalysis) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, report.IPAddress, report.Username, report.PortScan, report.Protocol, report.DNSAnalysis)
	if err != nil {
		logger.WithError(err).Error("Failed to insert report into database")
	}
	return err
}

func GetAllReports() ([]Report, error) {
	var reports []Report
	rows, err := db.Query("SELECT id, timestamp, ipAddress, username, portScan, protocol, dnsAnalysis FROM reports")
	if err != nil {
		logger.WithError(err).Error("Failed to execute SELECT query on database")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var report Report
		err := rows.Scan(&report.ID, &report.Timestamp, &report.IPAddress, &report.Username, &report.PortScan, &report.Protocol, &report.DNSAnalysis)
		if err != nil {
			logger.WithError(err).Error("Failed to scan row from SELECT query result")
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}
