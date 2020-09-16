## Project Description
The current project is the implementation of REST API for a web application.

The application has the signup, login endpoints as well as an additional resource `users` that can be accessed if a user is logged in.

To determine whether a request is from a logged in user or not, Json Web Tokens (https://jwt.io/) are utilized. The frontend will be sending requests with the JWT in the `x-authentication-token` header.

PostgreSQL is used for the database with the following Golang wrapper https://github.com/go-pg/pg.

An overview of the application is shown in the following figure.
![system overview](https://github.com/staheri14/go-playground/blob/master/webapplication/photos/overview.png)


## API Specs

### `POST /signup`
Endpoint to create an user row in postgres db. The payload should have the following fields:

```json
{
  "email": "test@axiomzen.co",
  "password": "axiomzen",
  "firstName": "Alex",
  "lastName": "Zimmerman"
}
```

where `email` is an unique key in the database.

The response body returns a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token" 
}
```

### `POST /login`
Endpoint to log an user in. The payload has the following fields:

```json
{
  "email": "test@axiomzen.co",
  "password": "axiomzen"
}
```

The response body returns a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token"
}
```

### `GET /users`
Endpoint to retrieve a json of all users. This endpoint has a valid `x-authentication-token` header to be passed in with the request.

The response body looks like:
```json
{
  "users": [
    {
      "email": "test@axiomzen.co",
      "firstName": "Alex",
      "lastName": "Zimmerman"
    }
  ]
}
```

### `PUT /users`
Endpoint to update the current user `firstName` or `lastName` only. This endpoint requires a valid `x-authentication-token` header to be passed in and it only updates the user of the JWT being passed in. The payload has the following fields:

```json
{
  "firstName": "NewFirstName",
  "lastName": "NewLastName"
}
```

The response can body can be empty.

## Final Notes
The current implementation only showcases my skills in development of a production-grade code. However, there are various improvements to be considered. 
The possible improvements include (but not limited to) the followings: 
- Having a password protected  connection to the database server
- Storing salted hash of passwords inside the database
- Narrowing down the errors and providing specific instructions to both the client and the server to handle errors
- Using a secure MAC key for the server
- Validating user's inputs (e.g., the format of email)  during the sign up procedure
- Including un-happy paths in the tests (the current tests only assess the happy paths)
- ...


