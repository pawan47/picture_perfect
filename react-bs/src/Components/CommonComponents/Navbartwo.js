import React from 'react'
import Navbar from 'react-bootstrap/Navbar'
import NavDropdown from 'react-bootstrap/NavDropdown'
import Nav from 'react-bootstrap/Nav'
import Button from 'react-bootstrap/Button'
import {Link} from 'react-router-dom'

// local imports
import {Genres} from '../Helper/selectList'
import logo from '../../images/logo.jpg'


const GenreList = Genres.map(gen => {
  return (
    // <NavDropdown.Item as = {Link} to = {"/searchpage?genre=" + gen} key = {gen}>{gen}</NavDropdown.Item>
    <NavDropdown.Item href = {"/searchpage?genre=" + gen} key = {gen}>{gen}</NavDropdown.Item>
  )
})


function Navbartwo() {
    return (
        <Navbar bg="light" expand="lg" variant="light">
        <div className = "container">
            <Navbar.Brand href ="/">
                <img src={logo} loading = "lazy" alt='logo' style={{ width: 40 }} />{''} <span style={{ color: "black" }}>Picture Perfect</span>
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav"></Navbar.Toggle>
            <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="mr-auto">
                    <NavDropdown title="Genre" style={{ color: "white" }} id="basic-nav-dropdown" >
                        {GenreList}
                        </NavDropdown>
                    <Nav.Link as={Link} disabled to = "/show"><span style={{ color: "grey", fontSize: "110%" }}>Shows</span></Nav.Link>

                </Nav>
                <Nav className="ml-auto">

                    <Nav.Link href="#login"><Button disabled variant="outline-primary" >SignUp</Button></Nav.Link>
                    <Nav.Link href="#login"><Button disabled variant="outline-primary">LogIn</Button></Nav.Link>

                </Nav>
            </Navbar.Collapse>
            </div>
        </Navbar>

    )
}

export default Navbartwo