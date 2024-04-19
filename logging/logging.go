package logging

import (
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.JSONFormatter{}
}

// LogUserCreation logs user creation event
func LogUserCreation(name, game string) {
	logger.WithFields(logrus.Fields{
		"event":     "UserCreation",
		"operation": "creating user",
		"createdAt": time.Now(),
		"Name":      name,
		"Game":      game,
	}).Info("User was created")
}

// LogUsersIndex logs the event of retrieving all users
func LogUsersIndex() {
	logger.WithFields(logrus.Fields{
		"event": "UsersIndex",
	}).Info("Retrieved all users")
}

// LogUsersShow logs the event of retrieving a single user
func LogUsersShow(id string) {
	logger.WithFields(logrus.Fields{
		"event": "UsersShow",
		"id":    id,
	}).Info("Retrieved user")
}

// LogUserDeletion logs user deletion event
func LogUserDeletion(name, game string) {
	logger.WithFields(logrus.Fields{
		"event":     "UserDeletion",
		"operation": "deleting user",
		"deletedAt": time.Now(),
		"Name":      name,
		"Game":      game,
	}).Info("User was deleted")
}

// LogUserUpdate logs user update event
func LogUserUpdate(name, game string) {
	logger.WithFields(logrus.Fields{
		"event":     "UserUpdate",
		"operation": "updating user",
		"updatedAt": time.Now(),
		"Name":      name,
		"Game":      game,
	}).Info("User was updated")
}

// LogFilteringSortingPaginating logs filtering, sorting, and pagination event
func LogFilteringSortingPaginating(filter, sort string, page int) {
	logger.WithFields(logrus.Fields{
		"event":     "FilteringSortingPaginating",
		"operation": "filtering, sorting, and paginating",
		"filter":    filter,
		"sort":      sort,
		"page":      page,
	}).Info("Performed filtering, sorting, and pagination")
}
