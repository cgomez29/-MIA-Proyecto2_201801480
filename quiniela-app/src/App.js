import React, { useState } from 'react';
import LoginForm from './components/LoginForm'

function App() {
  const adminUser = {
    email: "hola@gmail.com",
    password: "123"
  }

  const [user, setUser] = useState({name:"",email:"",password:""});
  const [error, setError] = useState("");

  const Login = details => {
    console.log(details);

    if(details.email == adminUser.email && details.password == adminUser.password) {
      console.log("Logged in");
      setUser({
        name: details.name,
        email: details.email 
      });

    } else {
      console.log("details do not match!");
      setError("details do not match!");
    }
  } 
  

  const Logout = () => {
    //console.log("Logout");
    setUser({name:"",email:"",password:""});
  }

  return (
    <div className="App">
        {(user.email != "") ? (
          <div className="welcome">
             <h2>Welcome,<span>{user.name}</span></h2>
             <button onClick={Logout}>Logout</button>
          </div> 
        ) : (
          <LoginForm Login={Login} error={error}/>
        )}
    </div>
  );
}

export default App;
