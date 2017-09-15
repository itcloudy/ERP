<template>

    <div  class="navbar-container">
        <el-menu  :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :router="true" >
            <el-row type="flex" justify="end">
                <el-menu-item index="home" index="/">首页</el-menu-item>
                <el-menu-item index="work_space" index="/work">我的工作台</el-menu-item>
                <el-submenu index="user_info"  >
                    <template slot="title">{{userinfo.NameZh}}</template>
                    <el-menu-item v-if="userinfo.IsBackground" index="/admin">进入后台</el-menu-item>
                    <el-menu-item v-if="userinfo.Name" >个人信息</el-menu-item>
                    <el-menu-item v-if="userinfo.Name" index="/admin/login">注销</el-menu-item>
                    <el-menu-item v-else index="/admin/login">登录</el-menu-item>
                </el-submenu>
                
            </el-row>
        </el-menu> 
        
    </div>
</template>
<script>
    import { mapState } from 'vuex';
    import localStore from '@/utils/local_store';
    export default{
        name:"navbar",
        data(){
            return{
                activeIndex: 'home',
                 
            }
        },
        methods:{
            handleSelect(key, keyPath) {
                console.log(key, keyPath);
            },                                                                
            showLeftSidebarClick:function(){
                 this.$store.dispatch('GLOBAL_TOGGLE_LEFT_SIDEBAR');
            },
        },
        computed:{
           ...mapState({
               userinfo: state => state.userinfo
           })
        }
    }
</script>
<style lang="scss" scoped>
     
</style>