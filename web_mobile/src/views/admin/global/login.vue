<template>
    <div >
        <mt-header fixed title="登录"></mt-header>
        <div id="login">
            <mt-field  placeholder="请输入用户名" v-model="data.username"></mt-field>
            <mt-field  placeholder="请输入密码" type="password" v-model="data.password"></mt-field>
            <mt-button type="primary">登录</mt-button>
        </div>
    </div>
</template>
<script>
    import localStore from '@/utils/local_store';
    import lazyLoadMenusRoutes from '@/utils/lazyload';
    import { mapMutations } from 'vuex';
    import { mapState } from 'vuex';

    export default {
        name: 'login',
        data() {
            return {
                logining:false,
                winSize: {
                    width: '',
                    height: ''
                },
                formOffset: {
                    position: 'absolute',
                    left: '',
                    top: ''
                },
                login_actions: {
                    disabled: false
                },
                data: {
                    username: 'admin',
                    password: 'cloudy',
                    token: ''
                },
                rule_data:{
                    username: [{
                        validator:(rule, value, callback)=>{
                            if (value === '') {
                                callback(new Error('请输入用户名'));
                            } else {
                                if(/^[a-zA-Z0-9_-]{2,16}$/.test(value)){
                                    callback();
                                }else{
                                    callback(new Error('用户名至少6位,由大小写字母和数字,-,_,组成'));
                                }
                            }
                        },
                        trigger: 'blur'
                    }],
                    password: [{
                        validator:(rule, value, callback)=>{
                            if (value === '') {
                                callback(new Error('请输入密码'));
                            } else {
                                if(!(/^[a-zA-Z0-9_-]{6,16}$/.test(value))){
                                    callback(new Error('密码至少6位,由大小写字母和数字,-,_组成'));
                                }else{
                                    callback();
                                }

                            }
                        },
                        trigger: 'blur'
                    }]
                },
            }
        },
        computed:{
           ...mapState({
               loadRoutersDone: state => state.loadRoutersDone,
               backgroundMenus: state => state.backgroundMenus
            })
        },
        created() {
            this.setSize();
        },
        methods: {
            // 设置窗口大小
            setSize() {
                this.winSize.width = window.innerWidth + 'px';
                this.winSize.height = window.innerHeight +'px' ;
                this.formOffset.left = (parseInt(this.winSize.width) / 2 - 175) + 'px';
                this.formOffset.top = (parseInt(this.winSize.height) / 2 - 178) + 'px';
            },
            onLogin(ref){
                this.$refs[ref].validate((valid)=>{
                    if(valid){
                        this.logining = true;
                        let params = {
                            username: this.data.username,
                            password: this.data.password
                        };
                        this.$ajax.post('/login',params).then(response=>{
                            this.logining = false;
                            let {code,msg,data} = response.data;
                            if (code=='success'){
                                let user = data.user;
                                //提示
                                this.$message({ message:msg, type: 'success' });
                                // 本地缓存用户信息
                                localStore.set('userinfo',JSON.stringify(user));
                                //更新store中的userinfo
                                this.setGlobalUserInfo(user)
                                // 本地缓存权限信息
                                localStore.set('groups',JSON.stringify(data.groups));
                                // 验证通过，获得菜单
                                let params = {
                                    groups:data.groups,
                                    isAdmin:user.IsAdmin,
                                }
                                this.$ajax.post("/menu",params).then(response=>{
                                    let {code,msg,data} = response.data;
                                    if(code=='success'){
                                        // 后台菜单
                                        let backgroundMenus = this.menuList2Json(data.menus);
                                        // 本地缓存后台菜单信息
                                        localStore.set('backgroundMenus',JSON.stringify(backgroundMenus));
                                        this.setGlobalUserMenu(backgroundMenus);
                                        this.loadBackgroundRouters();
                                    }else{
                                        this.$message({  message:msg,   type: 'error' });
                                    }
                                    //登录成功跳转到后台首页
                                    this.$router.push('/');
                                });
                            }else{
                                this.$message({  message:msg,   type: 'error' });
                            }
                        });
                    }
                });
            },
            menuList2Json(menuList){
                // 后台
                let resultJson =[];
               
                let stepList = [];
                if (menuList== null){
                    return resultJson;
                }
                let  menuLen = menuList.length;
                // 获得所有的步长,以及顶级才按
                for( let i=0;i<menuLen;i++){
                    let menu  = menuList[i];
                    let step = menu.ParentRight - menu.ParentLeft;
                    let hasStep = false;
                    for(let j=0;j<stepList.length;j++){
                        if (step==stepList[j]){
                            hasStep = true;
                        }
                    }
                    if (hasStep==false){
                        stepList.push(step);
                    }
                }
                // 对stepList排序
                stepList.sort();
                //循环处理menu,后期考虑删除已经处理的元素
                for(let j=0,len=stepList.length;j<len;j++){
                    let step = stepList[j];
                    for(let i=0;i<menuLen;i++){
                        let menu = menuList[i];
                        let menuStep = menu.ParentRight - menu.ParentLeft;
                        // 若不相等跳过
                        if (step!=menuStep){
                            continue
                        }
                        //排除顶级菜单
                        if (menu.Parent != null){
                            let parentIndex = menu.Parent.Index;
                            for(let k=0; k<menuLen;k++){
                                if (menuList[k].Index == parentIndex){
                                    if(!("children" in menuList[k])){
                                        menuList[k].children = [];
                                    } 
                                    delete menu.Parent;
                                    delete menu.ParentLeft;
                                    delete menu.ParentRight;
                                    // 增加path前缀，即地址为绝对地址
                                    menu.path = '/admin/' +menu.path;
                                    menuList[k].children.push(menu);
                                }
                            }
                        }else{
                            delete menu.Parent;
                            delete menu.ParentLeft;
                            delete menu.ParentRight;
                            // 增加path前缀，即地址为绝对地址
                            menu.path = '/admin/' +menu.path;
                            resultJson.push(menu)
                        }
                    }
                }
                 
                return resultJson;
            },
            loadBackgroundRouters(){
                if (this.loadRoutersDone==false){
                    let menusRoutes = lazyLoadMenusRoutes(this.backgroundMenus);
                    this.$router.addRoutes([menusRoutes]);
                    //动态加载路由
                    this.setloadRoutersDone(true);
                }
            },
            ...mapMutations({
                setGlobalUserInfo: "GLOBAL_SET_USERINFO",
                setGlobalUserMenu: "GLOBAL_SET_UER_MENUS",
                setloadRoutersDone: "GLOBAL_LOAD_ROUTES_DONE"
            })
        }
        
        
    }
</script>
<style lang="scss" scoped>
    #login{
        margin-top: 3rem;
    }
</style>
