// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuex from 'vuex'
import Mint from 'mint-ui'

import store from './store'
import localStore from 'utils/local_store';
import 'mint-ui/lib/style.css';
import axios from 'axios';

import stringTimeFormat from 'utils/filters';
import lazyLoadMenusRoutes from 'utils/lazyload';

Vue.prototype.$ajax = axios;

Vue.config.productionTip = false
Vue.use(Mint);
Vue.use(Vuex)

let backgroundMenus = JSON.parse(localStore.get("backgroundMenus"));
if (backgroundMenus) {
    store.commit("GLOBAL_SET_UER_MENUS", backgroundMenus);
    // 加载后台菜单
    let menusRoutes = lazyLoadMenusRoutes(backgroundMenus);
    router.addRoutes([menusRoutes]);
    store.commit("GLOBAL_LOAD_ROUTES_DONE");

}

// NProgress.configure({ showSpinner: false })

router.beforeEach((to, from, next) => {
    // NProgress.start();
    if (to.path == '/login') {
        localStore.remove('userinfo');
        localStore.remove("groups");
        localStore.remove("backgroundMenus");
    }
    let user = JSON.parse(localStore.get('userinfo'));
    let path = to.path;
    // 后台必须登录
    let adminPath = path.indexOf('/admin');
    if (adminPath == 0) {
        if (!user && path != '/login') {
            next({ path: '/login' })
        } else {
            next()
        }
    } else {
        next()
    }
});
new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')