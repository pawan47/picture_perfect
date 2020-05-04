package main

import (
	"database/sql"
	"fmt"
	"log"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/lib/pq"
)

// MoviesInfo will store movie details
type MoviesInfo struct {
	MovieID       int64   `json:"id"`
	Title         string  `json:"title"`
	Language      string  `json:"original_language"`
	ThumbnailLink string  `json:"poster_path"`
	Genre         string  `json:"genres"`
	Adult         bool    `json:"adult"`
	Overview      string  `json:"overview"`
	Tagline       string  `json:"tagline"`
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     int64   `json:"vote_count"`
}

func main() {
	var movies []int64
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	MovieArray := []MoviesInfo{}
	tmdbClient, err := tmdb.Init(("6b75cbe428b679bc311c9c50905fd313"))
	if err != nil {
		fmt.Println(err)
	}
	tmdbClient.SetClientAutoRetry()
	options := map[string]string{
		"language": "en-US",
		"page":     "3",
	}
	list, err := tmdbClient.GetDiscoverMovie(options)
	for i := 0; i < len(list.Results); i++ {
		movies = append(movies, (list.Results)[i].ID)
	}
	count := 0
	for i := 0; i < len(movies); i++ {

		det, err := tmdbClient.GetMovieDetails(int(movies[i]), nil)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mov := MoviesInfo{}
		mov.Adult = det.Adult
		mov.Genre = det.Genres[0].Name
		mov.Language = det.OriginalLanguage
		mov.MovieID = det.ID
		mov.Overview = det.Overview
		mov.Tagline = det.Tagline
		mov.ThumbnailLink = tmdb.GetImageURL(det.PosterPath, "original")
		mov.Title = det.Title
		mov.VoteAverage = det.VoteAverage
		mov.VoteCount = det.VoteCount

		stmt := "insert into movie_info (movie_id, movie_name, short_discription, long_discription, thumbnail_link, genre,language, vote_average, vote_count,actor,actress,director) values($1, $2, $3, $4, $5,$6,$7,$8,$9,$10,$11,$12);"
		// fmt.Println(mov.MovieID, mov.Title, mov.Tagline, mov.Overview, mov.ThumbnailLink, mov.Genre, mov.Language)
		err = db.QueryRow(stmt, mov.MovieID, mov.Title, mov.Tagline, mov.Overview, mov.ThumbnailLink, mov.Genre, mov.Language, mov.VoteAverage, mov.VoteCount, "", "", "").Scan()
		if err != nil {
			fmt.Println(err)
			continue
		}
		count++
		// MovieArray = append(MovieArray, mov)
	}
	defer db.Close()
	fmt.Println(count)
}

func connect() (*sql.DB, error) {
	pqurl, err := pq.ParseURL("postgres://ccmnzryy:12OjOODSZeS_yTLUB-sDdJ3sU7swHDuz@arjuna.db.elephantsql.com:5432/ccmnzryy")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(pqurl)
	db, err := sql.Open("postgres", pqurl)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}
	return db, nil
}
