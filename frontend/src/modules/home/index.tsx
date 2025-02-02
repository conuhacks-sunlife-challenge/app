import { useState } from 'react';
import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom';
import './index.css';


const Home = () => {
  const [email, setEmail] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e: { preventDefault: () => void; }) => {
    e.preventDefault();
    console.log("Email:", email);
    navigate('/create-account');
  };

  return (
    <div className="home">
      <div className="logo">Snappy</div>
      <div className="navbar">
        <button className="nav-button">Sign In</button>
      </div>
      <div className="main-content">
        <h1>Financial help in a snap</h1>
        <h3>Get the most up to date advice.</h3>

        <div className="sign-up">
          <h4>Enter your email to create an account.</h4>
          <form onSubmit={handleSubmit}>
            <input
              type="email"
              placeholder="Email address"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
            <button type="submit" className="get-started-button">Get Started</button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Home;
