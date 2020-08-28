package database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/staheri14/go-playground/webapplication/model"
	"strings"
)

// PGDataBase implements DatabaseHandler interface for PostgreSQL database
type PGDataBase struct {
	//postgres database
	PGDB *pg.DB
}

// NewDataBase takes the address of a postgreSQL server, its database name, and user and password for an authorized user (with access to the database)
// initializes a connection, embeds it inside a PGDataBase struct and returns a pointer to that struct
func NewDataBase(addr, databaseName string) (*PGDataBase, error) {
	// connect to the postgres database with the given information
	pgdb := pg.Connect(&pg.Options{
		Addr:     addr,
		Database: databaseName,
	})

	db := &PGDataBase{PGDB: pgdb}

	// create a users table inside the database
	err := db.CreateUserTable()
	if err != nil {
		return db, fmt.Errorf("could not create the users table %w", err)
	}
	return db, nil
}

// CreateUserTable creates a table for users inside the database
func (p *PGDataBase) CreateUserTable() error {
	m := model.User{}

	// check whether the table already exists
	ok, _ := p.PGDB.Model(&m).Exists()
	if ok {
		//return nil indicating the table exists
		return nil
	}

	// if does not exist, create a non-temporary table
	err := p.PGDB.Model(&m).CreateTable(&orm.CreateTableOptions{
		Temp: false,
	})
	if err != nil {
		// if the table could not be created, return an error
		return fmt.Errorf("could not create the table %w", err)
	}

	//return nil indicating the table is successfully created with no error
	return nil
}

// InsertUser inserts a user with the information  given in model.User into the database
func (p *PGDataBase) InsertUser(user model.User) error {
	// insert the user
	_, err := p.PGDB.Model(&user).Insert()
	if err != nil {
		return fmt.Errorf("could not insert the user %w", err)
	}

	// return nil indicating the users is successfully inserted
	return nil
}

// UpdateUser receives an email string representing an email address
// finds the record associated with the given email address
// and updates its first name and last name
// based on the information given in the passed newuser
// returns error if something goes wrong.
func (p *PGDataBase) UpdateUser(newuser model.User, email string) error {

	// Check whether the requested email exists
	tempUser := model.User{Email: email}
	err := p.PGDB.Model(&tempUser).WherePK().Select()

	if err != nil {
		return fmt.Errorf("could not find the email %w", err)
	}

	// attempt updating the user's information (only fist name and last name) using data inside newuser
	newuser.Email = email
	_, err = p.PGDB.Model(&newuser).Column(model.FistName, model.LastName).WherePK().Update()
	if err != nil {
		return fmt.Errorf("could not update the user's infromation %w", err)
	}

	// return nil indicating the user's information is successfully updated
	return nil
}

// GetUsers returns the values of the columns corresponding to the fields of model.UserPublicList
// for all the records of the database
func (p *PGDataBase) GetUsers() (model.UserPublicList, error) {

	// users is the model of the table to look for
	var users model.UserList
	// usersInfo will hold the retrieved information of users
	var usersInfo model.UserPublicList
	// Read the table corresponding to users model
	// and populate usersInfo with the users information
	err := p.PGDB.Model(&users).Select(&usersInfo)

	// the users information are successfully retrieved
	return usersInfo, err

}

// GetOneUser returns true if the email inside cred exists in the database otherwise false
func (p *PGDataBase) AuthenticateUser(providedCred model.LoginCredentials) (bool, error) {

	var user model.User
	user.Email = providedCred.Email
	var userCred model.LoginCredentials

	// retrieve users information based on the primary key and write into userCred
	err := p.PGDB.Model(&user).WherePK().Select(&userCred)
	if err != nil {
		return false, fmt.Errorf("could not get users infromation from the database %w", err)
	}

	//check the password
	if strings.Compare(userCred.Password, providedCred.Password) != 0 {
		return false, nil
	}

	// the user's credential is successfully verified
	return true, nil
}
