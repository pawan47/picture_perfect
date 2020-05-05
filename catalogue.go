package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetMoviesByID is a endpoint for the GET /movies/{name} request
func GetMoviesByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	var movie MoviesInfo
	stmt := "select * from movie_info where movie_id = $1;"
	row := Dbhandler.db.QueryRow(stmt, ID)
	err := row.Scan(&movie.MovieID, &movie.Title, &movie.Tagline, &movie.Overview, &movie.ThumbnailLink, &movie.Genre, &movie.Language, &movie.Actor, &movie.Actress, &movie.Director, &movie.VoteAverage, &movie.VoteCount)
	var error Error
	if err != nil {
		error.Message = "no movie present by given id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		fmt.Println(err)
		return
	}

	token := r.Header.Get("Authorization")
	// fmt.Println(token)
	tokenArr := strings.Split(token, " ")
	// fmt.Println(tokenArr,"s")
	if len(tokenArr) == 2 {
		vaild, _, _ := VerifyToken(tokenArr[1])
		// fmt.Println(vaild)

		if vaild == true {
			UserID, err := GetUserIdbyToken(tokenArr[1])
			// fmt.Println(UserID, err)
			if err == nil {
				// fmt.Println("get review")
				movie.UserRating = GetRatingbyID(UserID, movie.MovieID)
				movie.UserReview = GetReviewbyID(UserID, movie.MovieID)
			}

		}

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}
