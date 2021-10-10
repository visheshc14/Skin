import React from 'react';
import { Navbar, Container, Nav } from 'react-bootstrap';
import {BrowserRouter as Router} from 'react-router-dom'
import './navbar.css';

function MyNavbar(props) { 

    console.log(props);
    const {logOutHandle, isLogout} = props;
    console.log(isLogout);

    return ( 
        <div>
            <Router>
                <Navbar className="navbar">
                    <Container>
                        <Navbar.Brand href="/">homepage</Navbar.Brand>
                        <Nav className="me-auto">
                            { !isLogout && (
                                <React.Fragment>
                                    <Nav.Link href='/profile'>Profile</Nav.Link>
                                    <Nav.Link href="/" onClick={logOutHandle}>Log Out</Nav.Link> 
                                </React.Fragment> 
                            )}
                            { isLogout && (
                                <React.Fragment>
                                    <Nav.Link href="/signup">Register</Nav.Link>
                                    <Nav.Link href="/signin">Login</Nav.Link> 
                                </React.Fragment> 
                            )}
                        </Nav>
                    </Container>
                </Navbar>
            </Router>
        </div>
    );
}
 
export default MyNavbar;