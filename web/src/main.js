import babelpolyfill from 'babel-polyfill'
import Vue from 'vue'
import App from './App'
import router from './router';
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import store from './store'
import Vuex from 'vuex'

import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import localStore from 'utils/local_store';

// import 'font-awesome/css/font-awesome.min.css'
import 'styles/index.scss'
Vue.use(ElementUI)

Vue.use(Vuex)


// NProgress.configure({ showSpinner: false })

// router.beforeEach((to, from, next) => {
//     NProgress.start();
//     if (to.path == '/login') {
//         localStore.remove('user');
//         localStorage.remove("permissions");
//     }
//     let user = JSON.parse(localStore.get('user'));
//     if (!user && to.path != '/login') {
//         next({ path: '/login' })
//     } else {
//         next()
//     }
// })
// router.afterEach(() => {
//     NProgress.done();
// });
new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')