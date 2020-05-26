package main

import (
	"database/sql"
	"encoding/json"
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
	r.HandleFunc("/hulu", GetMoviesFilter).Methods("GET")
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
		Admin:    false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenstring, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenstring, nil

}

// VerifyToken will verify if token is valid or not
func VerifyToken(token string) (bool, string) {
	_, pre := blacklist[token]
	if pre == true {
		return false, "invalid token"
	}

	AuthToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return []byte(SecretKey), nil
	})
	if err != nil {
		return false, "invalid token"
	}
	if AuthToken.Valid {
		return true, ""
	}
	return false, "invalid token"

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

// JSONErrorWriter error writer
func JSONErrorWriter(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	var error Error
	error.Message = err
	json.NewEncoder(w).Encode(error)
}

// GetQuery will fetch query data and catches empty query
func GetQuery(r *http.Request, param string) (string, error) {
	params := r.URL.Query().Get(param)
	var err Error
	// fmt.Println(params,)
	if params == "" {
		err.Message = fmt.Sprintf("%s query not found", param)
		return "", &err
	}
	return params, nil
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
