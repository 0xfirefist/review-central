import Login from './components/login'
import Register from './components/register'
import { useState } from 'react'
import { ApolloProvider, ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: 'http://localhost:8000/graphql',
  cache: new InMemoryCache(),
  headers: {    authorization: localStorage.getItem('token') || '',  }
});


function Page() {
    const [login, setLogin] = useState('login')

    const triggerLoginState = () => {
        if (login === 'login')
            setLogin('register')
        else
            setLogin('login')
    }

    return (
        <ApolloProvider client={client}>
            <div>
                { login === 'login' && (
                    <Login toggle={triggerLoginState} />
                )}

                { login === 'register' && (
                    <Register toggle={triggerLoginState} />
                )}
            </div>
        </ApolloProvider>
    )
}

export default Page