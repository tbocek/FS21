const { VueLoaderPlugin } = require('vue-loader')
const NodePolyfillPlugin = require("node-polyfill-webpack-plugin")
module.exports = {
    mode: 'development',
    entry: {
        app: './app.js',
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /\.css$/,
                use: ['vue-style-loader', 'css-loader']
            }
        ]
    },
    resolve: {
        alias: {
            vue: 'vue/dist/vue.js'
        },
    },
    plugins: [
        new VueLoaderPlugin(),
        new NodePolyfillPlugin(),
    ]
}
