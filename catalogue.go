package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

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