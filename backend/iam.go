package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Reset will reset password on the basis of username #NotProtected
func Reset(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JSONErrorWriter(w, "Internal Error", http.StatusNoContent)
		return
	}
	if user.UserName == "" {
		JSONErrorWriter(w, "User name empty", http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		JSONErrorWriter(w, "passord Empty", http.StatusBadRequest)
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
		JSONErrorWriter(w, "User Doesn't Exist", http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("user credentials updated")

}

// Logout function will verify token then it will assign a new token which will exipre now only
func Logout(w http.ResponseWriter, r *http.Request) {
	var t JwtToken
	t.Token = r.Header.Get("Authorization")
	valid, message := t.CheckAuth()
	if valid == false {
		JSONErrorWriter(w, message, http.StatusUnauthorized)
		return
	}
	w.Header().Set("Authorization", "")
	blacklist[t.Token] = true
	fmt.Fprintf(w, "token erased and included in blacklist")

}

// Login endpoint for login check and generating jwt token
func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JSONErrorWriter(w, "Internal Error", http.StatusNoContent)
		return
	}

	if user.UserName == "" {
		JSONErrorWriter(w, "UserName empty", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		JSONErrorWriter(w, "Password Empty", http.StatusBadRequest)
		return
	}
	password := user.Password
	stmt := "select user_id, password from user_details where username=$1;"
	row := Dbhandler.db.QueryRow(stmt, user.UserName)
	err = row.Scan(&user.UserID, &user.Password)
	fmt.Println(user)

	if err != nil {
		JSONErrorWriter(w, "User Name not found in database", http.StatusBadRequest)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		JSONErrorWriter(w, "Password Incorrect", http.StatusBadRequest)
		return
	}
	token, err := Gentoken(user, false)
	if err != nil {
		JSONErrorWriter(w, "Failed to generate key", http.StatusInternalServerError)
		return
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
	json.NewDecoder(r.Body).Decode(&cred)

	if cred.UserName == "" {
		JSONErrorWriter(w, "User Name is missing", http.StatusBadRequest)
		return
	}
	if cred.Password == "" {
		JSONErrorWriter(w, "Empty Password", http.StatusBadRequest)
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
		JSONErrorWriter(w, "Duplicate Entry", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("user signed up")
}
