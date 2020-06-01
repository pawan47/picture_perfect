import {createStore, applyMiddleware, compose} from 'redux'
//createStore(reducers, initstate, enhancers)
import thunk from 'redux-thunk'
import rootReducer from './Reducers/rootReducer'

const initState = {}

const middleware = [thunk]

const store = createStore(rootReducer,
    initState,
    compose(
        applyMiddleware(middleware)
    )
    )

export default store;