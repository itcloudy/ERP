import Vue from 'vue'
import Router from 'vue-router'
const Login = resolve => require(['@/views/common/login'], resolve);
const Home = resolve => require(['@/views/common/home'], resolve);
Vue.use(Router)

let routes = [{
        path: "/",
        name: "Home",
        component: Home,
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    },


];

export default new Router({
    routes: routes
});