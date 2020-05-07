package main

import "github.com/dgrijalva/jwt-go"

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
	jwt.StandardClaims
}

// User will get user details from login page
type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

//Error will have all variable required to handle errors
type Error struct {
	Message string `json:"error"`
}

// JwtToken huhuhuhu njnn
type JwtToken struct {
	Token string `json:"jwttoken"`
}

// MoviesInfo will be used to hold various movies details from db
type MoviesInfo struct {
	MovieID       int     `json:"movie_id"`
	Title         string  `json:"movie_name"`
	Language      string  `json:"language"`
	ThumbnailLink string  `json:"thumbnail_link"`
	Genre         string  `json:"genre"`
	Overview      string  `json:"long_discription"`
	Tagline       string  `json:"short_discription"`
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     int64   `json:"vote_count"`
	Actor         *string `json:"actor"`
	Actress       *string `json:"actress"`
	Director      *string `json:"director"`
	UserRating    int     `json:"user_rating"`
	UserReview    string  `json:"user_review"`
}

type MovieListInfo struct {
	MovieID       int     `json:"movie_id"`
	Title         string  `json:"movie_name"`
	Language      string  `json:"language"`
	ThumbnailLink string  `json:"thumbnail_link"`
	Genre         string  `json:"genre"`
	Tagline       string  `json:"short_discription"`
	VoteAverage   float32 `json:"vote_average"`
}

// gorm struct
