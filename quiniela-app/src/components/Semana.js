import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        height: 140,
        width: 100,
    },
    control: {
        padding: theme.spacing(2),
    },
}));

const Semana = (props) => {
    const [spacing, setSpacing] = React.useState(0);
    const classes = useStyles();

    return (
        <Grid container className={classes.root} spacing={2}>
            <Grid item xs={12}>
                <Grid container justify="center" spacing={spacing}>
                    <Paper className={classes.paper}>
                        <Typography gutterBottom variant="subtitle1">
                            Standard license
                        </Typography>
                    </Paper>
                </Grid>
            </Grid>

        </Grid>
    );
};

export default Semana;