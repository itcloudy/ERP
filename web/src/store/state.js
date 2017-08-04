import localStore from 'utils/local_store';
const state = {
    //登录成功后的用户信息
    userinfo: JSON.parse(localStore.get('userinfo')) || {},
    //后台获得的菜单
    menus: [],

    //记住密码相关信息，现在暂且只做记住一个账号密码
    //后期：每次登录成功一次，就缓存到列表中，然后在登录表单，输入时，会出现下拉列表选择之前登录过得用户
    remumber: {
        remumber_flag: localStore.get('remumber_flag') ? true : false,
        remumber_login_info: localStore.get('remumber_login_info') || {
            username: '',
            token: ''
        }
    },
    //显示左侧菜单栏
    showLeftSidebar: true,
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