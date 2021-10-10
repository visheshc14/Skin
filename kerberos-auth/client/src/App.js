import React, {useState, useEffect} from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import SignUp from './components/signup/signup';
import SignIn from './components/signin/signin';
import HomePage from './components/homepage/home-page';
import MyNavbar from './components/navbar/navbar';
import ProfilePage from './components/profilepage/profile-page';

function App(props) {

  const [isLogout, setIsLogout] = useState(true);

  useEffect(() => {
    if(localStorage.getItem("token")){
      setIsLogout(false);
    } else {
      setIsLogout(true);
    }
  }, [isLogout]);

  const logOutHandle = () => {
    setIsLogout(true);
    localStorage.removeItem("token");
  }

  return (
    <Router>
        <MyNavbar logOutHandle={logOutHandle} isLogout={isLogout}/>
        <Switch>
          <Route path='/' component={HomePage} exact />
          <Route path='/signup' component={SignUp} />
          <Route path="/signin" render={(props) => (<SignIn {...props} isLogout={isLogout} setIsLogout={setIsLogout} />)}/>
          <Route path='/profile' render={(props) => (<ProfilePage {...props} isLogout={isLogout} setIsLogout={setIsLogout} />)}/>
        </Switch>
    </Router>
  );
}

export default App;
