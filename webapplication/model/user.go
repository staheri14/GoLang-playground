package model

var Email = "email"
var FistName = "firstname"
var LastName = "lastname"
var AuthenticationHeader="x-authentication-token"

type User struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	Email     string   `json:"email" pg:"email,pk"`
	PassWord  string   `json:"password" pg:"password"`
	FistName  string   `json:"firstName" pg:"firstname"`
	LastName  string   `json:"lastName" pg:"lastname"`
}

type UserList []*User

type UserPublicList []*struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	Email     string   `json:"email" pg:"email,pk"`
	FistName  string   `json:"firstName" pg:"firstname"`
	LastName  string   `json:"lastName" pg:"lastname"`
}

type UserPublicListWrapper struct {
	Users UserPublicList `json:"users"`
}

type LoginCredentials struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
}

type TokenCredentials struct {
	Token string `json:"token"`
}

// Claims struct that will be encoded to a JWT.
type Claims struct {
	Email string `json:"email"`
}
