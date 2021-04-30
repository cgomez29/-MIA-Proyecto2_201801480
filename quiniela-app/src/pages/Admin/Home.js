import React, {useEffect, useState} from 'react';
import {Grid} from "@material-ui/core";
import CardG from "../../components/CardG";
import axios from "axios";
import url from "../../config";
import {Redirect} from "react-router-dom";

const Home = () => {
    const [loggedIn, setLoggedIn] = useState(false)
    const  [data, setData] = useState([])

    const instance = axios.create({
        withCredentials: true,
    })

    useEffect(()=>{
        (
            async () => {
                await instance.get( `${url}/deporte`)
                    .then(res => {
                        setData(res.data)
                    }).catch( err =>{
                        setLoggedIn(true)
                    })
            }
        )();

    })

    if (loggedIn)
        return <Redirect to="/login"/>

    return (
        <React.Fragment>
            <Grid align="center">
                {
                    data.map((deporte) => (
                        <CardG key={deporte.id}
                               {...deporte}
                        />
                    ))
                }
            </Grid>
        </React.Fragment>
    );
};

export default Home;