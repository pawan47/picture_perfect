package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ReviewMovie required params movie_id, api_key, review
func ReviewMovie(w http.ResponseWriter, r *http.Request) {
	// check authentication

	var error Error
	review := r.URL.Query().Get("review")
	if review == "" {
		error.Message = "Review is missing"
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
		stmt = "update rating_review set review = $1 where movie_id = $2 AND user_id = $3"
		err = Dbhandler.db.QueryRow(stmt, review, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			error.Message = "updation failed"
			json.NewEncoder(w).Encode(error)
			// fmt.Println(err)
			return
		}

	} else {
		stmt = "insert into rating_review (movie_id, user_id, review) values($1,$2,$3)"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID, review).Scan()
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

// DelReviewMovie will delete review of the authenticated user
func DelReviewMovie(w http.ResponseWriter, r *http.Request) {
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
	stmt := "select review from rating_review where movie_id = $1 AND user_id = $2 AND review IS NOT NULL"
	if RowExists(stmt, ID, UserID) {
		stmt = "UPDATE rating_review SET review = NULL where movie_id = $1 AND user_id = $2"
		err = Dbhandler.db.QueryRow(stmt, ID, UserID).Scan()
		if err != nil && err != sql.ErrNoRows {
			error.Message = "deletion failed"
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
	} else {
		error.Message = "review DNE"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}
