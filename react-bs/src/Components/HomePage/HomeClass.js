import React, { Component } from 'react';
import { connect } from 'react-redux'

//local imports
import List from './../CommonComponents/List'
import Paginate from './../CommonComponents/Paginate'
import getMovies from '../../Actions/fetchMovies'


class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      currentPage: 1,
      moviesPerPage: 24,
    }
  }

  componentDidMount() {
    const pp = {
      search : null,
      genre : null,
      sortby : "vote_average DESC"
    }
    this.props.getMovie(this.props.domain, pp)
  }

  changePage = (pageno) => {
    this.setState({
      currentPage: pageno,
    })
  }


  render() {

    const result = this.props.result.slice((this.state.currentPage - 1) * this.state.moviesPerPage, this.state.currentPage * this.state.moviesPerPage);
    console.log(result)
    return (
      <div style={{ backgroundColor: "#d2d2d275" }}>
        <List list={result} catalogueHidden={false} />
        <Paginate totalMovies={this.props.totalMovies} currentPage={this.state.currentPage} moviePerPage={this.state.moviesPerPage} changePage={this.changePage} />
      </div >
    );
  }
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
    getMovie : (state, type) => dispatch(getMovies(state, type)),
  };
};


export default connect(mapStateToProps, mapDisptachToProps)(Home);
