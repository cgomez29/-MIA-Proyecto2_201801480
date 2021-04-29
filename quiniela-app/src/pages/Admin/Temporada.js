import React from 'react';
import SeasonTable from "../../components/Admin/SeasonTable";
import {Avatar, Grid} from "@material-ui/core";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";

const Temporada = () => {



    return (
        <Grid align="center">
            <h1>Temporada</h1>
            <SeasonTable/>
        </Grid>
    );
};

export default Temporada;