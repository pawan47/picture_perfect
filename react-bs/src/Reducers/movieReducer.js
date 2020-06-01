
import {FETCH_MOVIES} from './../Actions/types'

const initState = {
    moviesList : [],
    domain : "http://13.232.168.173/movie/catalogue",
    params : {
        search : "",
        genre : "",
        sortBy : "",
    },
    moviesPerPage : 24,
    currentPageNo : 1,
    total_movies : 0,
    // loading : true,
};


const movieReducer = (state = initState, action) => {
    switch(action.type){
        case FETCH_MOVIES:
            return (
                {
                    ...state,
                    moviesList : action.payload,
                    total_movies : len(action.payload),
                }
            )
        default:
            return state
    }
}

export default movieReducer;