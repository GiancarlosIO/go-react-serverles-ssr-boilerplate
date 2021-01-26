/** @type {import('webpack').Configuration} */
const path = require('path');

const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const webpackMerge = require('webpack-merge');

const base = require('./webpack.base');

module.exports = webpackMerge.merge(base, {
  mode: 'production',
  devtool: 'source-map',
  output: {
    filename: `app.[contenthash].min.js`,
    chunkFilename: '[name]-[chunkhash].min.js',
  },
  module: {
    rules: [
      {
        test: /\.scss$/i,
        use: [
          MiniCssExtractPlugin.loader,
          'css-loader',
          'sass-loader',
          'postcss-loader',
        ],
      },
    ],
  },
  plugins: [new MiniCssExtractPlugin()],
});
