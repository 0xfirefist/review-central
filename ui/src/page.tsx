import Login from './components/login'
import Register from './components/register'
import User from './components/user'
import {HashRouter, Route, Switch} from 'react-router-dom'
import { ApolloProvider, ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: 'http://localhost:8000/graphql',
  cache: new InMemoryCache(),
  headers: {    authorization: localStorage.getItem('token') || '',  }
});


function Page() {
    return (
        <ApolloProvider client={client}>
            <HashRouter>
                <Switch>
                    <Route exact path='/' component={User}/>
                    <Route exact path='/login' component={Login}/>
                    <Route exact path='/register' component={Register}/>
                </Switch>
            </HashRouter>
        </ApolloProvider>
    )
}

export default Page