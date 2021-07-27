const path = require("path");

const HtmlWebpackPlugin = require("html-webpack-plugin");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");
const VueLoaderPlugin = require("vue-loader/lib/plugin");
module.exports = {
  mode: "development",
  entry: "./src/main.js",
  output: {
    filename: "js/[name].js",
    path: path.resolve(__dirname, "dist"),
  },
  resolve: {
    // 配置省略文件路径的后缀名
    extensions: [".vue", ".tsx", ".ts", ".js"],
  },
  module: {
    rules: [
      /** vue */
      { test: /\.vue$/, use: ["vue-loader"] },
      {
        test: /\.ts$/,
        loader: "ts-loader",
        options: { appendTsSuffixTo: [/\.vue$/] },
      },
      {
        test: /\.js$/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-env"],
          },
        },
      },
      /** css 样式处理*/
      {
        test: /\.css$/,
        // use: [{ loader: "style-loader", options: {} }],
        use: ["style-loader", "css-loader"],
      },
      { test: "/.less$/", use: ["style-loader", "css-loader", "less-loader"] },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, "./public/index.html"), //
      filename: "index.html",
      title: "webpack搭建",
    }),
    new CleanWebpackPlugin(), // 情况打包目录
    new VueLoaderPlugin(),
  ],
};
