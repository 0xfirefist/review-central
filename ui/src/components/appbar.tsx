import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import UserMenu from './usermenu'
import Button from '@material-ui/core/Button';
import { isLoggedInVar } from '../cache'
import { Link } from "react-router-dom";
import { useReactiveVar } from '@apollo/client';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1,
        },
        menuButton: {
            marginRight: theme.spacing(2),
        },
        title: {
            flexGrow: 1,
        },
        link: {
            color: 'white',
            textDecoration: 'none',
        },
    }),
);

export default function ButtonAppBar() {
    const classes = useStyles();
    const isLoggedIn = useReactiveVar(isLoggedInVar);

    return (
        <div className={classes.root}>
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" className={classes.title}>
                        <Link to="/" variant="body2" className={classes.link}>
                            Review Central
                        </Link>
                    </Typography>
                    {isLoggedIn? (
                            <UserMenu />
                        ) : (
                            <Button color="inherit">
                                <Link to="/login" variant="body2" className={classes.link}>
                                    Login/Register
                                </Link>
                            </Button> 
                        )
                    }                    
                </Toolbar>
            </AppBar>
        </div>
    );
}