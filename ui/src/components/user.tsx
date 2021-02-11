import { Link } from "react-router-dom";

function User(){
    return(
    <div>
        <h1>Welcome to Review Central</h1>
        <Link to="/login" variant="body2">
                Already have an account? Sign in
        </Link>
    </div>
    )
}

export default User