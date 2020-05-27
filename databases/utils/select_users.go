package utils

import (
	"fmt"

	"github.com/coffemanfp/shachat/config"
	"github.com/coffemanfp/shachat/databases"
	"github.com/coffemanfp/shachat/models"
	db "upper.io/db.v3"
)

// SelectUsers - Consults a users list.
func SelectUsers(usersIDs []int) (users []models.User, err error) {
	settings, err := config.GetSettings()
	if err != nil {
		return
	}

	dbConn, err := databases.Get()
	if err != nil {
		return
	}

	usersTable := dbConn.Collection("users")

	if len(usersIDs) == 0 {
		err = usersTable.Find().Limit(settings.MaxElementsPerEP).All(&users)
		if err != nil {
			err = fmt.Errorf("failed to querying to the database:\n%s", err)
		}
		return
	}

	var stmtConditions []db.Compound

	for _, id := range usersIDs {
		stmtConditions = append(stmtConditions, db.Cond{"id": id})
	}

	err = usersTable.Find(db.Or(stmtConditions...)).All(&users)
	if err != nil {
		err = fmt.Errorf("failed to querying to the database:\n%s", err)
	}
	return
}
