package main

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Credentials :for stroing user cred
type Credentials struct {
	UserID    int    `json:"userid"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Admin     bool   `json:"isadmin"`
}

// Claims will have header deatils along with jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"userid"`
	Admin    bool   `json:"is_admin"`
	jwt.StandardClaims
}

// User will get user details from login page
type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// ReviewEndPoint will fetch review from an external api for now as we dont have much review
type ReviewEndPoint struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

//Error will have all variable required to handle errors
type Error struct {
	Message string `json:"error"`
}

func (e *Error) Error() string {
	return e.Message
}

// JwtToken huhuhuhu njnn
type JwtToken struct {
	Token string `json:"jwttoken"`
}

// CheckAuth Method is attached to JWTtoken struct which will check token validity
func (t *JwtToken) CheckAuth() (bool, string) {
	tokenArr := strings.Split(t.Token, " ")
	if len(tokenArr) == 2 {
		valid, message := VerifyToken(tokenArr[1])
		if valid == true {
			return true, ""
		}
		return false, message
	}
	return false, "Token missing"
}

// GetUserID method is attached to jwttoken
func (t *JwtToken) GetUserID() (int, error) {
	TokenNew, err := jwt.Parse(strings.Split(t.Token, " ")[1], nil)
	if TokenNew == nil {
		return 0, err
	}
	return int(((TokenNew.Claims.(jwt.MapClaims))["userid"]).(float64)), nil

}

// IsAdmin method is attached to jwttoken
func (t *JwtToken) IsAdmin() (bool, error) {
	TokenNew, err := jwt.Parse(strings.Split(t.Token, " ")[1], nil)
	if TokenNew == nil {
		return false, err
	}
	return ((TokenNew.Claims.(jwt.MapClaims))["userid"]).(bool), nil

}

// func (t *JwtToken)

// MoviesInfo will be used to hold various movies details from db
type MoviesInfo struct {
	MovieID       int              `json:"movie_id"`
	Title         string           `json:"movie_name"`
	Language      string           `json:"language"`
	ThumbnailLink string           `json:"thumbnail_link"`
	Genre         string           `json:"genre"`
	Overview      string           `json:"long_discription"`
	Tagline       string           `json:"short_discription"`
	VoteAverage   float32          `json:"vote_average"`
	VoteCount     int64            `json:"vote_count"`
	Actor         *string          `json:"actor"`
	Actress       *string          `json:"actress"`
	Director      *string          `json:"director"`
	UserRating    int              `json:"user_rating"`
	UserReview    string           `json:"user_review"`
	Review        []ReviewEndPoint `json:"review"`
}

// MovieListInfo will hold get /movie/{name} request
type MovieListInfo struct {
	MovieID       int     `json:"movie_id"`
	Title         string  `json:"movie_name"`
	Language      string  `json:"language"`
	ThumbnailLink string  `json:"thumbnail_link"`
	Genre         string  `json:"genre"`
	Tagline       string  `json:"short_discription"`
	VoteAverage   float32 `json:"vote_average"`
}
