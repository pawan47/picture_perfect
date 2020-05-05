package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// EndpointsInit will be used for router
func EndpointsInit() {
	r := mux.NewRouter()
	r.HandleFunc("/", Homepage)
	r.HandleFunc("/signup", Signup).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")
	r.HandleFunc("/reset", Reset).Methods("POST")
	// r.HandleFunc("/movie/catalogue/{}", GetMoviesList).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Gentoken is used to generate token at login
func Gentoken(user User, TokenInvalid bool) (string, error) {
	expire := time.Now()
	if TokenInvalid == false {
		expire = expire.Add(time.Hour * 2)
	}
	claims := &Claims{
		Username: user.UserName,
		UserID:   user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenstring, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenstring, nil

}

// VerifyToken will verify if token is valid or not
func VerifyToken(token string) (bool, int, string) {
	// error string
	// token := r.Header.Get("Auth")
	// tokenArr := strings.Split(token, " ")
	// if len(tokenArr) == 2{
	// token = tokenArr[1]
	_, pre := blacklist[token]
	if pre == true {
		return false, http.StatusUnauthorized, "invalid token"
	}
	AuthToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return []byte(SecretKey), nil
	})
	if err != nil {
		return false, http.StatusUnauthorized, "invalid token"
	}
	if AuthToken.Valid {
		return true, http.StatusAccepted, ""
	}
	return false, http.StatusUnauthorized, "invalid token"

}
