import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './index.css';
import snappy_logo from '../../assets/snappy_logo.svg';

const Home = () => {
  const [email, setEmail] = useState("");
  const navigate = useNavigate();
  const login = useNavigate();

  const createAccount = () => {
    navigate('/create-account');
  };

  const loginClick = (f: {preventDefault:()=>void;}) => {
    f.preventDefault();
    login('/login');

  }

  return (
    <>
    <div className="home">
      <div className="logo"><a href="/"><img src={snappy_logo} id="logo" /></a></div>
      <div className="navbar">
        <button className="nav-button" onClick={loginClick}>Sign In</button>
      </div>
      <div className="main-content">
        <h1>Financial help in a snap</h1>
        <h3>Get the most up to date advice.</h3>

        <div className="sign-up">
          <h4>Enter your email to create an account.</h4>
            <button onClick={createAccount} className="get-started-button">Get Started</button>
        </div>
      </div>
    </div>
    </>
  );
};

export default Home;
