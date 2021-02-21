import { Link } from "react-router-dom";
import CollapsibleTable from "./reviewboard";

function User(){
    return(
    <div>
        <h1>Welcome to Review Central</h1>
        <Link to="/login" variant="body2">
                Already have an account? Sign in
        </Link>
        <h1>REVIEW-BOARD</h1>
        <CollapsibleTable/>
    </div>
    )
}

export default User