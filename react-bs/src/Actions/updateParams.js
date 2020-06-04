import {UPDATE_PARAMS} from './types'

const updateParam = (search, genre, sortby) => dispatch => {
    return (
        dispatch({
            type : UPDATE_PARAMS,
            params : {
                search : search,
                genre : genre,
                sortby : sortby,
            }
        })
    )
}

export default updateParam;