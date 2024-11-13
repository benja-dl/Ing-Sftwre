import { useState } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";

export default function LoginPage() {

    const navigate = useNavigate();
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    async function handleLogin(e){
        e.preventDefault()
        try {
            const requestBody = {email, password}
            const response = await axios.post('http://localhost:8080/api/authentication/signin', requestBody)
            console.log(response)
            localStorage.setItem('access_token', response.data.token)
            console.log(response.data.token)
            navigate('/')
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <div className="container" style={{marginTop:"10vh"}}>
            <form onSubmit={handleLogin}>
                <h2>Login to your account</h2>
                <p>Welcome back!</p>
                <div className="mb-3">
                    <label htmlFor="email" className="form-label">Email address :</label>
                    <input onChange={e => {setEmail(e.target.value)}} type="email" className="form-control" id="email"/>
                </div>
                <div className="mb-3">
                    <label htmlFor="password" className="form-label">Password :</label>
                    <input onChange={e => {setPassword(e.target.value)}} type="password" className="form-control" id="password"/>
                </div>
                <button type="submit" className="btn btn-primary">LOG IN</button>
                <p><br />Demo User: <br />Email: user@example.com <br />Password: password12345</p>
            </form>
        </div>
    )

}