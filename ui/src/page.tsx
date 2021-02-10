import Login from './components/login'
import Register from './components/register'
import {BrowserRouter as Router,Switch,Route} from 'react-router-dom'
import { useState } from 'react'
import { ApolloProvider, ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: 'http://localhost:8000/graphql',
  cache: new InMemoryCache(),
  headers: {    authorization: localStorage.getItem('token') || '',  }
});


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