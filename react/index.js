const R = require('ramda');
import React from 'react';
import ReactDOM from 'react-dom';

const getDocumentRoot = () => {
  var element = document.createElement('div');
  document.body.appendChild(element);
  return element;
};

ReactDOM.render(
  <div>
    <h1>Hello</h1>
  </div>,
  getDocumentRoot()
);