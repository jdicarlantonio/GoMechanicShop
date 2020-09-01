import React, { Fragment } from 'react';
import './App.css';

import AddCustomer from './components/AddCustomer'

function App() {
  return (
    <Fragment>
      <div className="container">
        <AddCustomer />
      </div>
    </Fragment>
  );
}

export default App;
