package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// RateMovie required params movie_id, api_key, rating
func RateMovie(w http.ResponseWriter, r *http.Request) {
	// check authentication

	var error Error
	rating := r.URL.Query().Get("rate")
	if rating == "" {
		error.Message = "Rating missing"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	rate, err := strconv.Atoi(rating)
	if rate < 0 || rate > 10 {
		error.Message = "Invalid Rating"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	if err != nil {
		error.Message = "Rating should be integer"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	token := r.Header.Get("Authorization")
	tokenArr := strings.Split(token, " ")
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

	vars := mux.Vars(r)
	ID := vars["ID"]
	UserID, err := GetUserIdbyToken(tokenArr[1])
	if err != nil {
		error.Message = "token invalid"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	stmt := "select * from rating_review where movie_id = $1 AND user_id = $2"
	if RowExists(stmt, ID, UserID) {
		stmt = "update rating_review set rating = $1 where movie_id = $2 AND user_id = $3"
		err = Dbhandler.db.QueryRow(stmt, rate, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			error.Message = "updation failed"
			json.NewEncoder(w).Encode(error)
			// fmt.Println(err)
			return
		}
	} else {
		stmt = "insert into rating_review (movie_id, user_id, rating) values($1,$2,$3)"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID, rate).Scan()
		if err != nil && err != sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			error.Message = "Movie id is not present"
			json.NewEncoder(w).Encode(error)
			// fmt.Println(err)
			return
		}
	}
	w.WriteHeader(http.StatusAccepted)
}

// DelRateMovie will delete rating by user
func DelRateMovie(w http.ResponseWriter, r *http.Request) {
	var error Error
	token := r.Header.Get("Authorization")
	tokenArr := strings.Split(token, " ")
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

	vars := mux.Vars(r)
	ID := vars["ID"]
	UserID, err := GetUserIdbyToken(tokenArr[1])
	if err != nil {
		error.Message = "token invalid"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	stmt := "select rating from rating_review where movie_id = $1 AND user_id = $2 AND rating IS NOT NULL"
	if RowExists(stmt, ID, UserID) {
		stmt = "UPDATE table rating_review SET rating = NULL where movie_id = $1 AND user_id = $2"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			error.Message = "deletion failed"
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
		}
	} else {
		error.Message = "rating DNE"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}
