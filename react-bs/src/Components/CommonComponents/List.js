import React from 'react'
import Col from 'react-bootstrap/Col'
import Container from 'react-bootstrap/Container'

// local imports
import CatalogueButton from '../HomePage/CatalogueButton'
import DNE from '../../images/Poster_not_available.jpg'

const ProcessImage = (imglink, imageSize) => {
    const arr = imglink.split("/")
    if(arr.length !== 7){
        return DNE
    }
    arr[5] = imageSize
    imglink =  arr.join('/')
    return imglink
}


function List({ list, catalogueHidden}) {
    const postlist = list.length ? (
        list.map(movie => {
            return (
                <Col xs={12} sm={6} md={3} large={2} xl={2} className="mb-3 mt-3 justify-content-center text-center" key = {movie.movie_id}>

                    <div className='box '>
                        <a href={ProcessImage(movie.thumbnail_link, "w154")}>
                            <img className=' mb-3' src={ProcessImage(movie.thumbnail_link, "w154")} loading = "lazy" alt={movie.movie_id}
                                style={{ borderRadius: 10, width: 130, height: 194 }} />
                        </a>
                    </div>
                </Col>
            )
        })
    ) : (
            <div >
                <Col xl = {12} className="mb-3 mt-3 justify-content-center text-center" key = {-1}>No movie </Col>
            </div>
        )


    return (
        <div className='container body-cont' style={{ paddingTop: 5, paddingBottom: 15 }}>
            <main>
                <Container>
                    <CatalogueButton hidden = {catalogueHidden} /> 
                    <div className='row'>
                        {postlist}
                    </div>
                </Container>
            </main>
        </div>
    )
}

export default List;