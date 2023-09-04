package main

import (
	"fmt"
	"net/http"
	"nlearn/database"
	"nlearn/handlers"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver

	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	db := database.DB()
	defer db.Close()
	fmt.Fprint(w, "testing data goes here")

}

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/home", homepage).Methods("GET")            // Pass a reference to the homepage function
	router.HandleFunc("/add", handlers.Createuser).Methods("POST") // Pass a reference to the homepage function
	http.ListenAndServe(":8080", nil)

}
