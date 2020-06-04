import axios from 'axios'
import { FETCH_MOVIE } from './../Actions/types'

const fetchMovieDetails = (domain, movieId) => dispatch => {
    // console.log(domain + '/' + movieId)
    axios.get(domain + '/' + movieId)
        .then(res => {
            // console.log(res.data.review)
            dispatch(
                {
                    type: FETCH_MOVIE,
                    payload: {
                        movieId: res.data.movie_id,
                        movieName: res.data.movie_name,
                        poster: res.data.thumbnail_link,
                        rating: res.data.vote_average,
                        overview: res.data.long_discription,
                        genre: res.data.genre,
                        language: res.data.language,
                        popularity : res.data.vote_count,
                        review : res.data.review,
                    }
                }
            )
        }
        )
        .catch(err => console.log(err, "error at dispatch movie details"))
}


export default fetchMovieDetails;