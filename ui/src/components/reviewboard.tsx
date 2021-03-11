import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import Collapse from '@material-ui/core/Collapse';
import IconButton from '@material-ui/core/IconButton';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';
import { gql, useQuery, useReactiveVar } from '@apollo/client';
import AppBar from './appbar'

const GETREVIEWS = gql`
  query GetReviews($currentUser: Boolean!) {
    getReviews(input: {
      currentUser: $currentUser,
      }) {
        reviews {
          rating
          review
          timestamp
        }
    }
  }
`;

const useRowStyles = makeStyles({
  root: {
    '& > *': {
      borderBottom: 'unset',
    },
  },
});
    
function Row(props) {
  const { associatedReview } = props;
  const [open, setOpen] = React.useState(false);
  const classes = useRowStyles();
  var reviews = associatedReview.reviews;
  // sort reviews(based on timestamp)

  return (
    <React.Fragment>
      <TableRow className={classes.root}>
        <TableCell>
          <IconButton aria-label="expand row" size="small" onClick={() => setOpen(!open)}>
            {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
          </IconButton>
        </TableCell>
        <TableCell component="th" scope="row">
          {reviews[0].review}
        </TableCell>
      </TableRow>
      <TableRow>
        <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={10}>
          <Collapse in={open} timeout="auto" unmountOnExit>
            <Box margin={1}>
              <Typography variant="h6" gutterBottom component="div">
                Review-History
              </Typography>
              <Table size="small" aria-label="purchases">
                <TableHead>
                  <TableRow>
                    <TableCell>Rating</TableCell>
                    <TableCell align="right">Review</TableCell>
                    <TableCell align="right">Timestamp</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {reviews.map((review) => (
                    <TableRow>
                      <TableCell component="th" scope="row">{review.rating}</TableCell>
                      <TableCell align="right">{review.review}</TableCell>
                      <TableCell align="right">{review.timestamp}</TableCell>
                    </TableRow>
                  ))}
                  <TableCell>{<a href="/offset-review">Offset-Review</a>}</TableCell>
                </TableBody>
              </Table>
            </Box>
          </Collapse>
        </TableCell>
      </TableRow>
    </React.Fragment>
  );
}

Row.propTypes = {
  // row: PropTypes.shape({
  //   review_0: PropTypes.string.isRequired,
  //   reviewhistory: PropTypes.arrayOf(
  //     PropTypes.shape({
  //       rating: PropTypes.string.isRequired,
  //       review: PropTypes.string.isRequired,
  //       timestamp: PropTypes.string.isRequired,
  //     }),
  //   ).isRequired,
  // }).isRequired,
  associatedReview:PropTypes.any,
};

function ReviewBoard(props) {
  const { currentUser } = props;
  const { loading, error, data } = useQuery(GETREVIEWS, {
    variables: { currentUser: currentUser },
  });

  if(loading) {
    return (
      <div>
        <AppBar />
        <div>Loading ...</div>  
      </div>
    )
  }
  console.log(data)

  return (
    <div>
        <AppBar />
         <TableContainer component={Paper}>
         <Table aria-label="collapsible table">
           <TableHead>
             <TableRow>
               <TableCell />
             </TableRow>
           </TableHead>
           <TableBody>
             {/* {formatData(associatedReview).map} */}
             {data.getReviews.map((associatedReview)=>(
               <Row associatedReview={associatedReview} />
             ))}
           </TableBody>
         </Table>
       </TableContainer>
    </div>
  );
}

ReviewBoard.propTypes = {
  // row: PropTypes.shape({
  //   review_0: PropTypes.string.isRequired,
  //   reviewhistory: PropTypes.arrayOf(
  //     PropTypes.shape({
  //       rating: PropTypes.string.isRequired,
  //       review: PropTypes.string.isRequired,
  //       timestamp: PropTypes.string.isRequired,
  //     }),
  //   ).isRequired,
  // }).isRequired,
  currentUser:PropTypes.bool,
};

export default ReviewBoard