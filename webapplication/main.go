package main

import (
	"github.com/gorilla/mux"
	"github.com/staheri14/go-playground/webapplication/database"
	"github.com/staheri14/go-playground/webapplication/server"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {

	// Initialize a logger
	logger, err := zap.NewProduction()
	if err != nil {
		// stop the execution if error happens
		panic(err)
	}

	defer func() {
		err := logger.Sync()
		if err != nil {
			// stop the execution if error happens
			panic(err)
		}
	}()
	sugar := logger.Sugar()

	// Initialize a database
	database, err := database.NewDataBase("localhost:5432", "postgres")
	if err != nil {
		panic(err)
	}

	defer func() {
		database.PGDB.Close()
		if err != nil {
			// stop the execution if error happens
			panic(err)
		}
	}()

	// Initialize a server
	s := server.NewServer(sugar, database, []byte("secret"))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.Home).Methods("GET")
	router.HandleFunc("/signup", s.SignUp).Methods("POST")
	router.HandleFunc("/login", s.Login).Methods("POST")
	router.HandleFunc("/users", s.GetUsers).Methods("GET")
	router.HandleFunc("/users", s.UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))

}
