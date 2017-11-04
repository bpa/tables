const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: './js/index.js',
  output: {
    filename: 'tables.js',
    path: path.resolve(__dirname, 'dist')
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: 'index.html'
    })
  ],
  module: {
    rules: [
      { test: /\.css$/,
        use: ['style-loader', 'css-loader']
      },
      { test: /\.js$/,
        loader: 'babel-loader',
        include: path.resolve(__dirname, "js")
      }
    ]
  }
};
