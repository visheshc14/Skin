import React, {useState} from 'react';
import {useHistory} from 'react-router-dom'
import {requestForTGS, requestForAS, requestForServer} from "../../api/userApi";
import decryptMessages from "../../utils/decryptMessages";
import encryptMessages from "../../utils/encryptMessages";
import './signin.css'

function SignIn(props) {

    // console.log(props);
    const {setIsLogout, isLogout} = props;
    console.log(isLogout);
    const history = useHistory();

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const [signinResponse, setSigninResponse] = useState("");

    async function Server(messageForServer) {

        try {
            console.log("Inside Server calling api");
            const { data } = await requestForServer(messageForServer);
            console.log(data);

            if(data.success){
                console.log(data.success);
                setIsLogout(false);
                localStorage.setItem("token", data.accessToken);
                console.log(history);
                history.push("/");
            }
            
        } catch (err) {
            console.log(err);
            setSigninResponse({message: err})
        }

    }

    async function TGS(messageForTGS, TGS_SessionKey){
        try {
            console.log("inside TGS calling api");
            const { data } = await requestForTGS(messageForTGS);
            console.log(data);
            const encryptedServiceTicket = data.encryptedServiceTicket;
            const encryptedUserCredentials = data.encryptedUserCredentials;
            const serviceSecretKey = data.serviceSecretKey;

            const { username, timestamp, Service_SessionKey } = decryptMessages(encryptedUserCredentials, TGS_SessionKey);

            const timeElapsed = Date.now();
            const todaysDate = new Date(timeElapsed);
            const timeStamp = todaysDate.toUTCString();
            
            const userCredentials = {
                username,
                timeStamp
            };

            const newEncryptedUserCredentials = encryptMessages(userCredentials, Service_SessionKey);
            
            const messageForServer = {
                newEncryptedUserCredentials,
                encryptedServiceTicket,
                serviceSecretKey
            };

            Server(messageForServer);

        } catch (err) {
            console.log(err);
            setSigninResponse({message: err});
        }
    }

    async function handleSubmit(event) {
        event.preventDefault();

        const userData = {
            username,
            password
        };

        try {
            const { data } = await requestForAS(userData);
            console.log("data: ", data);
            setSigninResponse(data);
            // console.log("signinResponse: ", signinResponse);
            
            const encryptedClientMessage = data.encryptedClientMessage;
            // console.log(encryptedClientMessage);
            const clientSecretKey = data.clientSecretKey;  
            // console.log(clientSecretKey);
            const encryptedTGT = data.encryptedTGT;
            // console.log(encryptedTGT);
            const { TGS_ID, TGS_SessionKey, timestamp } = decryptMessages(encryptedClientMessage, clientSecretKey);
            // console.log(TGS_ID, timestamp, TGS_SessionKey);

            // console.log(TGS_SessionKey);
            const timeElapsed = Date.now();
            const todaysDate = new Date(timeElapsed);

            const userCredentials = {
                username,
                timeStamp: todaysDate.toUTCString()
            };

            const encryptedUserCredentials = encryptMessages(userCredentials, TGS_SessionKey);

            const messageForTGS = {
                encryptedTGT,
                encryptedUserCredentials
            };
            
            console.log(messageForTGS);

            TGS(messageForTGS, TGS_SessionKey);

        } catch(err) {
            // console.log("inside");
            console.log(err);
            setSigninResponse({message: err});
        }
    }

    return (
        
        <div className="content">
            <div className="signin-content">
                <div className="main main-raised">
                    <div className="container signin">
                    <div  className="signin-heading">
                        <h1 className="heading-per-page">Login</h1>
                    </div>
                        <form onSubmit={handleSubmit} className="signin-form">
                            <div className="input-values">
                                {/* {
                                    (!signinResponse.success) ?  <div className="err-msg"><div className="msg">{signinResponse.message}</div></div> : <div className="err-msg"></div>
                                } */}
                                <div className="input">
                                    <input type="text" name="username" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} /><br/> 
                                    <input type="password" name="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} /><br/>
                                </div>
                            </div>
                            <div className="submit-btn" Style={"margin-top: 1rem"}>
                                <button type="submit" className="submit-btn">Sign In</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>

    );
}

export default SignIn;
