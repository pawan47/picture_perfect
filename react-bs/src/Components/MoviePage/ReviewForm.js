import React from 'react'
import { makeStyles } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
// import { FormControl } from "@material-ui/core";
import Button from "@material-ui/core/Button";
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'


const useStyles = makeStyles(theme => ({
    root: {
        "& .MuiTextField-root": {
            margin: theme.spacing(1)
            // width: "100%"
        }
    },
    button: {
        flip: false,
        direction: "rtl"
    },
}));


const ReviewForm = ({movieName, UserReviewPosted, handlesubmit}) =>{ 
    const classes = useStyles();
    const [review, setreview] = React.useState("");
    const [name, setname] = React.useState("");

    const handleChange = event => {
        setname(event.target.value);
    };

    const handleChangereview = e => {
        setreview(e.target.value);
    };



    return(movieName ? (
        UserReviewPosted ? (
    null
) : (
    
    <div style={{ backgroundColor: "rgba(255, 255, 255, 0.53)" }} className="ml-2 mt-4 mb-4 row">
                
        <Row>
            <Col className="ml-4 mt-3" lg={12}>
                <h4> Post Your Review</h4>
            </Col>
            <Col className="">
                <form
                    className={classes.root}
                    style={{ margin: 15 }}
                    onSubmit={handlesubmit}
                >
                    <TextField
                        label="Name"
                        id="outlined-margin-dense"
                        className={classes.textField}
                        value={name}
                        // helperText="Some important text"
                        margin="dense"
                        variant="outlined"
                        required
                        onChange={handleChange}
                    />
                    <TextField
                        id="outlined-multiline-flexible"
                        label="Review"
                        multiline
                        rowsMax={4}
                        value={review}
                        onChange={handleChangereview}
                        variant="outlined"
                        rows={1}
                        fullWidth
                        required
                    />

                    <Button
                        style={{ margin: 10 }}
                        variant="outlined"
                        color="primary"
                        type="submit"
                    >
                        Post
        </Button>
                </form>
            </Col>
        </Row>
        </div>
    )) : (
        null
    )
    )
}


export default ReviewForm;