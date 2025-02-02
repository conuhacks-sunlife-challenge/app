import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './index.css';
import User from '../../types/user';
import endpoints from '../../endpoints';
import Credentials from '../../types/credentials';
import { useGlobalState } from '../../GlobalState';

const Login: React.FC = () => {

  const navigate = useNavigate();
  
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [incorrect, setIncorrect] = useState<boolean>(false);
  const {setCredentials, setLoggedIn} = useGlobalState();

  const authenticate = async (credentials: Credentials) => {
    const body = JSON.stringify(credentials)
    const req = new Request(endpoints.Auth, {

      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: body,
    })
    const res = await fetch(req)
    return res.status
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();


    console.log('Form submitted:', {email, password });
    const credentials = {
      Email: email,
      Password: password,
    }
    const authStatus = await authenticate(credentials)

    setPassword('');

    if (authStatus !== 200) {
      setIncorrect(true)
      return
    }

    setCredentials(credentials)
    setLoggedIn(true)
    setEmail('');

    navigate('/prompts');

  };

  return (
    <div className="login-container">
      <h2 className="login-title">Login</h2>
      <form className="login-form" onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        {incorrect && (<div className="password-mismatch">Incorrect credentials!</div>) }

        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default Login;
