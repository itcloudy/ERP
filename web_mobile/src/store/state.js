import localStore from 'utils/local_store';
const state = {
    //登录成功后的用户信息
    userinfo: JSON.parse(localStore.get('userinfo')) || {},
    //后台获得的菜单(后台菜单)
    backgroundMenus: JSON.parse(localStore.get('backgroundMenus')) || {},
    // 加载路由完成
    loadRoutersDone: false,
    //可以访问后台
    isBackgroundUser: false,

    // 多单位支持
    multiUnit: false,
};
export default state