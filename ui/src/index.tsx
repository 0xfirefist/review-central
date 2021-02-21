import React from 'react';
import ReactDOM from 'react-dom';
import HomePage from './homepage';
import { CookiesProvider } from "react-cookie";
import Login from './components/login'
import Register from './components/register'
import AddReview from './components/add-review'
import EditReview from './components/edit_review'
import Profile from './components/profile'
import Myreviews1 from './components/myreviews';
import {HashRouter, Route, Switch} from 'react-router-dom'
import { ApolloProvider, ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client';

const link = createHttpLink({
  uri: '/graphql',
  credentials: 'same-origin'
});

const client = new ApolloClient({
  link,
  cache: new InMemoryCache(),
});

ReactDOM.render(
  <CookiesProvider>    
        <ApolloProvider client={client}>
            <HashRouter>
                <Switch>
                    <Route exact path='/' component={HomePage}/>
                    <Route exact path='/edit-review' component={EditReview}/>
                    <Route exact path='/login' component={Login}/>
                    <Route exact path='/register' component={Register}/>
                    <Route exact path='/add-review' component={AddReview}/>
                    <Route exact path='/profile' component={Profile}/>
                    <Route exact path='/myreviews' component={Myreviews1}/>
                </Switch>
            </HashRouter>
        </ApolloProvider>
  </CookiesProvider>,
  document.getElementById('root')
);
