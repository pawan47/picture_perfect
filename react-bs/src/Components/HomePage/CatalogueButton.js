import React from 'react'
import ArrowForwardIosIcon from '@material-ui/icons/ArrowForwardIos';
import {Link} from 'react-router-dom'
const CatalogueButton = ({hidden}) => {
    // console.log(hidden)
    if( hidden ){
        return null
    }
    return (
        <Link to = "/searchPage" ><button type='button' className='btn btn-outline-primary my-3' >
                        Catalogue {' '}
                        <ArrowForwardIosIcon fontSize="small" />
                    </button>
                    </Link>
    )
}


export default CatalogueButton;