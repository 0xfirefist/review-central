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

function createData(productname,category) {
  return {
    productname,
    category,
    history: [
      { date: '2020-01-05', email_ID: 'nitesh212@gmail.com',review: 'The Phone was the Worst',price:20000,edit:'/edit_review'},
      { date: '2020-01-02', email_ID: 'kalradev@gmail.com',review: 'The Phone is Great',price:45000,edit:'Edit'},
    ],
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
          {row.productname}
        </TableCell>
        <TableCell align="right">{row.category}</TableCell>
      </TableRow>
      <TableRow>
        <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={5}>
          <Collapse in={open} timeout="auto" unmountOnExit>
            <Box margin={1}>
              <Typography variant="h6" gutterBottom component="div">
                History
              </Typography>
              <Table size="small" aria-label="purchases">
                <TableHead>
                  <TableRow>
                    <TableCell>Date</TableCell>
                    <TableCell>email_ID</TableCell>
                    <TableCell align="right">review</TableCell>
                    <TableCell align="right">price</TableCell>
                    <TableCell align="right"> Edit</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {row.history.map((historyRow) => (
                    <TableRow key={historyRow.date}>
                      <TableCell component="th" scope="row">{historyRow.date}</TableCell>
                      <TableCell>{historyRow.email_ID}</TableCell>
                      <TableCell align="right">{historyRow.review}</TableCell>
                      <TableCell align="right">{historyRow.price}</TableCell>
                      <TableCell align="right">{<Link to="/offset-review">Offset_Review</Link>}</TableCell>
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
    productname: PropTypes.string.isRequired,
    category: PropTypes.string.isRequired,
    history: PropTypes.arrayOf(
      PropTypes.shape({
        date: PropTypes.string.isRequired,
        email_ID: PropTypes.string.isRequired,
        review: PropTypes.string.isRequired,
        price: PropTypes.number.isRequired,
        edit: PropTypes.string.isRequired,
      }),
    ).isRequired,
  }).isRequired,
};

const rows = [
  createData('Nokia 7 Plus','Phones'),
  createData('One plus 8','Phones'),
];

export default function CollapsibleTable() {
  return (
    <TableContainer component={Paper}>
      <Table aria-label="collapsible table">
        <TableHead>
          <TableRow>
            <TableCell />
            <TableCell>Products</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <Row key={row.productname} row={row} />
          ))}
        </TableBody>
      </Table>
    </TableContainer>

  );
}
