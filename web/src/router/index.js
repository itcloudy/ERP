import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import * as global from './global';
const Login = resolve => require(['../page/global/login'], resolve);
let routes = [{
    path: '/login',
    name: 'login',
    component: Login
}]
export default new Router({
    scrollBehavior: () => ({ y: 0 }),
    routes: routes
});