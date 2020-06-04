import React, { Component } from 'react';
import { connect } from 'react-redux'

//local imports
import List from './../CommonComponents/List'
import Paginate from './../CommonComponents/Paginate'
import SimpleSelect from './Filter'
import {GetInitParams, convertParams} from './GetInitParams'

// ActionCreaters
import getMovies from './../../Actions/fetchMovies'
import updateParams from './../../Actions/updateParams'

class SearchPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            currentPage: 1,
            moviesPerPage: 24,
            search: GetInitParams("search", this.props.location.search),
            genre: GetInitParams("genre", this.props.location.search),
            sortBy: "vote_average DESC",
        }
    }

    componentDidMount() {
        this.props.getMovie(this.props.domain, convertParams(this.state.search, this.state.genre, this.state.sortBy))
    }


    // newUpdateParam = (key, param) => {
    //     this.setState({
    //         [key]: param,
    //     },
    //         () => this.getMovies()
    //     )
    // }

    updateParams = (search, genre, sortBy) => {
        this.props.getMovie(this.props.domain, convertParams(search, genre, sortBy))
    }

    changePage = (pageno) => {
        this.setState({
            currentPage: pageno,
        })
    }

    render() {
        const result = this.props.result.slice((this.state.currentPage - 1) * this.state.moviesPerPage, this.state.currentPage * this.state.moviesPerPage);
        const initparams = {
            search: this.state.search,
            genre: this.state.genre,
        }
        return (
            <div style={{ backgroundColor: "#d2d2d275" }}>
                <SimpleSelect updateParams={this.updateParams} initparams={initparams} />
                <List list={result} catalogueHidden={true} />
                <Paginate totalMovies={this.props.totalMovies} currentPage={this.state.currentPage} moviePerPage={this.state.moviesPerPage} changePage={this.changePage} />
            </div >
        );
    }
}


const mapStatesToProps = state => {
    return (
        {
            result: state.movies.moviesList,
            totalMovies: state.movies.total_movies,
            domain: state.movies.domain,
            params : state.movies.params,
        }
    )
}


const mapDisptachToProps = (dispatch) => {
    return {
      getMovie : (state, type) => dispatch(getMovies(state, type)),
      updateParam : (search, genre, sortby) => dispatch(updateParams(search, genre, sortby)),

    };
  };

export default connect(mapStatesToProps,mapDisptachToProps)(SearchPage);
