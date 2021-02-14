import React from 'react';
import ReactDOM from 'react-dom';
import HomePage from './homepage';
import { CookiesProvider } from "react-cookie";
import Login from './components/login'
import Register from './components/register'
import AddReview from './components/add-review'
import Profile from './components/profile'
import {HashRouter, Route, Switch} from 'react-router-dom'
import { ApolloProvider, ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: window.location.origin + "/graphql",
  cache: new InMemoryCache(),
  headers: {    authorization: localStorage.getItem('token') || '',  }
});

ReactDOM.render(
  <CookiesProvider>    
        <ApolloProvider client={client}>
            <HashRouter>
                <Switch>
                    <Route exact path='/' component={HomePage}/>
                    <Route exact path='/login' component={Login}/>
                    <Route exact path='/register' component={Register}/>
                    <Route exact path='/add-review' component={AddReview}/>
                    <Route exact path='/profile' component={Profile}/>
                </Switch>
            </HashRouter>
        </ApolloProvider>
  </CookiesProvider>,
  document.getElementById('root')
);
