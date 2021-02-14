import React from 'react';
import Button from '@material-ui/core/Button';
import { Link } from "react-router-dom";
import CollapsibleTable from './reviewboard';
import { makeStyles } from '@material-ui/core/styles';
function Myreviews(){
return(
    <div>
<h1>WELCOME TO YOUR REVIEWS</h1>
<br />
<br />
<Link to="/edit_review" variant="body2"><Button>Edit Review</Button>
        </Link>
<br />
<CollapsibleTable/>
</div>

)
}
export default Myreviews