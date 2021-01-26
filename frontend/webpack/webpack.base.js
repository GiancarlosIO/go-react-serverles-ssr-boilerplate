/** @type {import('webpack').Configuration} */
const path = require('path');

const HtmlWebpackPlugin = require('html-webpack-plugin');
const HtmlWebpackHarddiskPlugin = require('html-webpack-harddisk-plugin');
const FaviconsWebpackPlugin = require('favicons-webpack-plugin');

const {
  babelLoader,
  fileLoader,
  assetResource,
  cssLoader,
} = require('./loaders');

module.exports = {
  context: __dirname,
  entry: '../src/index.tsx',
  output: {
    filename: `app.min.js`,
    path: path.resolve(__dirname, '../dist/static'),
    chunkFilename: '[name]-[id].min.js',
    publicPath: '/static/',
  },
  resolve: {
    extensions: ['.mjs', '.js', '.jsx', '.ts', '.tsx'],
  },
  module: {
    rules: [babelLoader, fileLoader, assetResource, cssLoader],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: '../../templates/app.gohtml',
      alwaysWriteToDisk: true,
      filename: 'app.gohtml',
      minify: false,
    }),
    new FaviconsWebpackPlugin({
      logo: '../src/Images/mr-n-logo.png',
      cache: true,
      prefix: '',
      favicons: {
        appName: 'Mr N',
        appDescription: "Hello! I'm Mr. N, a frontend enginner :)",
        developerName: 'Mr N',
        developerURL: null,
        background: '#1794cd',
        theme_color: '#43c794',
      },
    }),
    new HtmlWebpackHarddiskPlugin(),
  ],
};
