import React, {useContext} from 'react';
import {Link} from "react-router-dom"
import axios from "axios";
import url from "../config";
import {AppBar, Button, IconButton, Toolbar, Typography} from "@material-ui/core";
import HomeIcon from '@material-ui/icons/Home';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import {UserContext} from "../Context/UserContext";
import Badge from '@material-ui/core/Badge';
import Avatar from '@material-ui/core/Avatar';
import { makeStyles, withStyles } from '@material-ui/core/styles';


const StyledBadge = withStyles((theme) => ({
    badge: {
        backgroundColor: '#44b700',
        color: '#44b700',
        boxShadow: `0 0 0 2px ${theme.palette.background.paper}`,
        '&::after': {
            position: 'absolute',
            top: 0,
            left: 0,
            width: '100%',
            height: '100%',
            borderRadius: '50%',
            animation: '$ripple 1.2s infinite ease-in-out',
            border: '1px solid currentColor',
            content: '""',
        },
    },
    '@keyframes ripple': {
        '0%': {
            transform: 'scale(.8)',
            opacity: 1,
        },
        '100%': {
            transform: 'scale(2.4)',
            opacity: 0,
        },
    },
}))(Badge);

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
            margin: theme.spacing(0),
        },
    },
}))

const Nav = () => {
    const {rol, setRol, src, setSrc} = useContext(UserContext)

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
        setRol('')
        setSrc('')
    }

    let menu;

    if (rol===2) {
        console.log(src)

        menu = (
            <React.Fragment>
                <IconButton component={Link} to="/user" color="inherit" aria-label="menu" className={classes.menuButton}>
                    <HomeIcon/>
                </IconButton>
                <Typography variant='h5' className={classes.title}>
                    QUINIELA
                </Typography>
                <div>
                    <IconButton
                        aria-label="account of current user"
                        aria-controls="menu-appbar"
                        aria-haspopup="true"
                        onClick={handleMenu}
                        color="inherit"
                    >
                        <div className={classes.root}>
                            <StyledBadge
                                overlap="circle"
                                anchorOrigin={{
                                    vertical: 'bottom',
                                    horizontal: 'right',
                                }}
                                variant="dot"
                            >
                                <Avatar
                                    style={{ width: 40, height: 40,}}
                                    alt=""
                                    src={src}
                                />
                            </StyledBadge>
                        </div>

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
    } else if(rol===1) {
        menu = (
            <React.Fragment>
                <IconButton component={Link} to="/admin" color="inherit" aria-label="menu" className={classes.menuButton}>
                    <HomeIcon/>
                </IconButton>
                <Typography variant='h6' className={classes.title}>
                    QUINIELA
                </Typography>
                <Button component={Link} to="/admin/temporada" variant="text" color="inherit"  >
                    Temporada
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
                        <div className={classes.root}>
                            <StyledBadge
                                overlap="circle"
                                anchorOrigin={{
                                    vertical: 'bottom',
                                    horizontal: 'right',
                                }}
                                variant="dot"
                            >
                                <Avatar
                                    style={{ width: 40, height: 40,}}
                                    alt=""
                                    src={src}
                                />
                            </StyledBadge>
                        </div>

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
    } else {
        menu = (
            <React.Fragment>
                <IconButton component={Link} to="/" color="inherit" aria-label="menu" className={classes.menuButton}>
                    <HomeIcon/>
                </IconButton>
                <Typography variant='h6' className={classes.title}>
                    QUINIELA
                </Typography>
                <Button component={Link} to="/login" variant="text" color="inherit">
                    LogIn
                </Button>
                <Button component={Link} to="/register" variant="text" color="inherit">
                    SignIn
                </Button>
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