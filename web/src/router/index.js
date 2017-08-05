import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)
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
            { path: '/country', component: notFound, name: 'country', hidden: true },
            { path: '/province', component: notFound, name: 'province' },
            { path: '/city', component: notFound, name: 'city' },
            { path: '/district', component: notFound, name: 'district' },
        ]
    },
    {
        path: '/',
        name: 'home',
        component: Home
    },
]
export default new Router({
    scrollBehavior: () => ({ y: 0 }),
    routes: routes
});