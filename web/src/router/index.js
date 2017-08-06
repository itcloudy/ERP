import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
import * as global from './global';
const Login = resolve => require(['../page/global/login'], resolve);
const Home = resolve => require(['../page/global/layout'], resolve);
const notFound = resolve => require(['../page/global/notFound'], resolve);

let routes = [{
        path: '/login',
        name: 'login',
        component: Login
    },

    // 地址管理
    {
        path: '/address',
        component: Home,
        name: 'address',
        children: [
            { path: 'country', component: notFound, name: 'country', hidden: true },
            { path: 'province', component: notFound, name: 'province' },
            { path: 'city', component: notFound, name: 'city' },
            { path: 'district', component: notFound, name: 'district' }
        ]
    },
    {
        path: '/',
        name: 'home',
        component: Home
    },
];
const scrollBehavior = (to, from, savedPosition) => {
    if (savedPosition) {
        // savedPosition is only available for popstate navigations.
        return savedPosition
    } else {
        const position = {}
            // new navigation.
            // scroll to anchor by returning the selector
        if (to.hash) {
            position.selector = to.hash
        }
        // check if any matched route config has meta that requires scrolling to top
        if (to.matched.some(m => m.meta.scrollToTop)) {
            // cords will be used if no selector is provided,
            // or if the selector didn't match any element.
            position.x = 0
            position.y = 0
        }
        // if the returned position is falsy or an empty object,
        // will retain current scroll position.
        return position
    }
}
export default new VueRouter({
    // mode: 'history',
    scrollBehavior,
    base: __dirname,
    routes: routes
});