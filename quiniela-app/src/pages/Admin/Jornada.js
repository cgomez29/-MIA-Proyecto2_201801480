import React from 'react';
import {Button, Grid} from "@material-ui/core";
import JornadaFullCalendar from "../../components/Admin/JornadaFullCalendar";

const Jornada = () => {
    return (
        <React.Fragment>
            <Grid container>
                <Button variant="contained" color="primary" style={{margin: 5}}>Agregar Evento</Button>
            </Grid>
            <JornadaFullCalendar/>
        </React.Fragment>
    );
};

export default Jornada;