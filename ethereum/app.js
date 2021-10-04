import Vue from 'vue'
import App from './src/App.vue'
//https://newbedev.com/webpack-config-how-to-just-copy-the-index-html-to-the-dist-folder
require('file-loader?name=[name].[ext]!./index.html');

Vue.component('app', App);

new Vue({
    el: '#app',
    component: { App }
})
