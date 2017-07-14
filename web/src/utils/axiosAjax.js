import Vue from 'vue';
import axios from 'axios';


Vue.prototype.$ajax = axios

/**
 * 封装axios的请求
 * @param {string} type 请求方式:get post
 * @param {string} url 请求地址
 * @param {object} data 请求数据
 * @param {function} callback 请求回调函数
 * @param {function} errback 错误梳理函数
 * @param {bool} tokenFlag 是否需要携带token参数，为true，不需要；false，需要。一般除了登录，都需要
 */
export default function axiosAjax(type, url, data, callback, errback, {
    cbFn,
    tokenFlag,
    errFn,
    host,
    headers,
    axios_opts
} = {}) {
    var options = {
        method: type,
        url: url,
        headers: headers && typeof headers === 'object' ? headers : {}
    };
    options[type === 'get' ? 'params' : 'data'] = data;

    if (tokenFlag !== true) {
        //如果你们的后台不会接受headers里面的参数，打开这个注释，即实现token通过普通参数方式传
        // data.token = this.$store.state.user.userinfo.token;
        options.headers.token = this.$store.state.global.userinfo.token;
    }
    //axios内置属性均可写在这里
    if (axios_opts && typeof axios_opts === 'object') {
        for (var f in axios_opts) {
            options[f] = axios_opts[f];
        }
    }
    //发送请求
    Vue.axios(options).then(callback(res)).catch(errback);
}