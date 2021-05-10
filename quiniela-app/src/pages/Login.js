import React, {useContext, useState} from 'react'
import {Redirect} from 'react-router-dom'
import url from '../config'
import axios from 'axios'
import {Avatar, Button, Grid, Link, Paper, TextField, Typography} from "@material-ui/core";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import {UserContext} from "../Context/UserContext";

const LoginPage = () => {
    const {setRol, setSrc} = useContext(UserContext)

    const paperStyle = {padding: 25, height:'50vh', width:300, margin: "20px auto"}
    const avatarStyle = { backgroundColor:'green'}
    const btnStyle = { margin:'8px 0' }

    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [redirectUser, setRedirectUser] = useState(false)
    const [redirectAdmin, setRedirectAdmin] = useState(false)

    const onSumbit = async (e) => {
        e.preventDefault()

        const auth = {
            username: username,
            password: password
        }

        const instance = axios.create({
            withCredentials: true,
        })

        await instance.post(`${url}/login`, auth)
            .then(res => {
                //console.log(res)
                if(res.data.idRol === 2){
                    setRedirectUser(true)
                } else {
                    setRedirectAdmin(true)
                }

                setRol(res.data.idRol)
                setSrc(`${url}/img/${res.data.file}`)
                console.log(res.data.file)
            })
    }

    if(redirectUser)
        return <Redirect to="/user"/>

    if(redirectAdmin)
        return <Redirect to="/admin"/>

    return (
        <React.Fragment>
            <form onSubmit={onSumbit}>
                <Paper elevation={10} style={paperStyle}>
                    <Grid align="center">
                        <Avatar style={avatarStyle}><LockOutlinedIcon/></Avatar>
                        <h2>Sign In</h2>
                    </Grid>
                    <TextField label='Username' placeholder="Enter username" fullWidth required onChange={e => setUsername(e.target.value)}/>
                    <TextField label='Password' placeholder="Enter password" type="password" fullWidth required onChange={e => setPassword(e.target.value)}/>
                    <Button type="submit"  color='primary' variant='contained' fullWidth style={btnStyle}>Sign in</Button>
                    <Typography>
                        <Link>
                            Forgot password
                        </Link>
                    </Typography>
                </Paper>
            </form>
        </React.Fragment>

    )
}

export default LoginPage;