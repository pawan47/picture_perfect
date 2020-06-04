import {combineReducers} from 'redux'

import movieReducer from './MovieReducer'
import movieDetailReducers from './movieDetailReducers'

const rootReducer = combineReducers({
    movies: movieReducer,
    movieDetail : movieDetailReducers,
})

export default rootReducer;