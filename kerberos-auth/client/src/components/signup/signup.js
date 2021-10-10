import React, { useState } from 'react';
import {Redirect, useHistory} from 'react-router-dom';
import {signupNewUser} from '../../api/userApi';
import './signup.css';

function SignUp(props) {

    const history = useHistory();

    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const [signupResponse, setSignupResponse] = useState("");

    async function handleSubmit(event){

        console.log("inside handel submit");
        event.preventDefault();
        const userData = {
            username,
            email,
            password
        };
 
        try {
            const { data } = await signupNewUser(userData);
            setSignupResponse(data);
            console.log(data);
            console.log(signupResponse);
            console.log(userData);

            if(data.success){
                console.log("Submitted successfully", " user: ", userData);
                history.push("/");
                return <Redirect to="/"/>
            }
            
        } catch (err) {
            setSignupResponse({message: err});
        }
    };

    
    return (
        
        <div className="content">
            <div className="signup-content">
                <div className="main main-raised">
                    <div className="container signup">
                        <div  className="signup-heading">
                            <h1 className="heading-per-page">Signup</h1>
                        </div>
                        <form onSubmit={handleSubmit} className="signup-form">
                            <div className="input-values">
                                {
                                    (!signupResponse.success) ?  <div className="err-msg"><div className="msg">{signupResponse.message}</div></div> : <div className="err-msg"></div>
                                }
                                <div className="input">
                                    <input type="text" name="username" placeholder="User Name" value={username} onChange={(e) => setUsername(e.target.value)}/><br/> 
                                    <input type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)}/><br/> 
                                    <input type="password" name="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)}/><br/>
                                </div>
                            </div>
                            <div className="submit-btn" Style={"margin-top: 1rem"}>
                                <button type="submit">Sign Up</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>

    );
    
}
 
export default SignUp;