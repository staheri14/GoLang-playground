package server

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/staheri14/go-playground/webapplication/mock"
	"github.com/staheri14/go-playground/webapplication/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

// TestGetUsers tests the happy path of getUsers request for an http request with a valid JWT token
func TestServer_GetUsers(t *testing.T) {

	// initialize a logger
	logger, err := zap.NewProduction()
	require.NoError(t, err)

	sugar := logger.Sugar()

	// set a mock database handler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock.NewMockDatabaseHandler(ctrl)

	// create a fake list of users
	plist := model.UserPublicList{{Email: "user1@email.com", FistName: "user1 fistname", LastName: "user1 lastname"}}
	// database shall return this list when gets called
	db.EXPECT().GetUsers().Return(plist, nil)

	// init the server
	s := NewServer(sugar, db, []byte{})

	// create a fake response
	w := NewFakeResponse(t)

	//make a request with an empty payload
	req, err := http.NewRequest("Get", "http://localhost:8081/", bytes.NewBuffer(make([]byte, 0)))
	require.NoError(t, err)

	// create a valid token
	token, err := s.createMacToken("user1@email.com")
	require.NoError(t, err)
	// set a valid token inside the request's x-authentication-token header
	req.Header.Set(model.AuthenticationHeader, "Bearer "+token)

	// call the get users method, no error is expected
	err = s.getUsers(w, req)
	assert.NoError(t, err)
	// the response header should contain http.StatusOK
	assert.Equal(t, w.status, http.StatusOK)

}

// TestServer_SignUp tests a happy path for sign up request where the user's email is valid and unique
func TestServer_SignUp(t *testing.T) {

	// initialize a logger
	logger, err := zap.NewProduction()
	require.NoError(t, err)
	sugar := logger.Sugar()

	// set a mock database handler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock.NewMockDatabaseHandler(ctrl)

	// create a user who will get inserted inside the database
	newuser := model.User{Email: "user1@email.com", FistName: "user1 fistname", LastName: "user1 lastname", PassWord: "123"}
	// database shall return nil when its InsertUser gets called on newuser indicating a successful error-free insertion
	db.EXPECT().InsertUser(newuser).Return(nil)

	// init the server
	s := NewServer(sugar, db, []byte{})

	// create a fake response
	w := NewFakeResponse(t)

	// prepare the body of the request
	body, err := json.Marshal(newuser)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:8081/", bytes.NewBuffer(body))
	require.NoError(t, err)

	// call the sign up method, no error is expected
	err = s.signUp(w, req)
	assert.NoError(t, err)
	// the response header should contain http.StatusOK
	assert.Equal(t, w.status, http.StatusOK)
}

// TestServer_Login tests the happy of login request where the user provided email and password are correct
func TestServer_Login(t *testing.T) {

	// initialize a logger
	logger, err := zap.NewProduction()
	require.NoError(t, err)
	sugar := logger.Sugar()

	// set a mock database handler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock.NewMockDatabaseHandler(ctrl)

	// create valid LoginCredentials
	userCred := model.LoginCredentials{Email: "user1@email.com", Password: "123"}
	// database shall return true when its AuthenticateUser gets called on userCred indicating a successful login
	db.EXPECT().AuthenticateUser(userCred).Return(true, nil)

	// init the server
	s := NewServer(sugar, db, []byte{})

	// create a fake response writer
	w := NewFakeResponse(t)

	// prepare the body of the request
	body, err := json.Marshal(&userCred)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:8081/", bytes.NewBuffer(body))
	require.NoError(t, err)

	// call the login method, no error is expected
	err = s.login(w, req)
	assert.NoError(t, err)

	// the response header should contain http.StatusOK
	assert.Equal(t, w.status, http.StatusOK)
}

// TestServer_UpdateUser tests the happy of update user request where the user provided token is valid
func TestServer_UpdateUser(t *testing.T) {

	// initialize a logger
	logger, err := zap.NewProduction()
	require.NoError(t, err)
	sugar := logger.Sugar()

	// set a mock database handler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock.NewMockDatabaseHandler(ctrl)

	userInfo := model.User{FistName: "user1 fistname", LastName: "user1 lastname"}
	db.EXPECT().UpdateUser(userInfo, "user1@email.com").Return(nil)

	// init the server
	s := NewServer(sugar, db, []byte{})

	// create a fake response writer
	w := NewFakeResponse(t)

	// prepare the body of the request
	body, err := json.Marshal(userInfo)
	require.NoError(t, err)
	req, err := http.NewRequest("PUT", "http://localhost:8081/", bytes.NewBuffer(body))
	require.NoError(t, err)
	//create a valid token
	token, err := s.createMacToken("user1@email.com")
	require.NoError(t, err)
	// embed the token insider the x-authentication-token header
	req.Header.Set(model.AuthenticationHeader, "Bearer "+token)

	// call the updateUser method, no error is expected
	err = s.updateUser(w, req)
	assert.NoError(t, err)

	// the response header should contain http.StatusOK
	assert.Equal(t, w.status, http.StatusOK)
}

// FakeResponse is a mocked fake http response taken from https://gist.github.com/karlseguin/5128461
type FakeResponse struct {
	t       *testing.T
	headers http.Header
	body    []byte
	status  int
}

func NewFakeResponse(t *testing.T) *FakeResponse {
	return &FakeResponse{
		t:       t,
		headers: make(http.Header),
	}
}

func (r *FakeResponse) Header() http.Header {
	return r.headers
}

func (r *FakeResponse) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

func (r *FakeResponse) Assert(status int, body string) {
	if r.status != status {
		r.t.Errorf("expected status %+v to equal %+v", r.status, status)
	}
	if string(r.body) != body {
		r.t.Errorf("expected body %+v to equal %+v", string(r.body), body)
	}
}
