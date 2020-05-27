package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/coffemanfp/shachat/config"
	dbu "github.com/coffemanfp/shachat/databases/utils"
	"github.com/coffemanfp/shachat/models"
	"github.com/coffemanfp/shachat/utils"
	"github.com/gin-gonic/gin"
)

// GetUsers - Get a users array.
func GetUsers(c *gin.Context) {

	var m models.ResponseMessage
	var usersIDs []int

	queryUsersIDs := c.Query("users")

	if queryUsersIDs != "" {
		stringArrayUsersIDs, err := utils.StringToArray(queryUsersIDs)
		if err != nil {
			m.Error = errors.New("User id invalid")

			c.JSON(http.StatusBadRequest, m)
			return
		}

		usersIDs, err = utils.StringArrayToIntArray(stringArrayUsersIDs)
		if err != nil {
			m.Error = errors.New("User id invalid")

			c.JSON(http.StatusBadRequest, m)
			return
		}
	}

	settings, err := config.GetSettings()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if len(usersIDs) == 0 {
		m.Message = fmt.Sprintf(
			"Alert! The limit was automatically set to %d",
			settings.MaxElementsPerEP,
		)
	} else if len(usersIDs) > settings.MaxElementsPerEP {
		m.Message = fmt.Sprintf(
			"Alert! You have exceeded the limit of elements. The first %d items were returned",
			settings.MaxElementsPerEP,
		)

		usersIDs = usersIDs[:settings.MaxElementsPerEP]
	}

	users, err := dbu.SelectUsers(usersIDs)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	m.Content = users

	c.JSON(http.StatusOK, m)
}
