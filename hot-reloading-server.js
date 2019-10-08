
require ('@babel/register' );
require('@babel/polyfill');

require('dotenv').config();

const express = require('express')
const app = express();
const server = require('http').createServer(app);
const path= require('path')

app.use(express.static(path.join(__dirname, '../public')));

if (process.env.NODE_ENV === 'development') {
  const webpack = require('webpack');
  const webpackMerge = require('webpack-merge');
  const webpackDevMiddleware = require('webpack-dev-middleware');
  const webpackHotMiddleware = require('webpack-hot-middleware');
  const history = require('connect-history-api-fallback');

  const configDev = require('./webpack.dev.js');
  const config = webpackMerge(configDev);

  const compiler = webpack(config);

  app.use(history());
  app.use(webpackDevMiddleware(compiler));
  app.use(webpackHotMiddleware(compiler));
} else {
  app.get ('*', (req, res) => {
    res.sendFile(path.join('/service', 'public', 'dist', 'index.html'));
  });
}
server.listen(process.env.HOT_RELOADING_SERVER_PORT, () => console.log('hot-reloading-server is ready'));