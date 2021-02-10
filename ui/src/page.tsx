import Login from './components/login'
import Register from './components/register'
import {BrowserRouter as Router,Switch,Route} from 'react-router-dom'
import { useState } from 'react'

function Page() {
    const [login, setLogin] = useState('login')
    

    return (
        <Router>
        <div>
        <Switch>
            <Route path='/login'>
                <Login />
            </Route>
            <Route path='/register'>
                <Register />
            </Route>
            <Route path='/'>
                <Login />
            </Route>
        </Switch>
        </div>
        </Router>
    )
}

export default Page