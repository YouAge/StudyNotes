# webpack使用

## 安装

1. 全局安装

   ```shell
   npm install webpack webpack-cli -g
   ```

## 创建项目

1.  `npm init -y` 完成项目初始化，

   > 自己创建文件一个项目文件， 只能使用英文， 在命令行中使用命令，
   >
   > 如果`npm init -y`初始化失败，使用为自己创建的项目名有问题，使用 `npm init`手动输入
   >
   > 创建成功后，目录下出现一个`package.json`文件



## vue的

>VSC 中 使用 vue的配置， `vetur` 插件
>
>ctrl+shift+p  搜索json ， 找到 open setting 的json 点开，配置

```json

{
  "editor.tabSize": 2,
  "editor.formatOnSave": true,
  // 在style样式中, 使用缩进, 缩进是2个空格
  "vetur.format.styleInitialIndent": true,
  // 在script样式中, 使用缩进, 缩进2个空格
  "vetur.format.scriptInitialIndent": true,
  "vetur.format.defaultFormatterOptions": {
    // 用于template部分的格式化
    "prettyhtml": {
      // 超过80个字符的标签换行显示
      "printWidth": 80
    },
    // 用于script部分的格式化
    "prettier": {
      // 使用单引号
      "singleQuote": true,
      // 不使用;号
      "semi": false,
      "proseWrap": "never",
      "printWidth": 80
    }
  },
  // stylus不使用{}
  "stylusSupremacy.insertBraces": false,
  // stylus不使用;
  "stylusSupremacy.insertSemicolons": false
}
```



http://blog.brojie.cn/web/todolist/#%E8%AF%BE%E7%A8%8B%E5%8C%85%E6%8B%AC









# vue 配置

## 抽取css 按需加载

```sh
 npm install mini-css-extract-plugin -D
```

###　`vue.config.js`配置

```js
chainWebpack: config => {
  let miniCssExtractPlugin = new MiniCssExtractPlugin({
    filename: 'assets/[name].[hash:8].css',
    chunkFilename: 'assets/[name].[hash:8].css'
  })
  config.plugin('extract-css').use(miniCssExtractPlugin)
}
```

## 图片按需加载

```sh
 npm install image-webpack-loader -D
```



```js
hainWebpack: config => {
config.module.rule('images')
    .test(/\.(png|jpe?g|gif|webp)(\?.*)?$/)
    .use('image-webpack-loader')
    .loader('image-webpack-loader')
    .options({
      bypassOnDebug: true
    })
    .end()
}
```

> 图片压缩可以在：https://tinypng.com/ 进行批量压缩



## 配置less的全局变量

####　方法一

```sh
npm i style-resources-loader -D
```

**vue-cli4 配置 vue,config,js**

> 配置成功，vue脚手架版本是4.5

```js
module.exports={
    chainWebpack: config => {
        // 配置全局的less 变量使用、
        const types = ['vue-modules', 'vue', 'normal-modules', 'normal']
        types.forEach(type => addStyleResource(config.module.rule('less').oneOf(type)))
    }
}

// 配置全局的less的变量使用， 先安装 npm install style-resources-loader -D
function addStyleResource(rule) {
    rule.use('style-resource')
        .loader('style-resources-loader')
        .options({
            patterns: [path.resolve(__dirname, "./src/assets/style/global.less")]
        })
}
```

#### 反法二

```shell
npm install style-resources-loader vue-cli-plugin-style-resources-loader less-loader less -S
```

**vue-cli3 配置 vue,config,js**

> 在最新版本4.5中没有配置成功，可能是版本问题

```js

const path = require('path');
module.exports = {
pluginOptions: {
    'style-resources-loader': {
      preProcessor: 'less',
      patterns: [
        path.resolve(__dirname, "./src/assets/style/global.less")
      ]
    }
  }
}
```



## gzip压缩代码

```shell
 npm install compression-webpack-plugin -D
```

**vue.config.js**

```js
const CompressionWebpackPlugin = require('compression-webpack-plugin');

configureWebpack: config => {
// 开启gzip压缩
  config.plugins.push(
    new CompressionWebpackPlugin(
      {
        filename: info => {
          return `${info.path}.gz${info.query}`
        },
        algorithm: 'gzip',
        threshold: 10240, // 只有大小大于该值的资源会被处理 10240
        test: new RegExp('\\.(' + ['js'].join('|') + ')$'
        ),
        minRatio: 0.8, // 只有压缩率小于这个值的资源才会被处理
        deleteOriginalAssets: false // 删除原文件
      }
    )
  )
}
```

##　echarts按需导入

```shell
npm install babel-plugin-equire -D
```

**1. 创建一个文件，echarts.js**

````js
// eslint-disable-next-line
  const echarts = equire([
    // 写上你需要的 echarts api
    "tooltip",
    "line"
  ]);

  export default echarts;
````

**2. babel.config.js配置**

````js
module.exports = {
  presets: [
    '@vue/app'
  ],
  plugins: [
    [
      "component",
      {
        libraryName: "element-ui",
        styleLibraryName: "theme-chalk"
      }
    ],
    "equire"
  ]
}
````

**3.单页面使用**

```js
// 直接使用 
 import echarts from '@/lib/util/echarts.js' 
 
 this.myChart = echarts.init(this.$refs.chart)
```



## lodash按需加载

```shell
npm install lodash-webpack-plugin --save-dev
```

**1.babel.config.js**

```js
module.exports = {
  presets: [
    '@vue/app'
  ],
  plugins: [
    [
      "component",
      {
        libraryName: "element-ui",
        styleLibraryName: "theme-chalk"
      }
    ],
    "lodash",
    "equire"
  ]
}

```

**2.vue.config.js**

````js
const LodashModuleReplacementPlugin = require("lodash-webpack-plugin");

chainWebpack: config => {
    config
    .plugin("loadshReplace")
    .use(new LodashModuleReplacementPlugin());
}
````

##　prefetch，preload

> 删除无用的插件，避免加载多余的资源（如果不删除的话，则会在 index.html 里面加载 无用的 js 文件）

```js
chainWebpack: config => {
    // 移除prefetch插件，避免加载多余的资源
    config.plugins.delete('prefetch')
    / 移除 preload 插件，避免加载多余的资源
    config.plugins.delete('preload');
}
```





# vue.config.js完整配置

> @vue-cli3.0以上使用  4.5版本

```js
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CompressionWebpackPlugin = require('compression-webpack-plugin');
const LodashModuleReplacementPlugin = require("lodash-webpack-plugin");

module.exports = {
  productionSourceMap: false, // 关闭生产环境的 source map
  lintOnSave: false,
  publicPath: process.env.VUE_APP_PUBLIC_PATH,
  devServer: {
    host: "localhost",
    port: 3002,
    proxy: {
      '/api': {
        target: "https://tapi.quanziapp.com/api/",
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      },
    }
  },

  chainWebpack: config => {

    // 移除 prefetch 插件
    config.plugins.delete('prefetch');
    // 移除 preload 插件，避免加载多余的资源
    config.plugins.delete('preload');
 
    config.optimization.minimize(true);

    config.optimization.splitChunks({
      chunks: 'all'
    })

    config
      .plugin('webpack-bundle-analyzer')
      .use(require('webpack-bundle-analyzer').BundleAnalyzerPlugin)

    if (process.env.NODE_ENV !== 'development') {

      let miniCssExtractPlugin = new MiniCssExtractPlugin({
        filename: 'assets/[name].[hash:8].css',
        chunkFilename: 'assets/[name].[hash:8].css'
      })
      config.plugin('extract-css').use(miniCssExtractPlugin)
      config.plugin("loadshReplace").use(new LodashModuleReplacementPlugin());

      config.module.rule('images')
        .test(/\.(png|jpe?g|gif|webp)(\?.*)?$/)
        .use('image-webpack-loader')
        .loader('image-webpack-loader')
        .options({
          bypassOnDebug: true
        })
        .end()
        .use('url-loader')
        .loader('file-loader')
        .options({
          name: 'assets/[name].[hash:8].[ext]'
        }).end()
      config.module.rule('svg')
        .test(/\.(svg)(\?.*)?$/)
        .use('file-loader')
        .loader('file-loader')
        .options({
          name: 'assets/[name].[hash:8].[ext]'
        })
    }
  },
  configureWebpack: config => {
    // config.plugins.push(["equire"]);

    if (process.env.NODE_ENV !== 'development') {
      config.output.filename = 'assets/[name].[hash:8].js'
      config.output.chunkFilename = 'assets/[name].[hash:8].js'
    }
    // 公共代码抽离
    config.optimization = {
      // 分割代码块
      splitChunks: {
        cacheGroups: {
          //公用模块抽离
          common: {
            chunks: 'initial',
            minSize: 0, //大于0个字节
            minChunks: 2, //抽离公共代码时，这个代码块最小被引用的次数
          },
          //第三方库抽离
          vendor: {
            priority: 1, //权重
            test: /node_modules/,
            chunks: 'initial',
            minSize: 0, //大于0个字节
            minChunks: 2, //在分割之前，这个代码块最小应该被引用的次数
          },
        },
      }
    }
    // 开启gzip压缩
    config.plugins.push(
      new CompressionWebpackPlugin(
        {
          filename: info => {
            return `${info.path}.gz${info.query}`
          },
          algorithm: 'gzip',
          threshold: 10240, // 只有大小大于该值的资源会被处理 10240
          test: new RegExp('\\.(' + ['js'].join('|') + ')$'
          ),
          minRatio: 0.8, // 只有压缩率小于这个值的资源才会被处理
          deleteOriginalAssets: false // 删除原文件
        }
      )
    )
  },
  css: {
    extract: true,
    sourceMap: false,
    loaderOptions: {
      sass: {
      },
    },
  },
}
```





# vue中使用rem单位



