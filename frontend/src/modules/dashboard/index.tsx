import './index.css';
import snappy_logo from'../../assets/snappy_logo.svg';

const MainMenu: React.FC = () => {

    return(
        <>
            <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4" id="top-navigation">
                <a className="navbar-brand" href="/main-menu"><img src={snappy_logo} id="logo" /></a>
                
                <div className="collapse navbar-collapse" id="navbarCollapse">
                    <ul className="navbar-nav">
                    <li className="nav-item">
                        <a className="nav-link" href="">Option 1</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="">Option 2</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="">Option 3</a>
                    </li>
                    </ul>
                </div>
            </nav>
            <div id="main-content-area">
                Hello!
            </div>
        </>
    );
}

export default MainMenu