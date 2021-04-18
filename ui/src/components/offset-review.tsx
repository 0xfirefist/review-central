
import React from 'react';
import Button from '@material-ui/core/Button';
import PropTypes from 'prop-types';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import { Link } from "react-router-dom";
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { gql, useMutation } from '@apollo/client';
import AppBar from './appbar';
import Snackbar from '@material-ui/core/Snackbar';
import MuiAlert from '@material-ui/lab/Alert';
import moment from 'moment';

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const ADDOFFSET = gql`
  mutation AddOffset($token: String!,$rating: Float!,$review: String!,$timestamp:String!) {
    addOffset(input: {
      token: $token,
      rating: $rating,
      review: $review,
      timestamp: $timestamp})
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
  },form1: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
    height: '150%',
  },

  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

function OffsetReview({ match }) {
  const [successMes, setSuccessOpen] = React.useState(false);
  const [errorMes, setErrorOpen] = React.useState(false);
  let params = match.params;
  const classes = useStyles();

  const [addOffset] = useMutation(ADDOFFSET, {
    onCompleted(_) {
      setSuccessOpen(true);
    },
    onError(_){
      setErrorOpen(true);
    }
  });

  const handleSubmit = (event: any) => {
    event.preventDefault()
    addOffset({
      variables: {
        token: event.target.token.value,
        rating: +event.target.rating.value,
        review: event.target.review.value,
        timestamp: moment.utc().format()
      }
    })
  }

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setErrorOpen(false);
    setSuccessOpen(false);
  };


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
            id="token"
            label="Token"
            name="token"
            autoComplete="token"
            autoFocus
            value={params.token}
            disabled
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="rating"
            label="Rating"
            type="rating"
            id="rating"
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="review"
            label="Review"
            type="review"
            id="review"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Offset Review
          </Button>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>

      <Snackbar open={successMes} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          Review Offsetted Succesfully!
        </Alert>
      </Snackbar>
      <Snackbar open={errorMes} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          Error
        </Alert>
      </Snackbar>

    </Container>
    </div>
  );
}

export default OffsetReview