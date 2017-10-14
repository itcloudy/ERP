<template>
    <div class="block">
        <mt-header fixed title="golangERP登录"></mt-header>
        <section id="login-section">
          <mt-field label="用户名" placeholder="请输入用户名" v-model="user.username"></mt-field>
          <mt-field label="密码" placeholder="请输入密码" type="password" v-model="user.password"></mt-field>
          <mt-button type="primary" id="btn-login"  @click="onLogin" size="large">登录</mt-button>
        </section>
    </div>
</template>
<script>
    import { Toast } from 'mint-ui';
    import localStore from '@/utils/local_store';
    import lazyLoadMenusRoutes from '@/utils/lazyload';
    import { mapMutations } from 'vuex';
    import { mapState } from 'vuex';
    export default {
    data() {
        return {
            user:{
                username:"admin",
                password:"cloudy"
            }
        };
    },
    computed:{
        ...mapState({
            loadRoutersDone: state => state.loadRoutersDone,
            backgroundMenus: state => state.backgroundMenus
        })
    },
    methods:{
      onLogin(){
         this.$ajax.post('/login',this.user).then(response=>{
            let {code,msg,data} = response.data;
            if (code=='success'){
                Toast({message:msg,duration: 500});
                let user = data.user;
                //提示
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
                        Toast({message:msg,duration: 500});
                    }
                    //登录成功跳转到后台首页
                    this.$router.push('/admin/menu');
                });
            }else{
                Toast({message:"登录失败",duration: 500});
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
<style lang="scss" scope>
    #login-section{
        margin-top: 5rem;
    }
</style>
