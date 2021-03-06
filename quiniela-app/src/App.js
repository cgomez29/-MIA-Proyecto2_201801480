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
import Deporte from "./pages/Admin/Deporte";
import Jornada from "./pages/Admin/Jornada";
import HomeAdmin from './pages/Admin/Home'
import HomeUser from './pages/User/Home'
import PerfilUser from './pages/User/Perfil'
import MembresiaUser from './pages/User/Membresia'
import RecompensaUser from './pages/User/Recompensa'
import EventoUser from './pages/User/Evento'
import {UserContext} from "./Context/UserContext";
import Bulkload from "./pages/Admin/Bulkload";

const App = () => {
    const [rol, setRol] = useState('')
    const [src, setSrc] = useState('')

    const instance = axios.create({
        withCredentials: true,
    })

    useEffect(()=>{
        const axiosDeportes = async () => {
            await instance.get( `${url}/user`)
                .then(res => {
                    setRol(res.data.idRol)
                    setSrc(`${url}/img/${res.data.file}`)
                })
        }
        axiosDeportes()
    })

    return (
        <Router>
            <UserContext.Provider value={{rol, setRol, src, setSrc}}>
                <Nav/>
                <Switch>
                    <Route exact path="/" component={Home}/>
                    <Route exact path="/Login" component={LoginPage}/>
                    <Route exact path="/register" component={SignIn}/>

                    <Route exact path="/admin" component={HomeAdmin}/>
                    <Route exact path="/admin/temporada" component={Temporada}/>
                    <Route exact path="/admin/deporte" component={Deporte}/>
                    <Route exact path="/admin/jornada" component={Jornada}/>
                    <Route exact path="/admin/bulkload" component={Bulkload}/>

                    <Route exact path="/user" component={HomeUser}/>
                    <Route exact path="/user/evento" component={EventoUser}/>
                    <Route exact path="/user/recompensa" component={RecompensaUser}/>
                    <Route exact path="/user/perfil" component={PerfilUser}/>
                    <Route exact path="/user/Membresia" component={MembresiaUser}/>

                    <Route component={NotFound}/>

                </Switch>
            </UserContext.Provider>
        </Router>
    )
}
export default App