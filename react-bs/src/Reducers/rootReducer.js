import {combineReducers} from 'redux'

import movieReducer from './MovieReducer'

const rootReducer = combineReducers({
    movies: movieReducer
})

export default rootReducer;