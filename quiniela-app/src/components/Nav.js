import React from 'react';
import {Link} from "react-router-dom"
import axios from "axios";
import url from "../config";
import {AppBar, Button, IconButton, makeStyles, Toolbar, Typography} from "@material-ui/core";
import HomeIcon from '@material-ui/icons/Home';
import Avatar from '@material-ui/core/Avatar';

const  useStyles = makeStyles(theme => ({
    offset: theme.mixins.toolbar,
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title:{
        flexGrow: 1
    },
    root: {
        display: 'flex',
        '& > *': {
            margin: theme.spacing(1),
        },
    },
}))


const Nav = (props: {name: string, setName: (name: string) => void }) => {
    const classes = useStyles()

    const instance = axios.create({
        withCredentials: true,
    })

    const logout = async () => {
        await instance.post(`${url}/logout`, {})
            .then(res => {
                //console.log(res)
            })
        props.setName('')
    }

    let menu;

    if(props.name===''){
        menu = (
            <React.Fragment>
                <Button component={Link} to="/login" variant="text" color="inherit"  >
                    LogIn
                </Button>
                <Button component={Link} to="/register" variant="text" color="inherit"  >
                    SignIn
                </Button>
            </React.Fragment>
        )
    } else {
        menu = (
            <React.Fragment>
                <Button component={Link} to="/login" variant="text" color="inherit"  >
                    Bulk Load
                </Button>
                <Button component={Link} to="/login" variant="text" color="inherit"  >
                    Jornada
                </Button>
                <Button component={Link} to="/temporada" variant="text" color="inherit"  >
                    Temporada
                </Button>
                <Button component={Link} to="/login" variant="text" color="inherit"  >
                    Recompensas
                </Button>
                <Button component={Link} to="/login" onClick={logout} variant="text" color="inherit"  >
                    LogOut
                </Button>
                <div className={classes.root}>
                    <Avatar alt="Remy Sharp" src="/static/images/avatar/1.jpg" />
                </div>
            </React.Fragment>
        )
    }

    return (
        <React.Fragment>
            <AppBar>
                <Toolbar>
                    <IconButton component={Link} to="/" color="inherit" aria-label="menu" className={classes.menuButton}>
                        <HomeIcon/>
                    </IconButton>
                    <Typography variant='h6' className={classes.title}>
                        QUINIELA
                    </Typography>
                    {menu}
                </Toolbar>
            </AppBar>
            <div className={classes.offset}></div>
        </React.Fragment>
    );
};

export default Nav;