/** @type {import('webpack').Configuration} */
const webpack = require('webpack');
const webpackMerge = require('webpack-merge');
const ReactRefreshWebpackPlugin = require('@pmmmwh/react-refresh-webpack-plugin');

const base = require('./webpack.base');

module.exports = webpackMerge.merge(base, {
  mode: 'development',
  devtool: 'eval-cheap-module-source-map',
  output: {
    publicPath: 'http://localhost:9000/static/',
  },
  devServer: {
    port: 9000,
    publicPath: 'http://localhost:9000/static/',
    hot: true,
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new ReactRefreshWebpackPlugin({
      overlay: {
        sockIntegration: 'wds',
      },
    }),
  ],
});
