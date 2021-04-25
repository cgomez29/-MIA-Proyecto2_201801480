import React, {SyntheticEvent ,useState} from 'react'
import {Redirect} from 'react-router-dom'
import url from '../config'
import axios from 'axios'

const Login = (props: {setName: (name: string) => void }) => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [redirect, setRedirect] = useState(false)

    const onSumbit = async (e: SyntheticEvent) => {
        e.preventDefault()

        const auth = {
            email: email,
            password: password
        }

        const instance = axios.create({
            withCredentials: true,
        })

        await instance.post(`${url}/login`, auth)
            .then(res => {
                //console.log(res)
                setRedirect(true)
                props.setName(res.data.username)
            })
    }


    if(redirect)
        return <Redirect to="/"/>

    return (
        <form onSubmit={onSumbit} >
            <div className="form-inner">
                <h2>Login</h2>
                <div className="form-group">
                    <label htmlFor="email">Email:</label>
                    <input type="text" name="email" onChange={e => setEmail(e.target.value)} />
                </div>
                <div className="form-group">
                    <label htmlFor="password">Password:</label>
                    <input type="password" name="password" onChange={e => setPassword(e.target.value)}/>
                </div>
                <input type="submit" value="LOGIN"/>
            </div>
        </form>
    )
}

export default Login;