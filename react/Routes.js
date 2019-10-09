import React from 'react';
import { hot } from 'react-hot-loader';

import User from './containers/User';

const Routes = () => {

  return(
    <div>
      <h1>Hello</h1>
      <User/>
    </div>
  )
}

export default hot(module)(Routes);