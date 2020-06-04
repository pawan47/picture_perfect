package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetMoviesByID is a endpoint for the GET /movies/{name} request
func GetMoviesByID(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)
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

	// As we dont have review I am fetching review from a different end point

	response, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	defer response.Body.Close()
	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		if err == nil {
			// movie.Review{}
			// stringdata := string(data)
			// fmt.Println(stringdata)
			// var reviewjson []ReviewEndPoint
			json.Unmarshal(data, &movie.Review)
			movie.Review = movie.Review[:10]
			// fmt.Println(movie.Review[:10])
		}
		// fmt.Println(ReviewMovie)
		// fmt.Println(string(data))
		// movie.Review = data

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

// GetMoviesFilter will handle get /movies{name} request
func GetMoviesFilter(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w) //for cors
	// fmt.Println()
	stmt := "SELECT movie_id, movie_name, short_discription, thumbnail_link, vote_average, genre, language FROM movie_info"
	var conditions []string
	genre, err := GetQuery(r, "genre")
	if err == nil {
		conditions = append(conditions, fmt.Sprintf("genre = '%s'", genre))
	}
	search, err := GetQuery(r, "search")

	if err == nil {
		s := "%"
		for i := 0; i < len(search); i++ {
			s += string(search[i]) + "%"
		}
		conditions = append(conditions, fmt.Sprintf("movie_name LIKE '%s'", s))
	}

	lan, err := GetQuery(r, "language")
	if err == nil {
		conditions = append(conditions, fmt.Sprintf("language = '%s'", lan))
	}
	if len(conditions) > 0 {
		stmt += " WHERE " + strings.Join(conditions[:], " AND ")
	}
	sortby, _ := GetQuery(r, "sortby")

	if sortby == "" {
		JSONErrorWriter(w, "invalid URL", http.StatusBadRequest)
		return
	}
	stmt += fmt.Sprintf(" ORDER BY %s", sortby)
	// fmt.Println(stmt)
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
