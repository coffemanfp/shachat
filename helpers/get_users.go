package helpers

// GetUsers - ...
func GetUsers(usersIDs []int) (users []models.User, err error) {
	users, err = dbu.SelectUsers(usersIDs)
	if err != nil {
		return
	}

	var usersContacts []models.User

	for _, user := range users {
		contacts, err := dbu.SelectUserContacts(user.ID)
		if err != nil {
			return
		}

		usersContacts.Contacts = append(usersContacts.Contacts, contacts)
	}

	users.Contacts = usersContacts
	return
}