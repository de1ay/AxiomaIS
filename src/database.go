package src

import (
	"database/sql"
	"axioma/conf"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var Connection DatabaseConnection

type DatabaseConnection struct {
	Connection *sql.DB
}

func (dbConnection *DatabaseConnection) getSource() string {
	source := conf.MYSQL_LOGIN + ":"
	source += conf.MYSQL_PASSWORD + "@tcp("
	source += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	source += conf.MYSQL_DB_NAME
	return source
}

func (dbConnection *DatabaseConnection) Connect() {
	var err error
	source := dbConnection.getSource()
	dbConnection.Connection, err = sql.Open("mysql", source); if err != nil {
		log.Fatal(err)
	}
	dbConnection.Connection.SetMaxOpenConns(conf.MYSQL_MAX_OPEN_CONNECTIONS)
	dbConnection.Connection.SetMaxIdleConns(conf.MYSQL_MAX_IDLE_CONNECTIONS)
}