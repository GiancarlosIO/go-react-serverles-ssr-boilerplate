const path = require('path');
const slsw = require('serverless-webpack');
const nodeExternals = require('webpack-node-externals');
const TsconfigPathsPlugin = require('tsconfig-paths-webpack-plugin');

const { babelLoader, fileLoader, assetResource } = require("../frontend/webpack/loaders")

const include = [
  path.resolve(__dirname, "../frontend/src"),
  path.resolve(__dirname, './src'),
]

const exclude = [
  path.resolve(__dirname, 'node_modules'),
  path.resolve(__dirname, '.serverless'),
  path.resolve(__dirname, '.webpack'),
]

const createLoader = (loader) => ({
  ...loader,
  exclude,
  include,
})

module.exports = {
  context: __dirname,
  mode: slsw.lib.webpack.isLocal ? 'development' : 'production',
  entry: slsw.lib.entries,
  devtool: slsw.lib.webpack.isLocal ? 'eval-cheap-module-source-map' : 'source-map',
  resolve: {
    extensions: ['.mjs', '.json', '.ts', '.tsx', '.js', '.jsx'],
    symlinks: false,
    cacheWithContext: false,
    plugins: [
      new TsconfigPathsPlugin({
        configFile: './tsconfig.paths.json',
      }),
    ],
  },
  output: {
    libraryTarget: 'commonjs',
    path: path.join(__dirname, '.webpack'),
    filename: '[name].js',
  },
  optimization: {
    concatenateModules: false,
  },
  target: 'node',
  externals: [nodeExternals()],
  module: {
    rules: [
      createLoader(babelLoader),
      createLoader(fileLoader),
      createLoader(assetResource),
    ],
  },
  plugins: [],
};