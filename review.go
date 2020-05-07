package main

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ReviewMovie required params movie_id, api_key, review
func ReviewMovie(w http.ResponseWriter, r *http.Request) {
	review := r.URL.Query().Get("review")
	if review == "" {
		JSONErrorWriter(w, "Review is missing", http.StatusBadRequest)
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
	if err != nil {
		JSONErrorWriter(w, "Token Invalid", http.StatusBadRequest)
		return
	}
	stmt := "select * from rating_review where movie_id = $1 AND user_id = $2"
	if RowExists(stmt, ID, UserID) {
		stmt = "update rating_review set review = $1 where movie_id = $2 AND user_id = $3"
		err = Dbhandler.db.QueryRow(stmt, review, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			JSONErrorWriter(w, "Updation Failed", http.StatusInternalServerError)
			return
		}
	} else {
		stmt = "insert into rating_review (movie_id, user_id, review) values($1,$2,$3)"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID, review).Scan()
		if err != nil && err != sql.ErrNoRows {
			JSONErrorWriter(w, "Movie id is not present", http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusAccepted)
}

// DelReviewMovie will delete review of the authenticated user
func DelReviewMovie(w http.ResponseWriter, r *http.Request) {
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
		JSONErrorWriter(w, "Token Invalid", http.StatusBadRequest)
		return
	}
	stmt := "select review from rating_review where movie_id = $1 AND user_id = $2 AND review IS NOT NULL"
	if RowExists(stmt, ID, UserID) {
		stmt = "UPDATE rating_review SET review = NULL where movie_id = $1 AND user_id = $2"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			JSONErrorWriter(w, "Deletion Failed", http.StatusInternalServerError)
			return
		}
	} else {
		JSONErrorWriter(w, "Review does not exist", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
