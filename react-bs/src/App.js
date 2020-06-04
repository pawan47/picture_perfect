import React, { Component } from 'react';
import {BrowserRouter, 
  Route
} from 'react-router-dom'


// local imports
import './App.css';
import Navbartwo from './Components/CommonComponents/Navbartwo'
import Footer from './Components/CommonComponents/Footer'
import Home from './Components/HomePage/NewHome'
import SearchPage from './Components/FilterPage/newSearchPage'
import movieInfo from './Components/MoviePage/movieInfo'

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div style={{ backgroundColor: "rgb(208, 208, 208)" }}>

          <Navbartwo />
          <Route exact path = '/' component = {Home} />
          <Route path = '/searchpage' component = {SearchPage} />
          <Route path = '/movie/:movieId' component = {movieInfo} />
          <Footer />
        </div >
      </BrowserRouter>
    );
  }
}

export default App;
