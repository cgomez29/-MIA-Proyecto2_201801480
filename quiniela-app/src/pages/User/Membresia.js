import React, {useState} from 'react';
import {Button, Grid, Paper} from "@material-ui/core";
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import axios from "axios";
import url from "../../config";
import Snackbar from '@material-ui/core/Snackbar';
import { makeStyles } from '@material-ui/core/styles';
import MuiAlert from '@material-ui/lab/Alert';

function Alert(props) {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        '& > * + *': {
            marginTop: theme.spacing(2),
        },
    },
}));

const Membresia = () => {
    const classes = useStyles()
    const paperStyle = {padding: 25, height:'35vh', width:300, margin: "20px auto"}
    const [open, setOpen] = useState(false);
    const [tier, setTier] = useState('')

    const handleChange = (event) => {
        setTier(event.target.value);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleOpen = () => {
        setOpen(true);
    };

    const [openSnack, setOpenSnack] = useState(false);

    const handleClickSnack = () => {
        setOpenSnack(true);
    };

    const handleCloseSnack = (event, reason) => {
        if (reason === 'clickaway') {
            return;
        }
        setOpenSnack(false);
    };


    const onSubmit = async (e) => {
        e.preventDefault()
        const data = {
            tier: tier
        }
        const instance = axios.create({
            withCredentials: true,
        })
        await instance.put(`${url}/user/membresia`, data)
            .then( res => {
                handleClickSnack()
            }).catch(err=>{
                console.log(err)
            })
    }

    return (
        <React.Fragment>
            <form onSubmit={onSubmit} >
                <Paper elevation={10} style={paperStyle}>
                    <Grid align="center">
                        <h2>Membresia</h2>
                    </Grid>
                    <h4>Membresia actual</h4>
                    <Select
                        labelId="demo-controlled-open-select-label"
                        id="demo-controlled-open-select"
                        open={open}
                        onClose={handleClose}
                        onOpen={handleOpen}
                        value={tier}
                        onChange={handleChange}
                        fullWidth
                    >
                        <MenuItem value={"3"}>GOLD</MenuItem>
                        <MenuItem value={"2"}>SILVER</MenuItem>
                        <MenuItem value={"1"}>BRONZE</MenuItem>
                    </Select>
                    <Button color='primary' variant='contained'
                            type='submit'
                        fullWidth
                    >
                        Pagar
                    </Button>

                </Paper>
            </form>
            <div className={classes.root}>
                <Snackbar open={openSnack} autoHideDuration={2500} onClose={handleCloseSnack}>
                    <Alert onClose={handleCloseSnack} severity="success">
                        This is a success message!
                    </Alert>
                </Snackbar>
            </div>
        </React.Fragment>
    );
};

export default Membresia;