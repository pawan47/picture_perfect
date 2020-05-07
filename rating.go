package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// RateMovie required params movie_id, api_key, rating
func RateMovie(w http.ResponseWriter, r *http.Request) {
	rating := r.URL.Query().Get("rate")
	if rating == "" {
		JSONErrorWriter(w, "Rating missing", http.StatusBadRequest)
		return
	}
	rate, err := strconv.Atoi(rating)
	if rate < 0 || rate > 10 {
		JSONErrorWriter(w, "Invalid Rating", http.StatusBadRequest)
		return
	}
	if err != nil {
		JSONErrorWriter(w, "Rating should be integer", http.StatusBadRequest)
		return
	}

	var t JwtToken
	t.Token = r.Header.Get("Authorization")
	valid, message := t.CheckAuth()
	if valid == false {
		JSONErrorWriter(w, message, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	ID := vars["ID"]
	UserID, err := GetUserIdbyToken(strings.Split(t.Token, " ")[1])
	// fmt.Println(err)
	if err != nil {
		JSONErrorWriter(w, "Token Invalid", http.StatusBadRequest)
		return
	}
	stmt := "select * from rating_review where movie_id = $1 AND user_id = $2"
	if RowExists(stmt, ID, UserID) {
		stmt = "update rating_review set rating = $1 where movie_id = $2 AND user_id = $3"
		err = Dbhandler.db.QueryRow(stmt, rate, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			JSONErrorWriter(w, "updation failed", http.StatusInternalServerError)
			return
		}
	} else {
		stmt = "insert into rating_review (movie_id, user_id, rating) values($1,$2,$3)"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID, rate).Scan()
		if err != nil && err != sql.ErrNoRows {
			JSONErrorWriter(w, "Movie id is not present", http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusAccepted)
}

// DelRateMovie will delete rating by user
func DelRateMovie(w http.ResponseWriter, r *http.Request) {
	var t JwtToken
	t.Token = r.Header.Get("Authorization")
	valid, message := t.CheckAuth()
	if valid == false {
		JSONErrorWriter(w, message, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	ID := vars["ID"]
	UserID, err := GetUserIdbyToken(strings.Split(t.Token, " ")[1])
	if err != nil {
		JSONErrorWriter(w, "token Invalid", http.StatusBadRequest)
		return
	}
	stmt := "select rating from rating_review where movie_id = $1 AND user_id = $2 AND rating IS NOT NULL"
	if RowExists(stmt, ID, UserID) {
		stmt = "UPDATE rating_review SET rating = NULL where movie_id = $1 AND user_id = $2"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			JSONErrorWriter(w, "deletion failed", http.StatusInternalServerError)
			return
		}
	} else {
		JSONErrorWriter(w, "rating does not exist", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
