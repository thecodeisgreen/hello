const R = require('ramda');
import React from 'react';
import ReactDOM from 'react-dom';
import Routes from './Routes';
import Environment from './_graphql/Environment';
import Tokens from './lib/Tokens';

let environment;

const getDocumentRoot = () => {
  var element = document.createElement('div');
  document.body.appendChild(element);
  return element;
};

Tokens.setKey(CLIENT_ID, CLIENT_SECRET);

(async () => {
  environment = await Environment();
  ReactDOM.render(
    <div>
      <Routes/>
   </div>,
    getDocumentRoot()
  );
})();

/*

ReactDOM.render(
  <div>
    <Routes/>
 </div>,
  getDocumentRoot()
);
*/

export {
  environment
};