
import { FETCH_MOVIES, UPDATE_PARAMS } from './../Actions/types'

const initState = {
    moviesList: [],
    domain: "http://localhost:8080/movie/catalogue",
    params: {
        search: null,
        genre: null,
        sortby: "vote_average DESC",
    },
    moviesPerPage: 24,
    currentPageNo: 1,
    total_movies: 0,
    loading : true,
};


const movieReducer = (state = initState, action) => {
    switch (action.type) {
        case FETCH_MOVIES:
            return (
                {
                    ...state,
                    moviesList: action.payload,
                    total_movies: action.payload.length,
                    params : action.params
                }
            )
        case UPDATE_PARAMS:
            return (
                {
                    ...state,
                    params: action.params
                }
            )
        default:
            return state
    }
}

export default movieReducer;