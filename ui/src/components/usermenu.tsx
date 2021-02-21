import React from 'react';
import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { useCookies } from "react-cookie";
import { isLoggedInVar } from '../cache'
import { Link } from "react-router-dom";
import { gql, useQuery, useReactiveVar } from '@apollo/client';


const USER = gql`
  query User {
    user {
        firstName
        middleName
        lastName
    }
  }
`;

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        buttonText: {
            color: 'white',
        },
        link: {
          color: 'black',
          textDecoration: 'none',
        }
    }),
);

export default function UserMenu() {
  const [cookies, setCookie, removeCookie] = useCookies(["user"]);
  const classes = useStyles();
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const { loading, error, data } = useQuery(USER);
  // console.log(data)
  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const logout = () => {
    isLoggedInVar(false)
    removeCookie("user")
    handleClose()
    window.location.href = "#"
  }

  return (
    <div>
      {loading? (
              <div>Loading .....</div>
          ) : (
            <div>
              <Button aria-controls="simple-menu" aria-haspopup="true" onClick={handleClick} className={classes.buttonText}>
                {data.user.firstName + " " + data.user.middleName + " " + data.user.lastName}
                {/* Username */}
              </Button>
              <Menu
                id="simple-menu"
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                <MenuItem onClick={handleClose}> 
                  <Link to="/profile" variant="body2" className={classes.link}>
                    Profile
                  </Link>
                </MenuItem>
                <MenuItem onClick={handleClose}> 
                  <Link to="/add-review" variant="body2" className={classes.link}>
                    Add Review
                  </Link>
                </MenuItem>
                <MenuItem>
                  <Link to="/my-reviews" variant="body2" className={classes.link}>
                    My Reviews
                  </Link>
                </MenuItem>
                <MenuItem onClick={logout}>Logout</MenuItem>
              </Menu>
            </div>
          )
      }           
    </div>
  );
}