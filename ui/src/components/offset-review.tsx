
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
import AppBar from './appbar';
import moment from 'moment';

const OFFSETREVIEW = gql`
  mutation OFFSETReview($token: String!,$rating: Float!,$review: String!,$timestamp:String!) {
    offsetReview(input: {
      token: $token,
      rating: $rating,
      review: $review,S
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

function Login() {
  const classes = useStyles();

  const [offsetReview] = useMutation(OFFSETREVIEW, {
    onCompleted(data) {
      console.log(data.offsetReview)
    }
  });

  const handleSubmit = (event: any) => {
    event.preventDefault()
    offsetReview({
      variables: {
        token: event.target.token.value,
        rating: +event.target.rating.value,
        review: event.target.review.value,
        timestamp: moment.utc().format()
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
            id="token"
            label="Token"
            name="token"
            autoComplete="token"
            autoFocus
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
    </Container>
    </div>
  );
}

export default Login