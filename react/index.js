const R = require('ramda');
import React from 'react';
import ReactDOM from 'react-dom';
import Routes from './Routes';

const getDocumentRoot = () => {
  var element = document.createElement('div');
  document.body.appendChild(element);
  return element;
};

ReactDOM.render(
  <div>
    <Routes/>
  </div>,
  getDocumentRoot()
);