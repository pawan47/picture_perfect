package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

//Dbtype :will contain all variables required for db
type Dbtype struct {
	db *sql.DB

}

var blacklist map[string]bool

// SecretKey used for jwt token verification and generation
var SecretKey string = "mykey"

// Credentials :for stroing user cred
type Credentials struct {
	UserID    int    `json:"userid"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Admin     bool   `json:"isadmin"`
}

// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	UserID int `json:"userid"`
	jwt.StandardClaims
}

// User will get user details from login page
type User struct {
	UserID   int    `json:"userid"`
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

// Dbhandler database struct global variable
var Dbhandler Dbtype

func main() {
	pqurl, err := pq.ParseURL("postgres://ccmnzryy:12OjOODSZeS_yTLUB-sDdJ3sU7swHDuz@arjuna.db.elephantsql.com:5432/ccmnzryy")
	if err != nil {
		fmt.Println(err)
	}
	blacklist = make(map[string]bool)
	// fmt.Println(pqurl)
	Dbhandler.db, err = sql.Open("postgres", pqurl)
	defer Dbhandler.db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = Dbhandler.db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", Homepage)
	r.HandleFunc("/signup", Signup).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/logout",Logout).Methods("POST")
	http.ListenAndServe(":8080", r)
}

// Homepage used for homepage routing
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the home page")
}


// Logout function will verify token then it will assign a new token which will exipre now only
func Logout(w http.ResponseWriter, r *http.Request) {
	var error Error 
	token := r.Header.Get("Auth")
	tokenArr := strings.Split(token, " ")
	if len(tokenArr) != 2{
		error.Message = "tokemissing"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(error)
		return
	}
	vaild, code, message := VerifyToken(tokenArr[1])
	if vaild == false{
		error.Message = message
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(error)
		return
	}

	// // w.Header["token"] = ""
	w.Header().Set("Auth", "")
	fmt.Println(tokenArr[1])
	blacklist[tokenArr[1]] = true
	fmt.Fprintf(w, "token erased and included in blacklist")

}


// VerifyToken will verify if token is valid or not
func VerifyToken(token string) (bool, int , string){
	// error string 
	// token := r.Header.Get("Auth")
	// tokenArr := strings.Split(token, " ")
	// if len(tokenArr) == 2{
		// token = tokenArr[1]
		_, pre := blacklist[token]
		if pre == true{
			return false, http.StatusUnauthorized, "invalid token"
		}
		AuthToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}

			return []byte(SecretKey), nil
		})
		if err != nil{
			return false, http.StatusUnauthorized, "invalid token"
		}
		if AuthToken.Valid{
			return true, http.StatusAccepted, ""
		}
		return false, http.StatusUnauthorized, "invalid token"


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
	token, err := Gentoken(user,false)
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

// Gentoken is used to generate token at login
func Gentoken(user User, TokenInvalid bool) (string, error) {
	expire := time.Now()
	if TokenInvalid == false{
		expire = expire.Add(time.Hour * 2)
	}
	claims := &Claims{
		Username: user.UserName,
		UserID: user.UserID,
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
