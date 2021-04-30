import React from 'react';
import SeasonTable from "../../components/Admin/SeasonTable";
import {Grid} from "@material-ui/core";

const Temporada = () => {

    return (
        <Grid align="center">
            <h1>Temporada</h1>
            <SeasonTable/>
        </Grid>
    );
};

export default Temporada;