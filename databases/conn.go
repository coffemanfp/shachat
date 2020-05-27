package databases

import (
	"errors"
	"fmt"

	"github.com/coffemanfp/shachat/config"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

var db sqlbuilder.Database

// Get - Get the conn to the database.
func Get() (dbConn sqlbuilder.Database, err error) {
	if db != nil {
		dbConn = db
		return
	}

	dbConn, err = OpenConn()
	if err != nil {
		return
	}

	db = dbConn
	return
}

// OpenConn - Open a conn to the database.
func OpenConn() (dbConn sqlbuilder.Database, err error) {
	settings, err := config.GetSettings()
	if err != nil {
		return
	}

	if !settings.ValidateDatabase() {
		err = errors.New(fmt.Sprint("database settings are not populated", settings))
		return
	}

	var dbSettings = postgresql.ConnectionURL{
		Host:     settings.Database.Host,
		User:     settings.Database.User,
		Password: settings.Database.Password,
		Database: settings.Database.Name,
	}

	dbConn, err = postgresql.Open(dbSettings)
	if err != nil {
		err = fmt.Errorf("error opening a database connection:\n%s", err)
		return
	}

	dbConn.SetMaxIdleConns(1)

	db = dbConn

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("error in ping to the database:\n%s", err)
	}

	db = dbConn
	return
}

// CloseConn - ...
func CloseConn() {
	if db == nil {
		return
	}

	db.Close()
}
