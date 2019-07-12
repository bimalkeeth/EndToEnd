package main

import (
	"EndToEnd/Jwt/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		barerToken := strings.Split(authHeader, " ")
		if len(barerToken) == 2 {
			authToken := barerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					return nil, fmt.Errorf("error in token")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				errorMessage(w, http.StatusUnauthorized, "error in token")
				return
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorMessage(w, http.StatusUnauthorized, "token is invalid")
			}
		} else {
			errorMessage(w, http.StatusUnauthorized, "token is invalid")
		}

	})
}

func errorMessage(w http.ResponseWriter, status int, message string) {
	var errorMessage models.Error
	errorMessage.Message = message
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorMessage)
}

func GenerateToken(user models.User) (string, error) {

	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Password,
		"iss":   "teonyx",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal("error in generating token string")
	}
	return tokenString, nil

}

func login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var jwt models.JWT

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	if user.Email == "" {
		errorMessage(w, http.StatusBadRequest, "invalid user")
		return
	}
	if user.Password == "" {
		errorMessage(w, http.StatusBadRequest, "empty password not accepted")
		return
	}
	password := user.Password
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err = row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			errorMessage(w, http.StatusNotFound, "user not found")
			return
		} else {
			log.Fatal("error occurred")
		}
	}
	hashPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		errorMessage(w, http.StatusUnauthorized, err.Error())
		return
	}
	token, err := GenerateToken(user)
	if err != nil {
		log.Fatal("error in token generation")
	}
	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	fmt.Println(token)
	_ := json.NewEncoder(w).Encode(jwt)

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
