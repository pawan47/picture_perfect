package main

import (
	"database/sql"
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
	r.HandleFunc("/movie/catalogue/{ID}", GetMoviesByID).Methods("GET")
	r.HandleFunc("/movies/rating/{ID}", RateMovie).Methods("POST")
	r.HandleFunc("/movies/review/{ID}", ReviewMovie).Methods("POST")
	r.HandleFunc("/movies/rating/{ID}", DelRateMovie).Methods("DELETE")
	r.HandleFunc("/movies/review/{ID}", DelReviewMovie).Methods("DELETE")
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

// GetUserIdbyToken will give userid given valid token
func GetUserIdbyToken(token string) (int, error) {
	TokenNew, err := jwt.Parse(token, nil)
	if TokenNew == nil {
		return 0, err
	}
	return int(((TokenNew.Claims.(jwt.MapClaims))["userid"]).(float64)), nil
}

// GetRatingbyID will return movie rating given movie id and user id
func GetRatingbyID(ID, MovieID int) int {
	rating := 0
	stmt := "select rating from rating_review where user_id = $1 AND movie_id = $2"
	err := Dbhandler.db.QueryRow(stmt, ID, MovieID).Scan(&rating)
	if err != nil {
		return 0
	}

	return rating
}

// GetReviewbyID will retern review given by user if not found will return empty string
func GetReviewbyID(ID, MovieID int) string {
	review := ""
	stmt := "select review from rating_review where user_id = $1 AND movie_id = $2"
	err := Dbhandler.db.QueryRow(stmt, ID, MovieID).Scan(&review)
	// fmt.Println("aaab")
	if err != nil {
		return ""
	}
	return review
}

// RowExists will
func RowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := Dbhandler.db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false
	}
	return exists
}
