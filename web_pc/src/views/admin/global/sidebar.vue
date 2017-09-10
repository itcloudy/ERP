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
                    <el-menu theme="dark" :router="true"  v-for="(menu,index) in backgroundMenus" :key="index">
                        <el-submenu   v-if="menu.expand"  index="menu.path">
                            <template slot="title"><i :class="menu.Icon"></i>{{menu.name}}</template>
                            <template v-for="(firstMenu,index) in menu.children">
                                <template v-if="firstMenu.expand">
                                    <el-submenu :key="index">
                                        <template slot="title"><i :class="firstMenu.Icon"></i>{{firstMenu.name}}</template>
                                        <template v-for="(secondMenu,index) in firstMenu.children" >
                                            <el-menu-item v-if="!secondMenu.expand" :index="secondMenu.path" :key="secondMenu.index">{{secondMenu.name}}</el-menu-item>
                                        </template>
                                    </el-submenu>
                                </template>
                                <template v-if="!firstMenu.expand"  >
                                    <el-menu-item :index="firstMenu.path" :key="index"><i :class="firstMenu.Icon"></i>{{firstMenu.name}}</el-menu-item>
                                </template>
                            </template>
                        </el-submenu>
                        <el-menu-item  v-if="!menu.expand" :class="menu.Icon" :index="menu.path"><i :class="menu.Icon"></i>{{menu.name}}</el-menu-item>
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
               backgroundMenus: state => state.backgroundMenus,
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