import React, {SyntheticEvent, useState} from 'react';
import {Avatar, Button, Grid, Paper, TextField} from "@material-ui/core";
import AccountCircleOutlinedIcon from '@material-ui/icons/AccountCircleOutlined';
import axios from 'axios'
import url from "../config";

const SignIn = () => {
    // styles
    const paperStyle = {padding: 25, height:'70vh', width:300, margin: "20px auto"}
    const avatarStyle = { backgroundColor:'green'}
    const btnStyle = { margin:'8px 0' }

    //register
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [name, setName] = useState('')
    const [surname, setSurname] = useState('')
    const [email, setEmail] = useState('')
    const [date, setDate] = useState('')

    // for upload images
    const [file, setFile] = useState([])

    const instance = axios.create({
        withCredentials: true,
        headers: { 'Content-Type': 'multipart/form-data' },
    })

    /*const config = {
        headers: { 'Content-Type': 'multipart/form-data' }
    }*/

    const onSubmit = async (e: SyntheticEvent) => {
        e.preventDefault()

        const img = file;

        const formData = new FormData()
        formData.append("username", username)
        formData.append("password", password)
        formData.append("name", name)
        formData.append("surname",surname)
        formData.append("email", email)
        formData.append("date", date)
        formData.append('file', img)

        await instance.post(`${url}/register`, formData)
            .then(res => {
                console.log(res)
            }).catch(err =>{
                console.log(err)
            })

    }

    return (
        <React.Fragment>
            <form onSubmit={e => onSubmit(e)}>
                <Paper elevation={10} style={paperStyle}>
                    <Grid align="center">
                        <Avatar style={avatarStyle}><AccountCircleOutlinedIcon/></Avatar>
                        <h2>Register</h2>
                    </Grid>
                    <TextField label='Username' placeholder="Enter username" fullWidth required onChange={e => setUsername(e.target.value)}/>
                    <TextField label='Password' placeholder="Enter password" type="password" fullWidth required onChange={e => setPassword(e.target.value)}/>
                    <TextField label='Name' placeholder="Enter name" fullWidth required onChange={e => setName(e.target.value)}/>
                    <TextField label='Surname' placeholder="Enter surname" fullWidth required onChange={e => setSurname(e.target.value)}/>
                    <TextField label='Email' placeholder="Enter email" fullWidth required onChange={e => setEmail(e.target.value)}/>
                    <TextField
                        id="date"
                        label="Birthday"
                        type="date"
                        defaultValue="2015-05-24"
                        InputLabelProps={{
                            shrink: true,
                        }}
                        fullWidth
                        required
                        onChange={e => setDate(e.target.value) }
                    />
                    <Button variant="contained" component="label" fullWidth style={btnStyle}>
                        Upload photo
                        <input type="file" name="file"hidden onChange={e => setFile(e.target.files[0])} />
                    </Button>
                    <Button type="submit"  color='primary' variant='contained' fullWidth style={btnStyle}
                    >Sign in</Button>
                </Paper>
            </form>
        </React.Fragment>
    );
};
export default SignIn;