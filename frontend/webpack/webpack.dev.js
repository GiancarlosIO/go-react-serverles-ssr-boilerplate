/** @type {import('webpack').Configuration} */
const path = require('path');
const webpackMerge = require('webpack-merge');

const base = require('./webpack.base');

module.exports = webpackMerge.merge(base, {
  mode: 'development',
  output: {
    publicPath: 'http://localhost:9000/static/',
  },
  devServer: {
    port: 9000,
    publicPath: 'http://localhost:9000/static/',
  },
});
