import {ADD_REVIEW} from './types'

const addReview = (name,review) => dispatch => {
    console.log("add review called")
    dispatch({
        type : ADD_REVIEW,
        payload : {
            name : name,
            body : review,
        }
    })
}

export default addReview;