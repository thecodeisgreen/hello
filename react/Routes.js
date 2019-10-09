import React, {useState, useEffect} from 'react';
import { hot } from 'react-hot-loader';


const Routes = () => {

  const [firstname, setFirstname] = useState(null)
  const [lastname, setLastname] = useState(null)

  useEffect(() => {
    const getFirstname = async () => {
      const response = await fetch('/firstname', {
        headers: { 'Content-Type': 'application/json' }
      });
      const json = await response.json();
      setFirstname(json.firstname);
    }

    const getLastname = async () => {
      const response = await fetch('/lastname', {
        headers: { 'Content-Type': 'application/json' }
      });
      const json = await response.json();
      setLastname(json.lastname);
    }

    getFirstname()
    getLastname()
  }, [])

  return(
    <div>
    <h1>{`hello ${firstname}`}</h1>
    <h2>- {lastname} -</h2>
    <h3>--------</h3>
    </div>
  )
}

export default hot(module)(Routes);