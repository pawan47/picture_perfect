import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux'

//local imports
import List from './../CommonComponents/List'
import Paginate from './../CommonComponents/Paginate'
import SimpleSelect from './Filter'
import { GetInitParams, convertParams } from './GetInitParams'

// ActionCreaters
import getMovies from './../../Actions/fetchMovies'
import updateParams from './../../Actions/updateParams'


const SearchPage = (props) => {
    const [currentPage, setcurrentPage] = useState(1);
    const moviesPerPage = 24;

    const [params, setparams] = useState(
        {
            search: GetInitParams("search", props.location.search),
            genre: GetInitParams("genre", props.location.search),
            sortby: "vote_average DESC",
        }
    )
    const {getMovie, domain} = props
    useEffect(() => {
        getMovie(domain, convertParams(params.search, params.genre, params.sortby))
    },[getMovie, domain, params])


    const result = props.result.slice((currentPage - 1) * moviesPerPage, currentPage * moviesPerPage);
    const initparams = {
        search: params.search,
        genre: params.genre,
    }
    return (
        <div style={{ backgroundColor: "#d2d2d275" }}>
            <SimpleSelect setparams={setparams} initparams={initparams} />
            <List list={result} catalogueHidden={true} />
            <Paginate totalMovies={props.totalMovies} currentPage={currentPage} moviePerPage={moviesPerPage} changePage={setcurrentPage} />
        </div >
    );

}


const mapStatesToProps = state => {
    return (
        {
            result: state.movies.moviesList,
            totalMovies: state.movies.total_movies,
            domain: state.movies.domain,
            params: state.movies.params,
        }
    )
}


const mapDisptachToProps = (dispatch) => {
    return {
        getMovie: (state, type) => dispatch(getMovies(state, type)),
        updateParam: (search, genre, sortby) => dispatch(updateParams(search, genre, sortby)),

    };
};

export default connect(mapStatesToProps, mapDisptachToProps)(SearchPage);
