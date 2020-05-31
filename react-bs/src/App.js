import React, { Component } from 'react';
import {BrowserRouter, 
  Route
} from 'react-router-dom'


// local imports
import './App.css';
import Navbartwo from './Components/CommonComponents/Navbartwo'
import Footer from './Components/CommonComponents/Footer'
import Home from './Components/HomePage/Home'
import SearchPage from './Components/FilterPage/SearchPage'


class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div style={{ backgroundColor: "#d2d2d275" }}>

          <Navbartwo />
          <Route exact path = '/' component = {Home} />
          <Route path = '/searchpage' component = {SearchPage} />
          <Footer />
        </div >
      </BrowserRouter>
    );
  }
}

export default App;
