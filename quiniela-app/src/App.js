import React, {useEffect, useState} from 'react';
import { BrowserRouter as Router, Route, Switch} from "react-router-dom";
import Home from "./pages/Home"
import LoginPage from "./pages/Login"
import Nav from "./components/Nav"
import NotFound from "./components/NotFound"
import axios from "axios";
import url from './config'
import SignIn from "./components/SignIn";
import Temporada from "./pages/Admin/Temporada";

const App = () => {
    const [name, setName] = useState('')

    const instance = axios.create({
        withCredentials: true,
    })

    useEffect(()=>{
        (
            async () => {
                await instance.get( `${url}/user`)
                    .then(res => {
                        setName(res.data.username)
                    })
            }
        )();
    })

    return (
        <Router>
            <Nav name={name} setName={setName}/>
            <Switch>
                <Route exact path="/" component={Home}/>
                <Route exact path="/Login" component={() => <LoginPage setName={setName}/>}/>
                <Route exact path="/register" component={SignIn}/>
                <Route exact path="/temporada" component={Temporada}/>
                <Route component={NotFound}/>
            </Switch>
        </Router>
    )
}
export default App