import localStore from 'utils/local_store';
const state = {
    //登录成功后的用户信息
    userinfo: JSON.parse(localStore.get('userinfo')) || {},
    //后台获得的菜单
    menus: JSON.parse(localStore.get('menus')) || [],
    // 加载路由完成
    loadRoutersDone: false,

    //显示左侧菜单栏
    showLeftSidebar: true,
    // 窗口高度
    windowHeight: "",
    //窗口样式
    windowStyles: {
        leftSidebarStyles: {
            display: "block",
            height: ""
        },
        rightMainStyles: {
            display: "block",
            height: "",
            backgroundColor: "transparent"
        }
    },
};
export default state