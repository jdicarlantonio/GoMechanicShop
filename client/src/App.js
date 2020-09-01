import React, { Fragment } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import './App.css';

import AddCustomer from './components/AddCustomer';
import AddMechanic from './components/AddMechanic';

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
              </Switch>
            </div>
          </div>
        </Router>

    </Fragment>
  );
}

export default App;
