# Instructions
- Run `docker-compose`, it will run containers for the PostgreSQL server and the http server.
- The `docker-compose` builds up two containers, `db` and `http-server`, where `db` is the PostgreSQL, and 
  is accessible to `http-server` on `db:5432`.
- The http server is accessible via `localhost:8080` 
- You may install Postman (https://www.postman.com/) to interact with the http server
    - To include a token in a request, inside the Postman software, click the "Authorization" tab, and from the "TYPE" drop-down menu, select the "Bearer Token", then copy the token in the "Token" field
    - Note that for the http server to be able to read tokens from Postman requests, the token header must be changed from `"x-authentication-token"` to `"authorization"`, you can do this by 
    setting the `AuthenticationHeader` variable (in model/user.go file) to `"authorization"`
    
    
