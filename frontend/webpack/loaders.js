const path = require('path');

module.exports.fileLoader = {
  test: /\.(png|jpe?g|gif)$/i,
  use: 'file-loader',
  include: path.resolve(__dirname, '../src'),
  exclude: /node_modules/,
};

module.exports.babelLoader = {
  test: /\.(js|jsx|ts|tsx)$/,
  use: 'babel-loader',
  include: path.resolve(__dirname, '../src'),
  exclude: /node_modules/,
};

module.exports.assetResource = {
  test: /\.(png|svg|jpg|jpeg|gif|ico|favicon|xml|webapp)$/i,
  type: 'asset/resource',
  exclude: /node_modules/,
};

module.exports.cssLoader = {
  test: /\.(css|sass|scss)$/i,
  use: ['style-loader', 'css-loader', 'sass-loader', 'postcss-loader'],
};
