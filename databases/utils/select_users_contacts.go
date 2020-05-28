package utils

// SelectUserContacts - Consults a user contacts list.
// func SelectUserContacts(userID int, offset int, limit int) (contacts []models.Contact, err error) {
// 	settings, err := config.GetSettings()
// 	if err != nil {
// 		return
// 	}

// 	dbConn, err := databases.Get()
// 	if err != nil {
// 		return
// 	}

// 	contactsTable := dbConn.Collection("contacts")
// 	err = contactsTable.Find("user_id", userID).Limit(settings.MaxElementsPerEP).All(&contacts)
// 	if err != nil {
// 		err = fmt.Errorf("failed to querying to the database:\n%s", err)
// 	}

// 	var stmtConditions []db.Compound

// 	for _, id := range usersIDs {
// 		stmtConditions = append(stmtConditions, db.Cond{"id": id})
// 	}

// 	err = contactsTable.Find(db.Or(stmtConditions...)).All(&users)
// 	if err != nil {
// 		err = fmt.Errorf("failed to querying to the database:\n%s", err)
// 	}
// 	return
// }
