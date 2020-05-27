package databases_test

import (
	"testing"

	"github.com/coffemanfp/shachat/config"
	"github.com/coffemanfp/shachat/databases"
)

func TestGet(t *testing.T) {
	err := config.SetSettingsByFile("../config.yaml")
	if err != nil {
		t.Error("unexpected error:\n", err)
	}

	dbConn, err := databases.Get()
	if err != nil {
		t.Error("unexpected error:\n", err)
	}
	defer databases.CloseConn()

	err = dbConn.Ping()
	if err != nil {
		t.Errorf("failed to ping to the database:%s\b", err)
	}
}

func TestMaxConns(t *testing.T) {
	err := config.SetSettingsByFile("../config.yaml")
	if err != nil {
		t.Error("unexpected error:\n", err)
	}

	dbConn, err := databases.Get()
	if err != nil {
		t.Error("unexpected error:\n", err)
	}
	defer databases.CloseConn()

	maxConns := dbConn.MaxIdleConns()

	if maxConns != 1 {
		t.Errorf("max connections (%d) invalid, expected (%d)", maxConns, 1)
	}
}
