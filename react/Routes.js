import React, { useEffect } from 'react';
import { hot } from 'react-hot-loader';
import { useCookies } from 'react-cookie';

const Routes = () => {
  const [cookies] = useCookies(['session']);

  useEffect(() => {
    fetch('/_/info')
  }, [])

  return(
    <div>
      <h1>Hello</h1>
      {JSON.stringify(cookies)}
    </div>
  )
}

export default hot(module)(Routes);