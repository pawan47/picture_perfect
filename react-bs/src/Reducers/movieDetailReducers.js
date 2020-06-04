import {FETCH_MOVIE, ADD_REVIEW} from './../Actions/types'


const initState = {
    movieId : null,
    movieName : null,
    poster : null,
    rating : null,
    review : null,
    overview : null,
    genre : null,
    language : null,
    popularity : null,
}


const movieDetailReducers = (state = initState, action) => {
    switch(action.type){
        case FETCH_MOVIE:
            return({
                ...state,
                movieId : action.payload.movieId,
                movieName : action.payload.movieName,
                poster : action.payload.poster,
                rating : action.payload.rating,
                overview : action.payload.overview,
                genre : action.payload.genre,
                language : action.payload.language,
                popularity : action.payload.popularity,
                review : action.payload.review,
            })
        case ADD_REVIEW:
            return({
                ...state,
                review : [action.payload, ...state.review]
            })
        default:
            return state
    }
}

export default movieDetailReducers;