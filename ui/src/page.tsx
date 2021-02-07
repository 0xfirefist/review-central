import Login from './components/login'
import Register from './components/register'
import { useState } from 'react'

function Page() {
    const [login, setLogin] = useState('login')

    const triggerLoginState = () => {
        if (login === 'login')
            setLogin('register')
        else
            setLogin('login')
    }

    return (
        <div>
            { login === 'login' && (
                <Login toggle={triggerLoginState} />
            )}

            { login === 'register' && (
                <Register toggle={triggerLoginState} />
            )}
        </div>
    )
}

export default Page