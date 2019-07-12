package main

import (
	"EndToEnd/Jwt/models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/user_account")
	if err != nil {
		log.Fatal("error in db connection")
	}

	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")

	router.HandleFunc("/protected", TokenVerifyMiddleWare(ProtectedEndpoint)).Methods("GET")
	err = http.ListenAndServe(":8000", router)

	if err != nil {
		log.Fatal("Error occured in starting server")
		os.Exit(1)
	}
	log.Println("Server started")
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Protected endpoint invkoed")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {

	return nil
}

func errorMessage(w http.ResponseWriter, status int, message string) {
	var errorMessage models.Error
	errorMessage.Message = message
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorMessage)
}

func login(writer http.ResponseWriter, request *http.Request) {

	_, err := writer.Write([]byte("successfully called login"))
	if err != nil {
		log.Fatal("error occurred in wring response")
	}
}

func signup(writer http.ResponseWriter, request *http.Request) {

	var user models.User
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Fatal("error occurred in de-serializing")
	}
	if user.Email == "" {
		errorMessage(writer, http.StatusBadRequest, "Email is not valid")
		return
	}
	if user.Password == "" {
		errorMessage(writer, http.StatusBadRequest, "password is not valid")
		return
	}

	password, errp := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errp != nil {
		errorMessage(writer, http.StatusInternalServerError, err.Error())
	}
	user.Password = string(password)

	stmt := "insert into users(email,password) values($1,$2)"
	db.QueryRow(stmt, user.Email, user.Password)
	_ = json.NewEncoder(writer).Encode(&user)

}
