package utils_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/coffemanfp/shachat/config"
	dbu "github.com/coffemanfp/shachat/databases/utils"
	_ "github.com/lib/pq"
	"upper.io/db.v3/postgresql"
	_ "upper.io/db.v3/postgresql"
)

func TestSelectUsers(t *testing.T) {
	err := config.SetSettingsByFile("../../config.yaml")
	if err != nil {
		t.Errorf("unexpected error on settings:\n%s", err)
	}

	_, err = config.GetSettings()
	if err != nil {
		t.Errorf("unexpected error getting the settings:\n%s", err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected error opening a stub database connection:\n%s", err)
	}
	defer db.Close()

	// columns := []string{"id", "link_id"}

	// mock.ExpectQuery("SELECT (.+) FROM users LIMIT ?").
	// 	WithArgs(settings.MaxElementsPerEP)

	dbConn, err := postgresql.New(db)
	if err != nil {
		t.Errorf("unexpected error getting a database mock:\n%s", err)
	}

	users, err := dbu.SelectUsers(dbConn, []int{})
	if err != nil {
		t.Errorf("unexpected error getting the users from database:\n%s", err)
	}

	t.Log("----------------------")
	t.Log(users)
	t.Log("----------------------")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestExecSelectUsers(t *testing.T) {
	err := config.SetSettingsByFile("../../config.yaml")
	if err != nil {
		t.Errorf("unexpected error on settings:\n%s", err)
	}

	settings, err := config.GetSettings()
	if err != nil {
		t.Errorf("unexpected error getting the settings:\n%s", err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected error opening a stub database connection:\n%s", err)
	}
	defer db.Close()

	query := fmt.Sprintf(`
		SELECT
			id, link_id, created_on, updated_on, last_login
		FROM
			users
		LIMIT
			%d
	`, settings.MaxElementsPerEP)

	mock.ExpectQuery(query)
	mock.ExpectExec(query)
	mock.ExpectPrepare(query)

	db.QueryRow(query)
	db.Exec(query)
	db.Prepare(query)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
