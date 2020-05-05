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

// gorm struct
