import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import {Grid} from "@material-ui/core";


const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        flexWrap: 'wrap',
        '& > *': {
            margin: theme.spacing(1),
        },
    },
}));

const Home = () => {
    const classes = useStyles();

    return (
        <Grid align="center" container>
            <div className={classes.root}>
                <Grid item xs={12} sm={6} md={4}>
                    <Paper elevation={3} >
                        hello
                    </Paper>
                </Grid>
                <Grid item xs={12} sm={6} md={4}>
                    <Paper elevation={3} >
                        hello
                    </Paper>
                </Grid>
            </div>
        </Grid>

    );
};

export default Home;