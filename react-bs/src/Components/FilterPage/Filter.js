import React from "react";
import { createStyles, makeStyles } from "@material-ui/core/styles";
import InputLabel from "@material-ui/core/InputLabel";
import MenuItem from "@material-ui/core/MenuItem";
import FormControl from "@material-ui/core/FormControl";
import Select from "@material-ui/core/Select";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";


import {Genres, sortBymenu} from '../Helper/selectList'

const useStyles = makeStyles((theme) =>
  createStyles({
    formControl: {
      margin: theme.spacing(3),
      minWidth: 120
    },
    formControlo: {
      margin: theme.spacing(3),
      marginTop: theme.spacing(4),
      minWidth: 120
    }
  })
);

const GenreList = Genres.map(gen => {
  return (
    <MenuItem value={gen} key={gen}>{gen}</MenuItem>
  )
})


const sortByList = sortBymenu.map(sort => {
  return (
    <MenuItem value={sort} key={sort}>{sort}</MenuItem>
  )
}) 

export default function SimpleSelect({ setparams, initparams }) {
  const classes = useStyles();
  const [genre, setGenre] = React.useState(initparams.genre);
  const [sortBy, setSortBy] = React.useState("vote_average DESC");
  const [searchField, setSearchField] = React.useState(initparams.search)
  
  const handleChange = (e) => {
    if (e.target.name === "Genre") {
      setGenre(e.target.value)
    } else if (e.target.name === "SortBy") {
      setSortBy(e.target.value)
    } else {
      setSearchField(e.target.value)
    }
  };

  const handlesubmit = (e) => {
    e.preventDefault()
    setparams({
      search : searchField, 
      genre : genre, 
      sortby : sortBy
    })
  }

  return (
    <div className="container" >
      <div className="card">
        <div className="container">
          <div className="searchbar-container">

            <form onSubmit={handlesubmit}>
              <div className="row">
                <div className="col-12 col-sm-6 col-md-3">
                  <FormControl className={classes.formControl} mx="auto" name="text">
                    <TextField id="standard-basic" label="Search" defaultValue={searchField} onChange={handleChange} />
                  </FormControl>
                </div>
                <div className="col-12 col-sm-6 col-md-3">
                  <FormControl className={classes.formControl} mx="auto">
                    <InputLabel id="Genre-select">Genre</InputLabel>
                    <Select
                      labelId="Genre-select"
                      id="genre"
                      value={genre}
                      onChange={handleChange}
                      name="Genre"
                    >
                      {GenreList}
                    </Select>
                  </FormControl>
                </div>
                <div className="col-12 col-sm-6 col-md-3">
                  <FormControl className={classes.formControl}>
                    <InputLabel id="sortBy-select">SortBy</InputLabel>
                    <Select
                      labelId="sortBy-select"
                      id="sortby"
                      value={sortBy}
                      onChange={handleChange}
                      name="SortBy"
                    >
                      {sortByList}
                    </Select>
                  </FormControl>
                </div>
                <div className="col-12 col-sm-6 col-md-3">
                  <FormControl className={classes.formControlo}>
                    <Button type="submit" variant="contained" color="primary">
                      Filter
              </Button>
                  </FormControl>
                </div>

              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}
