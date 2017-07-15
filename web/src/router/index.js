import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import * as global from './global';
const Login = resolve => require(['../page/global/login'], resolve);
const Home = resolve => require(['../page/global/layout'], resolve);

let routes = [{
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/',
        name: 'home',
        component: Home
    }
]
export default new Router({
    scrollBehavior: () => ({ y: 0 }),
    routes: routes
});