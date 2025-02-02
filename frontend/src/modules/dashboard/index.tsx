import './index.css';

const MainMenu: React.FC = () => {

    return(
        <>
            <div className="container-fluid" id="main-container">
                <div className="row flex-nowrap">
                    <div className="col-auto col-md-3 col-xl-2 px-sm-2 px-0 bg-dark">
                        <div className="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
                            <a href="/" className="d-flex align-items-center pb-3 mb-md-0 me-md-auto text-white text-decoration-none">
                                <span className="fs-5 d-none d-sm-inline">Navigation</span>
                            </a>
                            <ul className="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start" id="menu">
                                <li className="nav-item">
                                    <a href="/" className="nav-link px-0 align-middle">
                                        <i className="fs-4 bi-table"></i> <span className="ms-1 d-none d-sm-inline">Home</span>
                                    </a>
                                </li>
                                <li className="nav-item">
                                    <a href="/" className="nav-link px-0 align-middle">
                                        <i className="fs-4 bi-table"></i> <span className="ms-1 d-none d-sm-inline">Option 1</span> 
                                    </a>
                                </li>
                                <li className="nav-item">
                                    <a href="/" className="nav-link px-0 align-middle">
                                        <i className="fs-4 bi-table"></i> <span className="ms-1 d-none d-sm-inline">Option 2</span>
                                    </a>
                                </li>
                                <li className="nav-item">
                                    <a href="/" className="nav-link px-0 align-middle ">
                                        <i className="fs-4 bi-table"></i> <span className="ms-1 d-none d-sm-inline">Option 3</span>
                                    </a>
                                </li>
                                <li className="nav-item">
                                    <a href="/" className="nav-link px-0 align-middle">
                                        <i className="fs-4 bi-table"></i> <span className="ms-1 d-none d-sm-inline">Option 4</span> 
                                    </a>
                                </li>
                            </ul>
                        </div>
                    </div>
                    <div className="col py-3" id="main-content-area">
                        Hello
                    </div>
                </div>
            </div>
        </>
    );
}

export default MainMenu