package driver_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func MySQLConnection() (*sql.DB, error) {
	mysqlConn, _ := os.LookupEnv("MYSQL_CONNECTION")
	db, err := sql.Open("mysql", mysqlConn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
