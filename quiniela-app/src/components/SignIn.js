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

    // for upload images
    const [file, setFile] = useState([])

    const instance = axios.create({
        withCredentials: true,
    })

    const config = {
        headers: { 'Content-Type': 'multipart/form-data' }
    }

    const onSubmit = async (e: SyntheticEvent) => {
        e.preventDefault()

        const img = file;

        const formData = new FormData()
        formData.append('file', img)
        formData.append('name', "siiiiuu")

        await instance.post(`${url}/upload`, formData, config)
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