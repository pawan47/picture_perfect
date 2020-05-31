import React, { Component } from 'react';
import axios from 'axios'

//local imports
import List from './../CommonComponents/List'
import Paginate from './../CommonComponents/Paginate'
import SimpleSelect from './Filter'
import GetInitParams from './GetInitParams'


class SearchPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            currentPage: 1,
            moviesPerPage: 24,
            result: [],
            totalMovies: 0,
            search: GetInitParams("search", this.props.location.search),
            genre: GetInitParams("genre", this.props.location.search),
            sortBy: "vote_average DESC",
            domain: "http://13.232.168.173",
        }
    }

    getquery() {
        const query = [];
        if (this.state.search !== "") {
            query.push('search=' + this.state.search)
        }
        if (this.state.genre !== "All") {
            query.push('genre=' + this.state.genre)
        }
        query.push('sortby=' + this.state.sortBy)
        return this.state.domain + '/movie/catalogue?' + query.join('&')
    }


    getMovies() {
        axios.get(this.getquery())
            .then(res => {
                if (res.data !== null) {
                    this.setState({
                        result: res.data,
                        totalMovies: res.data.length,
                    }
                    )
                } else {
                    this.setState({
                        result: [],
                        totalMovies: 0,
                    }
                    )
                }
            }
            )
    }

    componentDidMount() {
        this.getMovies()
    }


    // newUpdateParam = (key, param) => {
    //     this.setState({
    //         [key]: param,
    //     },
    //         () => this.getMovies()
    //     )
    // }

    updateParams = (search, genre, sortBy) => {
        this.setState({
            search: search,
            genre: genre,
            sortBy: sortBy,
        },
            () => this.getMovies()
        )
    }

    changePage = (pageno) => {
        this.setState({
            currentPage: pageno,
        })
    }

    render() {


        const result = this.state.result.slice((this.state.currentPage - 1) * this.state.moviesPerPage, this.state.currentPage * this.state.moviesPerPage);

        const initparams = {
            search: this.state.search,
            genre: this.state.genre,
        }

        return (
            <div style={{ backgroundColor: "#d2d2d275" }}>

                <SimpleSelect updateParams={this.updateParams} initparams={initparams} />
                <List list={result} catalogueHidden={true} />
                <Paginate totalMovies={this.state.totalMovies} currentPage={this.state.currentPage} moviePerPage={this.state.moviesPerPage} changePage={this.changePage} />

            </div >
        );
    }
}

export default SearchPage;
