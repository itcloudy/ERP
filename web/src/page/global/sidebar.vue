<template>
    <div class="left-sidebar" :style="sidebarStyles">
        <div class="left-top">
             golangERP 
        </div>
        <div class="left-main">
        <el-row class="tac">
            <el-col :span="24">
                <!--
                    后台menu的path没有“/”,
                    需要在此增加，router=true,index作为路由，
                    若嵌套需要手动加上上级的path，
                    index要以“/”开头，否则为相对前一个url的地址
                -->
                <el-menu default-active="2" theme="dark" :router="true">
                    <template theme="dark" v-for="(menu,index) in menuList" >
                        <template v-if="menu.children" >
                            <el-submenu index="menu.path" :key="index">
                                <template slot="title"><i :class="menu.Icon"></i>{{menu.name}}</template>
                                <template v-for="(firstMenu,index) in menu.children"   >
                                    <template v-if="firstMenu.children"   >
                                        <el-submenu  >
                                            <template slot="title"><i :class="firstMenu.Icon"></i>{{firstMenu.name}}</template>
                                            <template v-for="(secondMenu,index) in firstMenu.children" >
                                                <el-menu-item v-if="!secondMenu.children" :index="menu.path + '/' + secondMenu.path" :key="secondMenu.index">{{secondMenu.name}}</el-menu-item>
                                            </template>
                                        </el-submenu>
                                    </template>
                                    <template v-if="!firstMenu.children"  >
                                        </i><el-menu-item :index="menu.path + '/' + firstMenu.path"><i :class="firstMenu.Icon"></i>{{firstMenu.name}}</el-menu-item>
                                    </template>
                                </template>
                            </el-submenu>
                        </template>
                        <template v-if="!menu.children" >
                            <el-menu-item :class="menu.Icon" :index="menu.path"><i :class="menu.Icon"></i>{{menu.name}}</el-menu-item>
                        </template>
                    </template>
                </el-menu>
            </el-col>
        </el-row>
  </div>
    </div>
</template>
<script>
    import { mapState } from 'vuex';
    export default{
        name:"sidebar",
        data(){
            return{
                sidebarStyles: this.$store.state.windowStyles.leftSidebarStyles,
                
            }
        },
        computed:{
           ...mapState({
               menuList: state => state.menus
           })
        }
    }
</script>
<style lang="scss" scoped>
    .left-sidebar{
        background-color: rgb(50, 65, 87);
        color: white;
        .left-top{
            background-color: #20A0FF;
            line-height:50px;
            font-size:x-large;
            text-align:center;
        }
    }
</style>