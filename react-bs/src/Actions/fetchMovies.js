import { FETCH_MOVIES } from './types'
import axios from 'axios'

export const getMovies = (domain, params) => dispatch => {
    axios
    .get(domain, {
        params: params
    })
    .then(res => {
            if (res.data !== null) {
                dispatch({
                    type : FETCH_MOVIES,
                    payload : res.data,
                    params : params
                })
                
            } else {
                dispatch({
                    type: FETCH_MOVIES,
                    payload: [],
                    params : params
                })
            }
        }
        )
    .catch(error => {
        console.log("error while fetching movies", error)
    })
}

export default getMovies;