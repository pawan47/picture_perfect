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
	var error Error
	if err != nil {
		error.Message = "no movie present by given id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		fmt.Println(err)
		return
	}

	token := r.Header.Get("Authorization")
	tokenArr := strings.Split(token, " ")
	if len(tokenArr) == 2 {
		vaild, _, _ := VerifyToken(tokenArr[1])

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

func GetMoviesFilter(w http.ResponseWriter, r *http.Request) {
	stmt := "SELECT movie_id, movie_name, short_discription, thumbnail_link, vote_average, genre, language FROM movie_info"
	var conditions []string
	var error Error
	res := r.URL.Query()["genre"]
	for _, ele := range res {
		conditions = append(conditions, fmt.Sprintf("genre = '%s'", ele))
	}
	ress := r.URL.Query().Get("language")
	if ress != "" {
		conditions = append(conditions, fmt.Sprintf("language = '%s'", ress))
	}
	if len(conditions) > 0 {
		stmt += " WHERE " + strings.Join(conditions[:], " AND ")
	}
	ress = r.URL.Query().Get("sortby")
	lim := r.URL.Query().Get("limit")
	off := r.URL.Query().Get("offset")
	if ress == "" || lim == "" || off == "" {
		error.Message = "invalid URL"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	limit, _ := strconv.Atoi(lim)
	offset, _ := strconv.Atoi(off)
	if offset < 1 {
		error.Message = "invalid URL"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	offset = (offset - 1) * limit

	stmt += fmt.Sprintf(" ORDER BY %s", ress)
	stmt += fmt.Sprintf(" LIMIT %d", limit)
	stmt += fmt.Sprintf(" OFFSET %d", offset)

	row, err := Dbhandler.db.Query(stmt)
	defer row.Close()
	if err != nil && err != sql.ErrNoRows {
		error.Message = "internal database error"

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	var result []MovieListInfo

	for row.Next() {
		var info MovieListInfo
		if err := row.Scan(&info.MovieID, &info.Title, &info.Tagline, &info.ThumbnailLink, &info.VoteAverage, &info.Genre, &info.Language); err != nil {
			panic(err)
		}
		result = append(result, info)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

// func GetMoviesSearch()
