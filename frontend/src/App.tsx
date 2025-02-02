import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './modules/home/index.tsx';
import Dashboard from './modules/login/index.tsx';
import MainMenu from './modules/dashboard/index.tsx';

const AppWrapper = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/create-account" element={<Dashboard />} />
        <Route path="/main-menu" element={<MainMenu />} />
      </Routes>
    </Router>
  );
};

export default AppWrapper;