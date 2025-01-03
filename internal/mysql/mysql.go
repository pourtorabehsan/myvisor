package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Open(user string, password string, host string, port int) (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port))
}

// Hostname returns the hostname of the MySQL server
func Hostname(db *sql.DB) string {
	query := "SELECT @@hostname"
	row := db.QueryRow(query)
	var hostname string
	err := row.Scan(&hostname)
	if err != nil {
		return ""
	}
	return hostname
}
