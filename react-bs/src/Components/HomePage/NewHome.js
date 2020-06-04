import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux'

//local imports
import List from './../CommonComponents/List'
import Paginate from './../CommonComponents/Paginate'
import getMovies from '../../Actions/fetchMovies'


const Home = (props) => {

    const [currentPage, setCurrentPage] = useState(1);
    // const [moviesPerPage, setMoviesPerPage] = useState(24);
    const [ignore, setIgnore] = useState(false)
    const moviesPerPage = 24;


    // have to destructre because of warning and it is suggested to destructure ---- github issue facebook/react
    const {domain, getMovie} = props;
    useEffect(() => {
        if (ignore === false) {
            const pp = {
                search: null,
                genre: null,
                sortby: "vote_average DESC"
            }
            getMovie(domain, pp)
            setIgnore(true);
        }
    }, [domain, getMovie,ignore, setIgnore])

    const result = props.result.slice((currentPage - 1) * moviesPerPage, currentPage * moviesPerPage);
    // console.log(result)
    return (
        <div style={{ backgroundColor: "rgb(208, 208, 208)" }}>
            <List list={result} catalogueHidden={false} />
            <Paginate totalMovies={props.totalMovies} currentPage={currentPage} moviePerPage={moviesPerPage} changePage={setCurrentPage} />
        </div >
    );

}



const mapStateToProps = state => {
    return (
        {
            result: state.movies.moviesList,
            totalMovies: state.movies.total_movies,
            domain: state.movies.domain,
        }
    )
}


const mapDisptachToProps = (dispatch) => {
    return {
        getMovie: (state, type) => dispatch(getMovies(state, type)),
    };
};


export default connect(mapStateToProps, mapDisptachToProps)(Home);
