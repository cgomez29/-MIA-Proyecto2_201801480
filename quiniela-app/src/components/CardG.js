import React, {useEffect, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import noImg from '../images/noimage.png';
import Modal from '@material-ui/core/Modal';
import Backdrop from '@material-ui/core/Backdrop';
import Fade from '@material-ui/core/Fade';
import EditIcon from '@material-ui/icons/Edit';
import axios from "axios";
import url from "../config";
import {TextField} from "@material-ui/core";
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import useMediaQuery from '@material-ui/core/useMediaQuery';
import { useTheme } from '@material-ui/core/styles';
import DeleteIcon from '@material-ui/icons/Delete';

const useStyles = makeStyles((theme) => ({
    root: {
        maxWidth: 345,
        [theme.breakpoints.down("md")] : {
            maxWidth: 200
        },
        minWidth: 245,
        margin: 10,
    },
    media: {
        height: 140,
    },
    modal: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
    paper: {
        backgroundColor: theme.palette.background.paper,
        border: '2px solid #000',
        boxShadow: theme.shadows[5],
        padding: theme.spacing(2, 4, 3),
    },
}));

const CardG = ({idDeporte,name, file}) => {
    const [file2, setFile] = useState(null)
    const theme = useTheme();
    const fullScreen = useMediaQuery(theme.breakpoints.down('sm'));
    const btnStyle = { margin:'8px 0', }
    const classes = useStyles();
    const [open, setOpen] = React.useState(false);
    const [openDialog, setOpenDialog] = React.useState(false);
    const [data, setData] = useState({
        color: '',
        file: '',
    })

    const handleClickOpenDialog = () => {
        setOpenDialog(true);
    };

    const handleCloseDialog = () => {
        setOpenDialog(false);
    };

    const loadDeporte = async () => {
        const instance = axios.create({
            withCredentials: true,
            headers: { 'Content-Type': 'multipart/form-data' },
        })
        console.log(`${url}/deporte/${idDeporte}`)
        await instance.get(`${url}/deporte/${idDeporte}`).then(res => {
            setData(res.data)
        })
    }

    useEffect(() => {
        loadDeporte();
    }, [])

    const handleInputChange = (e:InputEvent) => {
        setData({
            ...data,
            [e.target.name]:e.target.value
        })
    }

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

    const onSubmit = async (e) => {
        e.preventDefault()
        const formData = new FormData()
        formData.append("color", data.color)
        if (file != null) {
            formData.append('file', file2)
        } else {
            formData.append('file', data.file)
        }

        await instance.put(`${url}/deporte/${idDeporte}`, formData)
            .then(res => {
                handleClose()
            }).catch(err =>{
                console.log(err)
            })

    }

    const onSubmitDelete = async (e) => {
        e.preventDefault()
        await instance.delete(`${url}/deporte/${idDeporte}`)
            .then(res => {
                handleCloseDialog()
            }).catch(err =>{
                console.log(err)
            })

    }

    return (
        <React.Fragment>
            <Card className={classes.root}>
                <CardActionArea>
                    <CardMedia
                        className={classes.media}
                        image={`${url}/img/${file}` || noImg}
                        title="Contemplative Reptile"
                    />
                    <CardContent>
                        <Typography gutterBottom variant="h5" component="h2">
                            {name}
                        </Typography>
                        <Typography variant="body2" color="textSecondary" component="p">
                            Lizards are a widespread group of squamate reptiles, with over 6,000 species, ranging
                            across all continents except Antarctica
                        </Typography>
                    </CardContent>
                </CardActionArea>
                <CardActions>
                    <Button size="small" color="primary" onClick={handleOpen}>
                        Editar
                    </Button>
                    <Button size="small" color="primary"
                        onClick={handleClickOpenDialog}
                    >
                        Eliminar
                    </Button>
                </CardActions>
            </Card>
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

                            <h2 id="transition-modal-title"><EditIcon/>Editar: {name}</h2>
                            <p id="transition-modal-description">
                                <form onSubmit={onSubmit}>
                                    <TextField type='color'
                                               name='color'
                                               defaultValue="#000"
                                               value={data.color}
                                               fullWidth
                                               onChange={handleInputChange}
                                    />
                                    <TextField disabled name='file' value={data.file} fullWidth onChange={handleInputChange}/>
                                    <Button variant="contained" component="label" fullWidth style={btnStyle}>
                                        Upload photo
                                        <input type="file"  hidden onChange={e => setFile(e.target.files[0])} />
                                    </Button>
                                    <Button type="submit"  color='primary' variant='contained' fullWidth style={btnStyle}
                                    >Guardar</Button>
                                </form>
                                <Button variant='contained' onClick={handleClose} fullWidth style={btnStyle}
                                >Cancelar</Button>
                            </p>
                        </div>
                    </Fade>
                </Modal>
            </div>
            <div>
                <Dialog
                    fullScreen={fullScreen}
                    open={openDialog}
                    onClose={handleCloseDialog}
                    aria-labelledby="responsive-dialog-title"
                >
                    <DialogTitle id="responsive-dialog-title">
                        <DeleteIcon/>
                        Esta seguro que desea eliminar el deporte {name}?
                    </DialogTitle>
                    <DialogContent>
                        <DialogContentText>
                            Este proceso no es reversible.
                        </DialogContentText>
                    </DialogContent>
                    <DialogActions>
                        <Button autoFocus onClick={handleCloseDialog} color="primary">
                            Cancelar
                        </Button>
                        <form onSubmit={onSubmitDelete}>
                            <Button type='submit' color="primary">
                                Aceptar
                            </Button>
                        </form>
                    </DialogActions>
                </Dialog>
            </div>
        </React.Fragment>
    )
};

export default CardG;