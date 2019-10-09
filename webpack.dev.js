const R = require('ramda');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');

const path = require('path');

console.log(__dirname)
const PATHS = {
  src: path.join(__dirname, 'react'),
  build: path.join(__dirname, 'public', 'dist'),
};

console.log(JSON.stringify(PATHS))

module.exports = {
  mode: 'development',

  target: 'web',
  
  resolve: {
    alias: {
      'react-dom': '@hot-loader/react-dom'
    },
  },

  entry: [
    `${PATHS.src}/index.js`,
    'webpack-hot-middleware/client'
  ],

  output: {
    filename: '[name].bundle.js',
    path: PATHS.build,
    publicPath: '/'
  },
  
  devtool: 'inline-source-map',
  
  devServer: {
    hot: true,
    contentBase: PATHS.build
  },

  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/env', '@babel/react'],
            plugins: [
              '@babel/plugin-syntax-dynamic-import'
            ],
            cacheDirectory: true
          } 
        }
      },
      {
        test: [/\.less$/, /\.css$/],
        include: [path.join(__dirname, 'react'), path.join(__dirname, 'node_modules')],
        use: [
          { loader: 'style-loader' },
          { loader: 'css-loader' },
          {
            loader: 'postcss-loader', 
            options: {
              plugins: () => [
                require('autoprefixer')
              ]
            }
          },
          { loader: 'less-loader',
            options: {
              modifyVars: {
                'hack': `true; @import "${__dirname}/react/less/antd-custom.less";`
              },
              javascriptEnabled: true
            }
          }
        ]
      },
      {
        test: /\.(png|jpg)$/,
        loader: 'url-loader'
      },
      {
        test: /\.(woff(2)?|ttf|eot|svg)(\?v=\d+\.\d+\.\d+)?$/,
        use: [{
          loader: 'file-loader',
          options: {
            name: '[name].[ext]',
            outputPath: 'fonts/'
          }
        }]
      }
    ]
  },
    
  plugins: [
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify('development'),
      API_USER_ID: `"client_id"`,
      API_KEY: `"client_secret"`
    }),
    new CleanWebpackPlugin({ verbose: true }),
    new HtmlWebpackPlugin({
      title: 'hello',
      template: require('html-webpack-template'),
      filename: 'index.html',
      inject: false,
      mobile: true
    }),
    new webpack.NamedModulesPlugin(),
    new webpack.HotModuleReplacementPlugin()
  ]
};