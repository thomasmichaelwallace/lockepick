const path = require('path');
const webpack = require('webpack');

const TerserPlugin = require('terser-webpack-plugin');

module.exports = {
  mode: 'development',
  plugins: [new webpack.ProgressPlugin()],

  devServer: {
    publicPath: '/dist/',
  },

  module: {
    rules: [{
      test: /\.(js|jsx)$/,
      include: [path.resolve(__dirname, 'src')],
      loader: 'babel-loader',
    }],
  },

  optimization: {
    minimizer: [new TerserPlugin()],

    splitChunks: {
      cacheGroups: {
        vendors: {
          priority: -10,
          test: /[\\/]node_modules[\\/]/,
        },
      },

      chunks: 'async',
      minChunks: 1,
      minSize: 30000,
      name: 'chunk',
    },
  },
};
