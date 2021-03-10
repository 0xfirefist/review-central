
import React from 'react';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import { Link } from "react-router-dom";
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { gql, useMutation } from '@apollo/client';
import { useCookies } from "react-cookie";
import AppBar from './appbar';

const LOGIN = gql`
  mutation Login($email: String!, $password: String!) {
      login(input: {
        email: $email,
        password: $password
      })
    }
`;

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Review Central
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

function Login() {
  const classes = useStyles();
  const [_, setCookie] = useCookies(["user"]);

  const [login] = useMutation(LOGIN, {
    onCompleted(data) {
      // localStorage.setItem("token",data.createUser)
      setCookie("user", data.login, {      
        path: "/",
        // secure: true
        sameSite: 'strict'
      });
    }
  });

  const handleSubmit = (event: any) => {
    event.preventDefault()
    login({
      variables: {
        email: event.target.email.value,
        password: event.target.password.value
      }
    })
  }

  return (
    <div>
      <AppBar/>
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Typography component="h1" variant="h5">
          Offset Review
        </Typography>
        <form className={classes.form} noValidate onSubmit={handleSubmit}>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="review_id"
            label="Review_ID"
            id="review_id"
            autoComplete="review_id"
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="productname"
            label="Product Name"
            name="productname"
            autoComplete="productname"
            autoFocus
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Add Offset Review
          </Button>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
    </div>
  );
}

export default Login