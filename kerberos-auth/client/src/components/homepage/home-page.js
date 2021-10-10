import React from 'react';
import './home-page.css';

function HomePage(props) {
    return ( 
        <div className="home-page content">
            <div className="container intro">
                <h1 className="heading-per-page">Kerberos Authentication</h1>
                <div className="container blog-description">
                    <p>
                        This project represents login/signup system which works on the principles of Kerberos Authentication.
                    </p>
                    <br/>
                    <p>
                        For Web Security Class
                    </p>
                    <br />
                    <p>
                        Team:<br />
                        <b>Aditi Jain</b><br />
                        <b>Pranav Singh</b><br />
                        <b>Vishesh Choudhary</b><br />
                    </p>
                </div>
            </div>
        </div>
    );
}
 
export default HomePage;