import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './index.css';
import User from '../../types/user';
import endpoints from '../../endpoints';

const Dashboard: React.FC = () => {

  const navigate = useNavigate();
  
  const [firstName, setFirstName] = useState<string>('');
  const [lastName, setLastName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [confirmPassword, setConfirmPassword] = useState<string>('');
  const [passwordMatch, setPasswordMatch] = useState<boolean>(true);
  const [failedCreateUser, setFailedCreateUser] = useState<boolean>(false);

  const authenticate = async (user: User) => {
    const body = JSON.stringify(user)
    const req = new Request(endpoints.NewUser, {

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

    if (!passwordMatch) {
      return
    }
    console.log('Form submitted:', { firstName, lastName, email, password });
    const authStatus = await authenticate({
      Email: email,
      Password: password,
      FirstName: firstName,
      LastName: lastName
    })

    console.log("auth status:", authStatus)
    if (authStatus !== 200) {
      setFailedCreateUser(true)
      return
    }
    
    setFirstName('');
    setLastName('');
    setEmail('');
    setPassword('');
    setConfirmPassword('');

    navigate('/login');

  };

  const handleConfirmPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setConfirmPassword(e.target.value);
    setPasswordMatch(e.target.value === password);
  };

  const LoginInformation = () => {
    if (failedCreateUser) {
      return (
            <div className="password-mismatch">User already exists!</div>
      )
    }
    if (!passwordMatch) {
      return (
            <div className="password-mismatch">Passwords do not match!</div>

      )
    }
    if (passwordMatch && confirmPassword) {
      return (
            <div className="password-match">Passwords match!</div>
      )
    }
  }

  return (
    <div className="dashboard-container">
      <h2 className="create-account-title">Create Your Account</h2>
      <form className="dashboard-form" onSubmit={handleSubmit}>
        <div>
          <label htmlFor="first-name">First Name</label>
          <input
            type="text"
            id="first-name"
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
            placeholder="Enter your first name"
            required
          />
        </div>
        <div>
          <label htmlFor="last-name">Last Name</label>
          <input
            type="text"
            id="last-name"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
            placeholder="Enter your last name"
            required
          />
        </div>
        <div>
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Enter your email"
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
            placeholder="Enter your password"
            required
          />
        </div>
        <div>
          <label htmlFor="confirm-password">Confirm Password</label>
          <input
            type="password"
            id="confirm-password"
            value={confirmPassword}
            onChange={handleConfirmPasswordChange}
            placeholder="Confirm your password"
            required
          />
          <LoginInformation/>
        </div>
        <button type="submit">Sign Up</button>
      </form>
    </div>
  );
};

export default Dashboard;
