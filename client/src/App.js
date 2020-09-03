import React, { Fragment } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import './App.css';

import AddCustomer from './components/AddCustomer';
import AddMechanic from './components/AddMechanic';
import AddCar from './components/AddCar';
import OpenServiceRequests from './components/OpenServiceRequests';
import CreateServiceRequest from './components/CreateServiceRequest';
import CloseServiceRequest from './components/CloseServiceRequest';

function App() {
  return (
    <Fragment>

        <Router>
          <div>
            <nav>
              <ul>
                <li>
                  <Link to="/">Home</Link>
                </li>
                <li>
                  <Link to="/addCustomer">Add Customer</Link>
                </li>
                <li>
                  <Link to="/addMechanic">Add Mechanic</Link>
                </li>
                <li>
                  <Link to="/addCar">Add Car</Link>
                </li>
                <li>
                  <Link to="/createServiceRequest">Open a Service Request</Link>
                </li>
                <li>
                  <Link to="/getOpenServiceRequests">Get Open Service Requests</Link>
                </li>
                <li>
                  <Link to="closeServiceRequest">Close a Service Request</Link>
                </li>
              </ul>
            </nav>
            <div className="container">
              <Switch>
                <Route path="/addCustomer">
                  <AddCustomer />
                </Route>
                <Route path="/addMechanic">
                  <AddMechanic />
                </Route>
                <Route path="/addCar">
                  <AddCar />
                </Route>
                <Route path="/createServiceRequest">
                  <CreateServiceRequest />
                </Route>
                <Route path="/getOpenServiceRequests">
                  <OpenServiceRequests />
                </Route>
                <Route path="/closeServiceRequest">
                  <CloseServiceRequest />
                </Route>
              </Switch>
            </div>
          </div>
        </Router>

    </Fragment>
  );
}

export default App;
