package database

import "github.com/staheri14/go-playground/webapplication/model"

// DatabaseHandler represents a handler for a database
type DatabaseHandler interface {

	// CreateUserTable creates a table for users inside the database
	CreateUserTable() error

	// InsertUser inserts a user with the information  given in model.User into the database
	InsertUser(model.User) error

	// UpdateUser receives a string representing an email address
	// finds the record associated with the given email address
	// and updates its first name and last name
	// based on the information given in the passed model.User
	UpdateUser(model.User, string) error

	// GetUsers returns the values of the columns corresponding to model.UserPublicList
	// for all the records of the database
	GetUsers() (model.UserPublicList, error)

	// GetOneUser returns true if the passed credentials model.LoginCredentials is valid (exists in the database),
	// otherwise, false
	AuthenticateUser(model.LoginCredentials) (bool, error)
}
