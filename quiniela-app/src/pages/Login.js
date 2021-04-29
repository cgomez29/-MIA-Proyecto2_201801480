import React, {SyntheticEvent ,useState} from 'react'
import {Redirect} from 'react-router-dom'
import url from '../config'
import axios from 'axios'
import {Avatar, Button, Grid, Link, Paper, TextField, Typography} from "@material-ui/core";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";

const LoginPage = (props: {setName: (name: string) => void }) => {
    const paperStyle = {padding: 25, height:'50vh', width:300, margin: "20px auto"}
    const avatarStyle = { backgroundColor:'green'}
    const btnStyle = { margin:'8px 0' }

    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [redirect, setRedirect] = useState(false)

    const onSumbit = async (e: SyntheticEvent) => {
        e.preventDefault()

        const auth = {
            email: email,
            password: password
        }

        const instance = axios.create({
            withCredentials: true,
        })

        await instance.post(`${url}/login`, auth)
            .then(res => {
                //console.log(res)
                setRedirect(true)
                props.setName(res.data.username)
            })
    }


    if(redirect)
        return <Redirect to="/"/>

    return (
        <React.Fragment>
            <form onSubmit={onSumbit}>
                <Paper elevation={10} style={paperStyle}>
                    <Grid align="center">
                        <Avatar style={avatarStyle}><LockOutlinedIcon/></Avatar>
                        <h2>Sign In</h2>
                    </Grid>
                    <TextField label='Email' placeholder="Enter email" fullWidth required onChange={e => setEmail(e.target.value)}/>
                    <TextField label='Passwor' placeholder="Enter password" type="password" fullWidth required onChange={e => setPassword(e.target.value)}/>
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