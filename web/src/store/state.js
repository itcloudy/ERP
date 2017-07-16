import localStore from 'utils/local_store';
const store = {
    //登录成功后的用户信息
    userinfo: localStore.get('userinfo') || {},

    //记住密码相关信息，现在暂且只做记住一个账号密码
    //后期：每次登录成功一次，就缓存到列表中，然后在登录表单，输入时，会出现下拉列表选择之前登录过得用户
    remumber: {
        remumber_flag: localStore.get('remumber_flag') ? true : false,
        remumber_login_info: localStore.get('remumber_login_info') || {
            username: '',
            token: ''
        }
    },
};
export default store