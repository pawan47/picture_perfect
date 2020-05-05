package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Reset will reset password on the basis of username #NotProtected
func Reset(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		error.Message = "internal error"
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if user.UserName == "" {
		error.Message = "username empty"
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		error.Message = "password empty"
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cost := 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	user.Password = string(hashed)

	// stmt :=
	fmt.Println(user)
	stmt := "update user_details set password = $1 where username= $2 RETURNING user_id;"
	err = Dbhandler.db.QueryRow(stmt, user.Password, user.UserName).Scan(&user.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error.Message = "user doesn't exist"
		json.NewEncoder(w).Encode(error)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("user credentials updated")

}

// Logout function will verify token then it will assign a new token which will exipre now only
func Logout(w http.ResponseWriter, r *http.Request) {
	var error Error
	token := r.Header.Get("Authorization")
	// fmt.Println(token)
	tokenArr := strings.Split(token, " ")
	// fmt.Println(tokenArr,"s")
	if len(tokenArr) != 2 {
		error.Message = "tokemissing"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(error)
		return
	}
	vaild, code, message := VerifyToken(tokenArr[1])
	if vaild == false {
		error.Message = message
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(error)
		return
	}

	// // w.Header["token"] = ""
	w.Header().Set("Authorization", "")
	fmt.Println(tokenArr[1])
	blacklist[tokenArr[1]] = true
	fmt.Fprintf(w, "token erased and included in blacklist")

}

// Login endpoint for login check and generating jwt token
func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// fmt.Fprintf(w, err)
		error.Message = "internal error"
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if user.UserName == "" {
		error.Message = "username empty"
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		error.Message = "password empty"
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	password := user.Password
	stmt := "select user_id, password from user_details where username=$1;"
	row := Dbhandler.db.QueryRow(stmt, user.UserName)
	err = row.Scan(&user.UserID, &user.Password)
	fmt.Println(user)

	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusBadRequest)
		error.Message = "username not found"
		json.NewEncoder(w).Encode(error)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error.Message = "password incorrect"
		json.NewEncoder(w).Encode(error)
		return
	}
	token, err := Gentoken(user, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error.Message = "failed to generate key"
		json.NewEncoder(w).Encode(error)
	}
	var jwttoken JwtToken
	jwttoken.Token = token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwttoken)
	// cookie := http.Cookie{
	// 	Name:    "token",
	// 	Value:   token,
	// 	Expires: time.Now().AddDate(0, 0, 1),
	// }
	// http.SetCookie(w, &cookie)
	// return
}

// Signup uses POST method
func Signup(w http.ResponseWriter, r *http.Request) {

	var cred Credentials
	var error Error
	json.NewDecoder(r.Body).Decode(&cred)

	if cred.UserName == "" {
		error.Message = "username is missing"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	if cred.Password == "" {
		error.Message = "Empty Password"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	cost := 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(cred.Password), cost)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	cred.Password = string(hashed)

	stmt := "insert into user_details (first_name, last_name, username, password, is_admin) values($1, $2, $3, $4, $5) RETURNING user_id;"
	err = Dbhandler.db.QueryRow(stmt, cred.FirstName, cred.LastName, cred.UserName, cred.Password, false).Scan(&cred.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error.Message = "duplicate entry"
		json.NewEncoder(w).Encode(error)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("user signed up")
}
