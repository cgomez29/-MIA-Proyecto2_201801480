import React, {useEffect, useState} from 'react';
import {Button, Grid, TextField} from "@material-ui/core";
import JornadaFullCalendar from "../../components/Admin/JornadaFullCalendar";
import axios from "axios";
import url from "../../config";
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import PeopleAltIcon from '@material-ui/icons/PeopleAlt';
import EventAvailableIcon from '@material-ui/icons/EventAvailable';
import EventBusyIcon from '@material-ui/icons/EventBusy';
import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import EditIcon from "@material-ui/icons/Edit";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary,
    },
    rootModal: {
        maxWidth: 345,
        [theme.breakpoints.down("sm")] : {
            maxWidth: 200
        },
        minWidth: 100,
        margin: 10,
    },
    mediaModal: {
        height: 150,
    },
    modal: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
    paperModal: {
        backgroundColor: theme.palette.background.paper,
        border: '2px solid #000',
        boxShadow: theme.shadows[4],
        padding: theme.spacing(2, 4, 3),
    },
}));

const Jornada = () => {
    const classes = useStyles();
    const btnStyle = { margin:'8px 0', }
    const [open, setOpen] = useState(false)
    const [openSelect, setOpenSelect] = useState(false)
    const [deporte, setDeporte] = useState([])
    const [datos, setDatos] = useState({
        temporada: '',
        participantes: '',
        fecha: '',
        jornada: '',
        idJornada: '',
    })
    const [evento, setEvento] = useState({
        fecha: '',
        hora: '',
        local: '',
        visitante: '',
        idDeporte: '',
    })


    const instance = axios.create({
        withCredentials: true,
    })

    const loadDeportes = async () => {
        await instance.get(`${url}/deporte`)
            .then(res => {
                setDeporte(res.data)
                console.log(res.data)
            })
    }

    useEffect(()=>{
        (
            async () => {
                await instance.get(`${url}/jornada/detalle`)
                    .then(res => {
                        setDatos(res.data)
                    })
            }
        )();
        loadDeportes()
    },[])

    const onSubmit = async (e) => {
        e.preventDefault()
        console.log(evento)
        await instance.post(`${url}/evento`, evento)
            .then(res => {
                handleClose()
            }).catch(err => {
                console.log(err)
            })
    }

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleOpenSelect = () => {
        setOpenSelect(true);
    };

    const handleCloseSelect = () => {
        setOpenSelect(false);
    };

    const handleInputChange = (e) => {
        setEvento({
            ...evento,
            [e.target.name]:e.target.value
        })
    }

    return (
        <React.Fragment>
            <Grid container spacing={3}>
                <Grid item xs={6} sm={3} direction="column" spacing={2}>
                    <Paper className={classes.paper}>
                        <EventAvailableIcon/>
                        <Typography gutterBottom variant="subtitle1">
                            Jornada actual:
                        </Typography>
                        <Typography variant="subtitle1">{datos.jornada}</Typography>
                    </Paper>
                </Grid>
                <Grid item xs={6} sm={3} direction="column" spacing={2}>
                    <Paper className={classes.paper}>
                        <EventAvailableIcon/>
                        <Typography gutterBottom variant="subtitle1">
                            Temporada:
                        </Typography>
                        <Typography variant="subtitle1">{datos.temporada}</Typography>
                    </Paper>
                </Grid>
                <Grid item xs={6} sm={3} direction="column" spacing={2}>
                    <Paper className={classes.paper}>
                        <PeopleAltIcon/>
                        <Typography gutterBottom variant="subtitle1">
                            Participantes:
                        </Typography>
                        <Typography variant="subtitle1">{datos.participantes}</Typography>
                    </Paper>
                </Grid>
                <Grid item xs={6} sm={3} direction="column" spacing={2}>
                    <Paper className={classes.paper}>
                        <EventBusyIcon/>
                        <Typography gutterBottom variant="subtitle1">
                            Finaliza:
                        </Typography>
                        <Typography variant="subtitle1">{datos.fecha}</Typography>
                    </Paper>
                </Grid>
            </Grid>
            <Grid container>
                <Button variant="contained"
                        color="primary"
                        style={{margin: 5}}
                        onClick={handleOpen}
                >
                    Agregar Evento
                </Button>
            </Grid>
            <JornadaFullCalendar/>
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
                        <div className={classes.paperModal}>
                            <h2 id="transition-modal-title"><EditIcon/>Agregar</h2>
                            <form onSubmit={onSubmit}>
                                <Select
                                    label="Deporte"
                                    labelId="demo-controlled-open-select-label"
                                    id="demo-controlled-open-select"
                                    open={openSelect}
                                    name="idDeporte"
                                    onClose={handleCloseSelect}
                                    onOpen={handleOpenSelect}
                                    value={evento.idDeporte}
                                    onChange={handleInputChange}
                                    fullWidth
                                >
                                    {
                                        deporte.map((value) => (
                                            <MenuItem key={value.idDeporte} value={value.idDeporte}>{value.name}</MenuItem>
                                        ))
                                    }

                                </Select>
                                <TextField label='Local'
                                           name='local'
                                           onChange={handleInputChange}
                                           value={evento.local}
                                           onChange={handleInputChange}
                                           placeholder="Enter name"
                                           fullWidth required
                                />
                                <TextField label='Visitante'
                                           name='visitante'
                                           onChange={handleInputChange}
                                           value={evento.visitante}
                                           onChange={handleInputChange}
                                           placeholder="Enter name"
                                           fullWidth required
                                />
                                <TextField
                                    id="date"
                                    label="Fecha"
                                    type="date"
                                    InputLabelProps={{
                                        shrink: true,
                                    }}
                                    name='fecha'
                                    value={evento.fecha}
                                    onChange={handleInputChange}
                                    fullWidth
                                    required
                                />
                                <TextField
                                    id="time"
                                    label="Hora"
                                    type="time"
                                    InputLabelProps={{
                                        shrink: true,
                                    }}
                                    name='hora'
                                    value={evento.hora}
                                    onChange={handleInputChange}
                                    fullWidth
                                    required
                                />
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

export default Jornada;