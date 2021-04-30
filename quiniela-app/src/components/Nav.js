import React from 'react';
import {Link} from "react-router-dom"
import axios from "axios";
import url from "../config";
import {AppBar, Button, IconButton, makeStyles, Toolbar, Typography} from "@material-ui/core";
import HomeIcon from '@material-ui/icons/Home';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';


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
    const [anchorEl, setAnchorEl] = React.useState(null);
    const open = Boolean(anchorEl);

    const handleClose = () => {
        setAnchorEl(null);
    };
    const handleMenu = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const instance = axios.create({
        withCredentials: true,
    })

    const logout = async () => {
        handleClose()
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
                <IconButton component={Link} to="/" color="inherit" aria-label="menu" className={classes.menuButton}>
                    <HomeIcon/>
                </IconButton>
                <Typography variant='h6' className={classes.title}>
                    QUINIELA
                </Typography>
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
                <IconButton component={Link} to="/admin" color="inherit" aria-label="menu" className={classes.menuButton}>
                    <HomeIcon/>
                </IconButton>
                <Typography variant='h6' className={classes.title}>
                    QUINIELA
                </Typography>
                <Button component={Link} to="/admin/bulkload" variant="text" color="inherit"  >
                    Bulk Load
                </Button>
                <Button component={Link} to="/admin/jornada" variant="text" color="inherit"  >
                    Jornada
                </Button>
                <Button component={Link} to="/admin/temporada" variant="text" color="inherit"  >
                    Temporada
                </Button>
                <Button component={Link} to="/admin/recompensa" variant="text" color="inherit"  >
                    Recompensas
                </Button>
                <Button component={Link} to="/admin/deporte" variant="text" color="inherit"  >
                    Deportes
                </Button>
                <div>
                    <IconButton
                        aria-label="account of current user"
                        aria-controls="menu-appbar"
                        aria-haspopup="true"
                        onClick={handleMenu}
                        color="inherit"
                    >
                        <AccountCircle />
                    </IconButton>
                    <Menu
                        id="menu-appbar"
                        anchorEl={anchorEl}
                        anchorOrigin={{
                            vertical: 'top',
                            horizontal: 'right',
                        }}
                        keepMounted
                        transformOrigin={{
                            vertical: 'top',
                            horizontal: 'right',
                        }}
                        open={open}
                        onClose={handleClose}
                    >
                        <MenuItem onClick={handleClose}>Perfil</MenuItem>
                        <MenuItem component={Link} to="/login" onClick={logout}>Logout</MenuItem>
                    </Menu>
                </div>

            </React.Fragment>
        )
    }

    return (
        <React.Fragment>
            <AppBar>
                <Toolbar>
                    {menu}
                </Toolbar>
            </AppBar>
            <div className={classes.offset}>
            </div>
        </React.Fragment>
    );
};

export default Nav;