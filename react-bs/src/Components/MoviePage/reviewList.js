import React from 'react'
import {Row, Col} from 'react-bootstrap'

const ReviewList = ({review}) => 
    {
        return(
            review ? (
                <div style={{ backgroundColor: "rgba(255, 255, 255, 0.53)" }} className="ml-2 mt-4 mb-4 row">
                <Row>
                    <Col className="m-3"><h4 style={{ margin: 10 }}>{review.length} Comments</h4></Col>
                </Row>
                <br></br>
                <ul style={{ listStyle: "none" }}>
                    {review.map(comments => {
                        return (
                            <li className="mt-2 mr-4" key={comments.id} >
                                <span className="blue" style={{ color: "blue" }}><strong>{comments.name}</strong></span>
                                <hr />
                                <h5 className="text-justify" style={{ fontStyle: "italic" }}>{comments.body}</h5>
                                <br />
                            </li>
                        )
                    }
    
                    )}
                </ul>
    
            </div>
    
            ) : (
                null
            )
        )
    }


export default ReviewList;