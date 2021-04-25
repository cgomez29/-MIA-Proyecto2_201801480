import React from 'react';
import {Link} from "react-router-dom"
import axios from "axios";
import url from "../config";

const Nav = (props: {name: string, setName: (name: string) => void }) => {
    const instance = axios.create({
        withCredentials: true,
    })

    const logout = async () => {
        await instance.post(`${url}/logout`, {})
            .then(res => {
                //console.log(res)
            })
        props.setName('')
    }

    let menu;

    if(props.name===''){
        menu = (
            <ul className="nav navbar-nav">
                <li className="active">
                    <Link to="/">HOME</Link>
                </li>
                <li>
                    <Link to="/login">Login</Link>
                </li>
                <li>
                    <Link to="/login">Register</Link>
                </li>
            </ul>
        )
    } else {
        menu = (
            <ul className="nav navbar-nav">
                <li className="active">
                    <Link to="/">HOME</Link>
                </li>
                <li>
                    <Link to="/login" onClick={logout}>Logout</Link>
                </li>
            </ul>
        )
    }


    return (
        <nav className="navbar navbar-default">
            <div className="container-fluid">
                <div className="navbar-header">
                </div>
                {menu}
            </div>
        </nav>
    );
};

export default Nav;