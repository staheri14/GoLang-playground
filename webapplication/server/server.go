package server

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"github.com/staheri14/go-playground/webapplication/database"
	"github.com/staheri14/go-playground/webapplication/model"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type Server struct {
	logger   *zap.SugaredLogger
	MacKey   []byte
	database database.DatabaseHandler
}

// NewServer initializes and returns an instance of a Server struct
func NewServer(logger *zap.SugaredLogger, db database.DatabaseHandler, macKey []byte) *Server {
	s := &Server{
		logger:   logger,
		database: db,
		MacKey:   macKey,
	}
	return s
}

// GetUsers receives a request r that contains a JWT token embedded in the x-authentication-token header
// If the token is valid, then the list of all signed up users gets
// encoded in the body of the response w
// an appropriate HTTP status code also gets set in w
func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	// process the request
	err := s.getUsers(w, r)
	if err != nil {
		s.logger.Error("could not retrieve user", "error:", err.Error())
		return
	}

	s.logger.Info("the Get users is done successfully")

}

// getUsers receives a request r that contains a JWT token embedded in the x-authentication-token header
// If the token is valid, then the list of all signed up users gets
// encoded in the body of the response w and returns nil
// Otherwise, returns an error
// an appropriate HTTP status code also gets set in w
func (s *Server) getUsers(w http.ResponseWriter, r *http.Request) error {
	authorizationToken := r.Header.Get(model.AuthenticationHeader)

	// Authenticate the token
	_, err := s.AuthorizeToken(authorizationToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return fmt.Errorf("could not authorize the user: %w", err)
	}

	// Read the list of users from the database
	userList, err := s.database.GetUsers()
	if err != nil {
		// if the list of users could not be retrieved return an error
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("could not retrieve users from the database: %w", err)
	}

	// Encode the list of users inside the response w
	err = json.NewEncoder(w).Encode(model.UserPublicListWrapper{Users: userList})
	if err != nil {
		// if the encoding does not take place, return an error
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("could not encode user list: %w", err)
	}

	// The request is processed successfully
	w.WriteHeader(http.StatusOK)
	return nil
}

// UpdateUser receives a request r containing a user's  firstName or lastName in the payload
// and a JWT token in the x-authentication-token header
// If the token is valid, then the function updates the user of the JWT being passed in.
// an appropriate HTTP status code also gets set in w
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	err := s.updateUser(w, r)
	if err != nil {
		s.logger.Error("could not process the update request error:", err.Error())
		return
	}
	s.logger.Info("the update is done successfully")

}

// updateUser receives a request r containing a user's  firstName or lastName in the payload
// and a JWT token in the x-authentication-token header.
// If the token is valid, then the function updates the user of the JWT being passed in and returns nil
// Otherwise, returns an error
// an appropriate HTTP status code also gets set in w
func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) error {

	// Read and decode the payload of the request r into a User struct
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If the structure of the body is wrong, set an HTTP  status code  inside w
		w.WriteHeader(http.StatusBadRequest)
		// return an error
		return fmt.Errorf("could not decode the body of the request: %w", err)
	}

	// Extract the x-authentication-token header
	authorizationHeader := r.Header.Get(model.AuthenticationHeader)
	//Authorize the token embedded in the authorization header
	token, err := s.AuthorizeToken(authorizationHeader)
	if err != nil {
		//If the authorization does not pass successfully, set an HTTP  status code  inside w
		w.WriteHeader(http.StatusUnauthorized)
		// return an error
		return fmt.Errorf("could not authorize the authorization header: %w", err)
	}

	// Decode the claim part of the token into a Claim struct
	var claim model.Claims
	err = mapstructure.Decode(token.Claims, &claim)
	if err != nil {
		// If the structure of the body is wrong, set an HTTP  status code  inside w
		w.WriteHeader(http.StatusBadRequest)
		// return an error
		return fmt.Errorf("could not extract the claimed request from the token: %w", err)
	}

	// Update the record of the claimed email based on the new information inside the user
	err = s.database.UpdateUser(user, claim.Email)
	if err != nil {
		// If the information could not be updated, set an HTTP  status code  inside w
		w.WriteHeader(http.StatusBadRequest)
		// return an error
		return fmt.Errorf("could not update the user's record in the database: %w", err)
	}

	// The update is done successfully
	w.WriteHeader(http.StatusOK)
	return nil
}

// Login receives a request r with the user's email and password inside the payload
// if the (email,password) combination is correct then a JWT token gets embedded inside the response w
// an appropriate HTTP status code also gets set in w
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {

	err := s.login(w, r)
	if err != nil {
		s.logger.Error("could not process the login request error:", err.Error())
		return
	}
	s.logger.Info("the login is done successfully")
}

// login receives a request r with the user's email and password inside the payload
// if the (email,password) combination is correct then a JWT token gets embedded inside the response w
// and the function returns nil
// otherwise returns an error
// an appropriate HTTP status code also gets set in w
func (s *Server) login(w http.ResponseWriter, r *http.Request) error {

	// decode the payload into a LoginCredentials
	var creds model.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, set an HTTP status code  inside w
		w.WriteHeader(http.StatusBadRequest)
		// return an error
		return fmt.Errorf("could not decode the body of the request: %w", err)
	}

	// Authenticate the user's email and password (embedded inside the creds)
	ok, err := s.database.AuthenticateUser(creds)
	if err != nil {
		// If an error happens, set an HTTP status code inside w
		w.WriteHeader(http.StatusInternalServerError)
		// return an error
		return fmt.Errorf("could not retrieve the user's record from the database: %w", err)
	}
	if !ok {
		// if the authentication of user's email and password  failed, set an HTTP status code inside w
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	var t model.TokenCredentials
	t.Token, err = s.createMacToken(creds.Email)
	if err != nil {
		// if the token creation fails, then set an HTTP status code inside w
		w.WriteHeader(http.StatusInternalServerError)
		// return an error
		return fmt.Errorf("could not create a token: %w", err)
	}

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		// if the token encoding fails, then set an HTTP status code inside w
		w.WriteHeader(http.StatusInternalServerError)
		// return an error
		return fmt.Errorf("could not encode the token: %w", err)
	}

	// The login is done successfully
	w.WriteHeader(http.StatusOK)
	return nil

}

// SignUp receives user's email, first name, last name and password inside the payload of a request r
// and adds the user to the database and encodes a valid token inside the response w
// an appropriate HTTP status code also gets set in w
func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	err := s.signUp(w, r)
	if err != nil {
		s.logger.Error("could not process the sign up request error:", err.Error())
		return
	}
	s.logger.Info("the sign up is done successfully")
}

// signUp receives user's email, first name, last name and password inside the payload of a request r
// and adds the user to the database, encodes a valid token inside the response w and returns nil
// if any error happens, return an error
// an appropriate HTTP status code also gets set in w
func (s *Server) signUp(w http.ResponseWriter, r *http.Request) error {

	// Decode the payload into a User instance
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If the structure of the body is wrong, set an HTTP status code inside w
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("could not decode the body of the request: %w", err)
	}

	// create a token
	var token model.TokenCredentials
	token.Token, err = s.createMacToken(user.Email)
	if err != nil {
		// if the token creation fails, set an HTTP status code inside w
		w.WriteHeader(http.StatusInternalServerError)
		// return an error
		return fmt.Errorf("could not create a token: %w", err)
	}

	// adds the user into the database
	err = s.database.InsertUser(user)
	if err != nil {
		// if the insertion fails, set an HTTP status code inside response w
		w.WriteHeader(http.StatusInternalServerError)
		//return an error
		return fmt.Errorf("the user could not be added to the database: %w", err)
	}

	// encode the token inside the response w
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		// if the encoding fails, set an HTTP status code inside response w
		w.WriteHeader(http.StatusInternalServerError)
		//return an error
		return fmt.Errorf("could not encode the token: %w", err)
	}

	// the sing up is done successfully
	w.WriteHeader(http.StatusOK)
	return nil

}

// createMacToken creates a jwt token with the email embedded in the claim part
func (s *Server) createMacToken(email string) (string, error) {

	// create a jwt token
	token := jwt.New(jwt.SigningMethodHS256)

	// fill in the claimed email inside the token
	token.Claims = &jwt.MapClaims{
		model.Email: email}

	// sign the token
	tokenString, err := token.SignedString(s.MacKey)
	if err != nil {
		// if the sign does not success, return an empty token and an error
		return "", fmt.Errorf("could not sign the token %w \n", err)
	}

	return tokenString, nil

}

// AuthorizeToken parses the authorizationHeader and authorizes the token therein
// An appropriate  error message is returned in case of authorization failure
func (s *Server) AuthorizeToken(authorizationHeader string) (*jwt.Token, error) {
	// if the authorizationHeader is empty
	if len(authorizationHeader) < 1 {
		return nil, fmt.Errorf("the authorization header could not be found")
	}

	// Check whether the header has two space separated parts
	bearerToken := strings.Split(authorizationHeader, " ")
	if len(bearerToken) != 2 {
		return nil, fmt.Errorf("invalid format of the authorization token")
	}

	// Parse, validate, and return a token
	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("couldn't parse the token")
		}
		return s.MacKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid format of authorization token")
	}

	// Check the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token could not be verified")
	}

	return token, nil
}
