import React, {useEffect, useState} from 'react';
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import Button from "@material-ui/core/Button";
import {makeStyles} from "@material-ui/core/styles";
import axios from "axios";
import url from "../config";
import {TextField} from "@material-ui/core";

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

const EventFullCalendar = () => {
    const classes = useStyles();
    const btnStyle = { margin:'8px 0', }
    const [open, setOpen] = useState(false)
    const [eventos, setEventos] = useState([])
    const [data, setData] = useState({
        id: 0,
        title:'',
    })
    const [prediccion, setPrediccion] = useState({
        idPrediccion: -1,
        local: 0,
        visitante: 0,
    })


    const instance = axios.create({
        withCredentials: true,
    })
    useEffect(()=> {
        (
            async () => {
                await instance.get(`${url}/evento`)
                    .then(res =>{
                        setEventos(res.data)
                        console.log(res.data)
                    }).catch(err =>{

                    })
            }
        )();
    }, [])

    const loadPrediccion = async () => {
        await instance.get(`${url}/prediccion/${data.id}`)
            .then(res =>{
                setPrediccion(res.data)

            }).catch(err =>{

            })
    }

    const handleInputChange = (e:InputEvent) => {
        setPrediccion({
            ...prediccion,
            [e.target.name]:e.target.value
        })
    }

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const prediction = {
        local: prediccion.local,
        visitante: prediccion.visitante,
        idEvento: data.id
    }

    const onSubmitPrediccion = async () => {
        const instance = axios.create({
            withCredentials: true,
            headers: {'Content-Type': 'application/json'},
        })

        console.log(prediction)
        instance.post(`${url}/prediccion`,prediction)
            .then(res => {
                handleClose()
            }).catch(err=>{
                console.log(err)
            })
    }

    return (
        <React.Fragment>
            <FullCalendar
                plugins={[ dayGridPlugin ]}
                initialView="dayGridMonth"
                headerToolbar={{
                    left: 'prev next today',
                    center: 'title',
                    right: 'dayGridMonth,dayGridWeek'
                }}
                events={eventos}
                eventClick={
                    function (arg) {
                        data.title = arg.event.title
                        data.id = arg.event.id
                        data.start = arg.event.start
                        loadPrediccion()
                        handleOpen()
                    }
                }
            />
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
                            <h2 id="transition-modal-title">{data.id}{data.title}</h2>
                            <p id="server-modal-description">Predicción:</p>
                            <form>
                                <TextField type="number"
                                           name="local"
                                           label="Local"
                                           placeholder='Local'
                                           value={prediccion.local}
                                           onChange={handleInputChange}
                                           fullWidth
                                />
                                <TextField type="number"
                                           name="visitante"
                                           label="Visitante"
                                           placeholder='Visitante'
                                           value={prediccion.visitante}
                                           onChange={handleInputChange}
                                           fullWidth
                                />
                                <Button onClick={onSubmitPrediccion} color='primary' variant='contained' fullWidth style={btnStyle}>
                                    Guardar prediccion
                                </Button>
                            </form>
                            <Button onClick={handleClose} variant='contained' fullWidth style={btnStyle}>
                                Cancelar
                            </Button>
                        </div>
                    </Fade>
                </Modal>
            </div>
        </React.Fragment>
    );
};

export default EventFullCalendar;