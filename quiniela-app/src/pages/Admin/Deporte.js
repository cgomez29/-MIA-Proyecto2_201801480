import React, {useEffect, useState} from 'react';
import CardG from "../../components/CardG";
import {Grid, TextField} from "@material-ui/core";
import axios from "axios";
import url from "../../config";
import {Redirect} from "react-router-dom";
import Button from "@material-ui/core/Button";
import {makeStyles} from "@material-ui/core/styles";
import AddOutlinedIcon from '@material-ui/icons/AddOutlined';
import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import EditIcon from "@material-ui/icons/Edit";

const useStyles = makeStyles((theme) => ({
    root: {
        maxWidth: 345,
        [theme.breakpoints.down("sm")] : {
            maxWidth: 200
        },
        minWidth: 100,
        margin: 10,
    },
    media: {
        height: 150,
    },
    modal: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
    paper: {
        backgroundColor: theme.palette.background.paper,
        border: '2px solid #000',
        boxShadow: theme.shadows[4],
        padding: theme.spacing(2, 4, 3),
    },
}));

const Deporte = () => {
    const classes = useStyles();
    const btnStyle = { margin:'8px 0', }
    const [loggedIn, setLoggedIn] = useState(false)
    const [open, setOpen] = React.useState(false);
    const [data, setData] = useState([])
    const [file, setFile] = useState([])
    const [deporte, setDeporte] = useState({
        name: '',
        color: '',
    })
    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const instance = axios.create({
        withCredentials: true,
        headers: { 'Content-Type': 'multipart/form-data' },
    })

    useEffect(()=>{
        (
            async () => {
                await instance.get( `${url}/deporte`)
                    .then(res => {
                        setData(res.data)
                        console.log(res.data)
                    }).catch( err =>{
                        setLoggedIn(true)
                    })
            }
        )();

    },[])

    const onSubmit = async (e) => {
        e.preventDefault()
        const formaData = new FormData()
        formaData.append("nombre", deporte.name)
        formaData.append("color", deporte.color)
        formaData.append("file", file)

        await instance.post(`${url}/deporte`, formaData)
            .then(res => {
                handleClose()
            }).catch(err => {
                console.log(err)
            })
    }

    const handleInputChange = (e) => {
        setDeporte({
            ...deporte,
            [e.target.name]:e.target.value
        })
    }

    if (loggedIn)
        return <Redirect to="/login"/>

    return (
        <React.Fragment>
            <Grid container>
                <Grid>
                    <Button variant="contained" color="primary"
                            startIcon={<AddOutlinedIcon/>}
                            onClick={handleOpen}
                            className={{ display: 'flex', }}
                    >
                        Agregar deporte
                    </Button>
                </Grid>
                <Grid align="center" container>

                    {
                        data.map((deporte) => (
                            <Grid item xs={12} sm={6} md={4}>
                                <CardG

                                    key={deporte.id}
                                    {...deporte}
                                />
                            </Grid>
                        ))
                    }
                </Grid>
            </Grid>

            <div>
                <Modal
                    aria-labelledby="transition-modal-title"
                    aria-describedby="transition-modal-description"
                    className={classes.modal}
                    open={open}
                    onClose={handleClose}
                    closeAfterTransition
                    BackdropComponent={Backdrop}
                    BackdropProps={{
                        timeout: 500,
                    }}
                >
                    <Fade in={open}>
                        <div className={classes.paper}>
                            <h2 id="transition-modal-title"><EditIcon/>Agregar</h2>
                            <form onSubmit={onSubmit}>
                                <TextField label='name' name='name' value={deporte.name} onChange={handleInputChange} placeholder="Enter name" fullWidth required />
                                <TextField type='color'
                                           name='color'
                                           defaultValue="#000"
                                           value={deporte.color}
                                           fullWidth
                                           onChange={handleInputChange}
                                />
                                <Button variant="contained" component="label" fullWidth style={btnStyle}>
                                    Upload photo
                                    <input type="file"  hidden onChange={e => setFile(e.target.files[0])} />
                                </Button>
                                <Button type="submit" color='primary' variant='contained' fullWidth>
                                    Agregar
                                </Button>
                            </form>
                            <Button variant='contained' fullWidth style={btnStyle}>
                                Cancelar
                            </Button>
                        </div>
                    </Fade>
                </Modal>
            </div>
        </React.Fragment>
    );
};

export default Deporte;