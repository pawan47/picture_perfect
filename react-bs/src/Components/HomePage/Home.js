import React, { Component } from 'react';
import axios from 'axios'

//local imports
import List from './../CommonComponents/List'
import Paginate from './../CommonComponents/Paginate'



class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      currentPage: 1,
      moviesPerPage: 24,
      result: [],
      totalMovies : 0,
      domain : "http://13.232.168.173"
    }
  }

  componentDidMount() {
    axios.get(this.state.domain + '/movie/catalogue?sortby=vote_average DESC')
      .then(res => {
        // console.log(res.data)
        this.setState({
          result : res.data,
          totalMovies : res.data.length,
        }
        )
      }
      )
  }

  changePage = (pageno) => {
    this.setState({
      currentPage : pageno,
    })
  }

  
  render() {
    const result = this.state.result.slice(( this.state.currentPage - 1) * this.state.moviesPerPage,this.state.currentPage * this.state.moviesPerPage);
    return (
      <div style={{ backgroundColor: "#d2d2d275" }}>
        <List list = {result} catalogueHidden = {false}/>
        <Paginate totalMovies = {this.state.totalMovies} currentPage = {this.state.currentPage}  moviePerPage = {this.state.moviesPerPage} changePage = {this.changePage} />
      </div >
    );
  }
}

export default Home;
