import React, { useState, useEffect } from 'react'
import Typography from '@material-ui/core/Typography';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import Link from '@material-ui/core/Link';
import { connect } from 'react-redux'
import Col from 'react-bootstrap/Col'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Rating from "@material-ui/lab/Rating";
import StarBorderIcon from "@material-ui/icons/StarBorder";



import ReviewForm from './ReviewForm'
import ReviewList from './reviewList'
import fetchMovieDetails from './../../Actions/fetchMovieDetails'
import addReview from './../../Actions/addReview'
import ProcessImage from './../CommonComponents/ProcessImage'





const MovieInfo = (props) => {

    const [userRating, setuserRating] = useState(0)
    const movieId = props.match.params.movieId
    const { domain } = props;
    const getMovieInfo = props.getMovieInfo;
    const [UserReviewPosted, setUserReviewPosted] = useState(false)
    let Ree = props.review

    const handlesubmit = e => {
        e.preventDefault();
        Ree.unshift({
            name: e.target[0].value,
            body: e.target[2].value,
            postid: 101,
            id: 101,
            email: "pawanagr@iitk.ac.in"
        })
        setUserReviewPosted(true)
    };


    useEffect(() => {
        getMovieInfo(domain, movieId)
    }, [getMovieInfo, movieId, domain])



    const bread = (
        <Breadcrumbs aria-label="breadcrumb" className="mt-4 mb-4">
            <Link color="inherit" href="/" >
                HomePage
        </Link>
            <Typography color="textPrimary">{props.movieName}</Typography>
        </Breadcrumbs>
    )


    const Moviecomponent = props.movieName ? (
        <div >
            <Row>
                <Col className = "ml-2">
                    {bread}
                </Col>
                </Row>
            <Row className="ml-2 mt-2 mb-4" style={{ backgroundColor: "rgba(255, 255, 255, 0.53)" }} >
                <Col className="col-12 col-sm-3 col-md-2 col-xl-2 col-lg-2 mt-3 mb-3 justify-content-center text-center">
                    <img src={ProcessImage(props.poster, "w154")} alt="movie Poster" height="250px" className="responsive"></img>
                </Col>

                <Col className="col mt-2 ml-0 mb-4 mt-3 mb-3">
                    <Row>
                        <Col className="col-12 col-sm-6 col-md-9 col-lg-9 ml-0"><h2>{props.movieName}</h2></Col>
                        <Col className="justify-content-center text-center">
                            <Typography component="legend">Your Rating</Typography>
                            <Rating
                                name="size-small"
                                defaultValue={userRating}
                                precision={0.5}
                                emptyIcon={<StarBorderIcon fontSize="inherit" />}
                                onChange={(value, v) => {
                                    setuserRating(v);
                                }}
                            />
                        </Col>
                    </Row>
                    <Row className=" mt-4 mb-3">
                        <Col className="col-12">{props.overview}</Col>


                        <Col className="col-12 col-sm-6 mt-1">
                            <strong>Genre</strong> : {props.genre}
                        </Col>
                        <Col className="col-12 col-sm-6 mt-1">
                            <strong>Language</strong> : {props.language}
                        </Col>
                        <Col className="col-12 col-sm-6 mt-1">
                            <strong>Rating</strong> : {props.rating}
                        </Col>
                        <Col className="col-12 col-sm-6 mt-1">
                            <strong>Popularity</strong> : {props.popularity}
                        </Col>

                    </Row>
                </Col>
            </Row>
        </div>

    ) : (
            null
        )



    return (
        <div style={{ minHeight: 700 }}>
            <Container>
                <Row>
                    <Col>
                    {Moviecomponent}
                    </Col>
                </Row>
            </Container>
            <Container>
                <ReviewForm movieName={props.movieName} UserReviewPosted={UserReviewPosted} handlesubmit={handlesubmit} />
                <ReviewList review={Ree} />
            </Container>

        </div>
    )
}


const mapStateToProps = state => {
    return (
        {
            movieId: state.movieDetail.movieId,
            movieName: state.movieDetail.movieName,
            poster: state.movieDetail.poster,
            rating: state.movieDetail.rating,
            overview: state.movieDetail.overview,
            genre: state.movieDetail.genre,
            language: state.movieDetail.language,
            domain: state.movies.domain,
            popularity: state.movieDetail.popularity,
            review: state.movieDetail.review,
        })
}


const mapDispatchToProps = dispatch => {
    return ({
        getMovieInfo: (domain, movieId) => dispatch(fetchMovieDetails(domain, movieId)),
        addReview: (name, review) => dispatch(addReview(name, review)),
    })
}

export default connect(mapStateToProps, mapDispatchToProps)(MovieInfo);