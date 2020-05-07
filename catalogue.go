package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	if err != nil {
		JSONErrorWriter(w, "No movie present by given id", http.StatusBadRequest)
		return
	}
	var t JwtToken
	t.Token = r.Header.Get("Authorization")
	valid, _ := t.CheckAuth()
	if valid == true {
		UserID, err := GetUserIdbyToken(strings.Split(t.Token, " ")[1])
		if err == nil {
			movie.UserRating = GetRatingbyID(UserID, movie.MovieID)
			movie.UserReview = GetReviewbyID(UserID, movie.MovieID)
		}

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

// GetMoviesFilter will handle get /movies{name} request
func GetMoviesFilter(w http.ResponseWriter, r *http.Request) {
	stmt := "SELECT movie_id, movie_name, short_discription, thumbnail_link, vote_average, genre, language FROM movie_info"
	var conditions []string
	res := r.URL.Query()["genre"]
	for _, ele := range res {
		conditions = append(conditions, fmt.Sprintf("genre = '%s'", ele))
	}
	ress, err := GetQuery(r, "language")
	if err == nil {
		conditions = append(conditions, fmt.Sprintf("language = '%s'", ress))
	}
	if len(conditions) > 0 {
		stmt += " WHERE " + strings.Join(conditions[:], " AND ")
	}
	ress, _ = GetQuery(r, "sortby")
	lim, _ := GetQuery(r, "limit")
	off, _ := GetQuery(r, "offset")
	if ress == "" || lim == "" || off == "" {
		JSONErrorWriter(w, "invalid URL", http.StatusBadRequest)
		return
	}
	limit, _ := strconv.Atoi(lim)
	offset, _ := strconv.Atoi(off)
	if offset < 1 {
		JSONErrorWriter(w, "invalid URL", http.StatusBadRequest)
		return
	}
	offset = (offset - 1) * limit

	stmt += fmt.Sprintf(" ORDER BY %s", ress)
	stmt += fmt.Sprintf(" LIMIT %d", limit)
	stmt += fmt.Sprintf(" OFFSET %d", offset)

	row, err := Dbhandler.db.Query(stmt)
	defer row.Close()
	if err != nil && err != sql.ErrNoRows {
		JSONErrorWriter(w, "internal database error", http.StatusInternalServerError)
		return
	}
	var result []MovieListInfo

	for row.Next() {
		var info MovieListInfo
		if err := row.Scan(&info.MovieID, &info.Title, &info.Tagline, &info.ThumbnailLink, &info.VoteAverage, &info.Genre, &info.Language); err != nil {
			continue
		}
		result = append(result, info)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

// func GetMoviesSearch()
