import React,{useState} from 'react';
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import EditIcon from "@material-ui/icons/Edit";
import Button from "@material-ui/core/Button";
import {makeStyles} from "@material-ui/core/styles";

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
    const [open, setOpen] = useState(false)
    const [data, setData] = useState({
        title:'',
    })
    const btnStyle = { margin:'8px 0', }

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    return (
        <React.Fragment>
            <FullCalendar
                plugins={[ dayGridPlugin ]}
                initialView="dayGridMonth"
                headerToolbar={{
                    left: 'prev, next today',
                    center: 'title',
                    right: 'dayGridMonth,dayGridWeek'
                }}
                events={[{"title":"Cristian", "start":"2021-05-10 10:00", "end":"2021-05-11 10:00", "color":"red"},
                    {"title":"Gomez", "start":"2021-05-10 12:00", "end":"2021-05-10 14:00", "color":"#000"}]}
                eventClick={
                    function (arg) {
                        data.title = arg.event.title
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
                            <h2 id="transition-modal-title">{data.title}</h2>

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

export default EventFullCalendar;