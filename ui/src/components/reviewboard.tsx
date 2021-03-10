import React from 'react';
import Button from '@material-ui/core/Button';
import { Link } from "react-router-dom";
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
import { ButtonBaseClassKey } from '@material-ui/core';
const useRowStyles = makeStyles({
  root: {
    '& > *': {
      borderBottom: 'unset',
    },
  },
});

function createData(review_0,rev) {
  return {
    review_0,
    reviewhistory:[...rev],
  };
}
    
function Row(props) {
  const { row } = props;
  const [open, setOpen] = React.useState(false);
  const classes = useRowStyles();

  return (
    <React.Fragment>
      <TableRow className={classes.root}>
        <TableCell>
          <IconButton aria-label="expand row" size="small" onClick={() => setOpen(!open)}>
            {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
          </IconButton>
        </TableCell>
        <TableCell component="th" scope="row">
          {row.review_0}
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
                    <TableCell align="right">review</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {row.reviewhistory.map((historyRow) => (
                    <TableRow key={historyRow.date}>
                      <TableCell component="th" scope="row">{historyRow.rating}</TableCell>
                      <TableCell align="right">{historyRow.review}</TableCell>
                      <TableCell align="right">{historyRow.timestamp}</TableCell>
                    </TableRow>
                  ))}
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
  row: PropTypes.shape({
    review_0: PropTypes.string.isRequired,
    reviewhistory: PropTypes.arrayOf(
      PropTypes.shape({
        rating: PropTypes.string.isRequired,
        review: PropTypes.string.isRequired,
        timestamp: PropTypes.string.isRequired,
      }),
    ).isRequired,
  }).isRequired,
};

const rows = [
  createData('Below are the Reviews',[{rating:'2.0',review:'best',timestamp:'10:30'}, {rating:'3.0',review:'greAT',timestamp:'10:30'}]),
];

export default function CollapsibleTable() {
  return (
    <TableContainer component={Paper}>
      <Table aria-label="collapsible table">
        <TableHead>
          <TableRow>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <Row key={row.review_0} row={row} />
          ))}
        </TableBody>
      </Table>
    </TableContainer>

  );
}