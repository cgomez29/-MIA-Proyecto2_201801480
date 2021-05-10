import React, {useContext, useEffect, useState} from 'react';
import {Avatar, Button, Grid, Paper, TextField} from "@material-ui/core";
import {UserContext} from "../../Context/UserContext";
import url from "../../config";
import axios from "axios";
import moment from "moment"

const UpdateUser = () => {
    const [file, setFile] = useState(null)
    const [data, setData] = useState({
        username: '',
        password: '',
        name: '',
        surname: '',
        email: '',
        date: '',
        file: '',
    })

    //Context
    const {src} = useContext(UserContext)

    // styles
    const paperStyle = {padding: 25, height:700, width:350, margin: "20px auto"}
    const avatarStyle = { backgroundColor:'green', width:200, height: 200}
    const btnStyle = { margin:'8px 0', }

    const instance = axios.create({
        withCredentials: true,
        headers: { 'Content-Type': 'multipart/form-data' },
    })

    const onSubmit = async (e) => {
        e.preventDefault()
        let fecha = moment(data.date).format("YYYY/MM/DD")
        const formData = new FormData()
        formData.append("username", data.username)
        formData.append("password", data.password)
        formData.append("name", data.name)
        formData.append("surname",data.surname)
        formData.append("email", data.email)
        formData.append("fecha", fecha)

        if (file != null) {
            formData.append('file', file)
        } else {
            formData.append('file', data.file)
        }

        await instance.put(`${url}/user`, formData)
            .then(res => {
                console.log(res)
            }).catch(err =>{
                console.log(err)
            })
    }

    const loadUser = async () => {
        const instance = axios.create({
            withCredentials: true,
            headers: { 'Content-Type': 'multipart/form-data' },
        })
        await instance.get(`${url}/user`).then(res => {
            setData(res.data)
        })
    }

    useEffect(() => {
        loadUser();
    }, [])

    const handleInputChange = (e:InputEvent) => {
        setData({
            ...data,
            [e.target.name]:e.target.value
        })
    }

    return (
        <React.Fragment>
            <form onSubmit={onSubmit}>
                <Paper elevation={10} style={paperStyle}>
                    <Grid align="center">
                        <Avatar style={avatarStyle} src={src}/>
                        <h2>Modificar datos personales</h2>
                    </Grid>
                    <TextField label='Username'
                               placeholder='Username'
                               name='username'
                               fullWidth required
                               value={data.username}
                               onChange={handleInputChange}
                    />
                    <TextField label='Password'
                                placeholder='Password'
                               type="password"
                               name='password'
                               fullWidth required
                               onChange={handleInputChange}
                               value={data.password}
                    />
                    <TextField label='Name'
                               placeholder='Name'
                               name='name'
                               fullWidth required
                               value={data.name}
                               onChange={handleInputChange}
                    />
                    <TextField label='Surname'
                               placeholder='Surname'
                               name='surname'
                               fullWidth required
                               value={data.surname}
                               onChange={handleInputChange}
                    />
                    <TextField label='Email'
                               placeholder='Email'
                               fullWidth required
                               name='email'
                               value={data.email}
                               onChange={handleInputChange}
                    />
                    <TextField
                        id="date"
                        label="Birthday"
                        type="date"
                        name="date"
                        defaultValue={moment(data.date).format("YYYY-MM-DD")}
                        value={moment(data.date).format("YYYY-MM-DD")}
                        InputLabelProps={{
                            shrink: true,
                        }}
                        fullWidth
                        required
                        onChange={handleInputChange}
                    />
                    <Button variant="contained" component="label" fullWidth style={btnStyle}>
                        Upload photo
                        <input type="file" name="file"hidden onChange={e => setFile(e.target.files[0])} />
                    </Button>
                    <Button type="submit"  color='primary' variant='contained' fullWidth style={btnStyle}
                    >Guardar</Button>
                </Paper>
            </form>
        </React.Fragment>
    );
};

export default UpdateUser;