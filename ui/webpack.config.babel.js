// @flow

import webpack from 'webpack'
import path from 'path'

import { WDS_PORT } from './src/constants/config'
import { isProd } from './src/constants/util'

export default {
  entry: [
    'react-hot-loader/patch',
    './src/index.js'
  ],
  resolve: {
    extensions: ['.js', '.jsx'],
  },
  module: {
    rules: [
      { 
        test: /\.(js|jsx)$/,
        use: 'babel-loader', 
        exclude: /node_modules/ 
      },
    ],
  },
  output: {
    filename: 'js/bundle.js',
    path: path.resolve(__dirname, 'dist'),
    publicPath: isProd 
      ? '/static' 
      : `http://localhost:${WDS_PORT}/dist/`,
  },
  devtool: isProd 
    ? false 
    : 'source-map',
  devServer: {
    contentBase: path.join(__dirname, "public/"),
    publicPath: `http://localhost:${WDS_PORT}/dist/`,
    port: WDS_PORT,
    hot: true,
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin()
  ],
}
