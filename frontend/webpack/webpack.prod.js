/** @type {import('webpack').Configuration} */
const path = require('path');
const webpackMerge = require('webpack-merge');

const base = require('./webpack.base');

module.exports = webpackMerge.merge(base, {
  mode: 'production',
  devtool: 'source-map',
  output: {
    filename: `app.[contenthash].min.js`,
    chunkFilename: '[name]-[chunkhash].min.js',
  },
});
