import React from 'react';
import {Avatar, Button, Grid, Paper, TextField} from "@material-ui/core";
import AccountCircleOutlinedIcon from '@material-ui/icons/AccountCircleOutlined';


const SignIn = () => {
    const paperStyle = {padding: 25, height:'70vh', width:300, margin: "20px auto"}
    const avatarStyle = { backgroundColor:'green'}
    const btnStyle = { margin:'8px 0' }

    const [selectedDate, setSelectedDate] = React.useState(new Date('2014-08-18T21:11:54'));

    const handleDateChange = (date) => {
        setSelectedDate(date);
    };

    return (
        <React.Fragment>
            <form>
                <Paper elevation={10} style={paperStyle}>
                    <Grid align="center">
                        <Avatar style={avatarStyle}><AccountCircleOutlinedIcon/></Avatar>
                        <h2>Register</h2>
                    </Grid>
                    <TextField label='Username' placeholder="Enter username" fullWidth required />
                    <TextField label='Password' placeholder="Enter password" type="password" fullWidth required/>
                    <TextField label='Name' placeholder="Enter name" fullWidth required />
                    <TextField label='Surname' placeholder="Enter surname" fullWidth required />

                    <TextField label='Email' placeholder="Enter email" fullWidth required />
                    <Button type="submit"  color='primary' variant='contained' fullWidth style={btnStyle}>Sign in</Button>
                </Paper>
            </form>
        </React.Fragment>
    );
};
export default SignIn;