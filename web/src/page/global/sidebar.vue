<template>
    <div class="left-sidebar" :style="sidebarStyles">
        <div class="left-top">
             golangERP 
        </div>
        <div class="left-main">
        <el-row class="tac">
            <el-col :span="24">
                <el-menu default-active="2" theme="dark" :router="true">
                    <template theme="dark" v-for="(menu,index) in menuList" :>
                        <template v-if="menu.children" >
                            <el-submenu index="menu.Path" :key="index">
                                <template slot="title"><i :class="menu.Icon"></i>{{menu.Name}}</template>
                                <template v-for="(firstMenu,index) in menu.children"   >
                                    <template v-if="firstMenu.children"   >
                                        <el-submenu  >
                                            <template slot="title"><i :class="firstMenu.Icon"></i>{{firstMenu.Name}}</template>
                                            <template v-for="(secondMenu,index) in firstMenu.children" >
                                                <el-menu-item v-if="!secondMenu.children" :index="'/' + menu.Path + '/' + secondMenu.Path" :key="secondMenu.index">{{secondMenu.Name}}</el-menu-item>
                                            </template>
                                        </el-submenu>
                                    </template>
                                    <template v-if="!firstMenu.children"  >
                                        </i><el-menu-item :index="'/' + menu.Path + '/' + firstMenu.Path"><i :class="firstMenu.Icon"></i>{{firstMenu.Name}}</el-menu-item>
                                    </template>
                                </template>
                            </el-submenu>
                        </template>
                        <template v-if="!menu.children" >
                            <el-menu-item :class="menu.Icon" :index="'/' + menu.Path"><i :class="menu.Icon"></i>{{menu.Name}}</el-menu-item>
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