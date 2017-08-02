<template>
    <div  class="navbar-container">
        
        <span  @click="showLeftSidebarClick"><img src="/static/images/menu.png" alt="菜单显示切换"></span>
        <div class="navbar-right">
            <el-dropdown trigger="click">
                <el-button type="primary" class="el-dropdown-link">
                    系统管理员<i class="el-icon-caret-bottom el-icon--right"></i>
                </el-button>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item>个人信息</el-dropdown-item>
                    <el-dropdown-item>注销</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>
<script>
    import { mapState } from 'vuex'
    export default{
        name:"navbar",
        data(){
            return{
                
            }
        },
        methods:{
            showLeftSidebarClick:function(){
                 this.$store.dispatch('GLOBAL_TOGGLE_LEFT_SIDEBAR');
            }
        },
        computed:{
           
        },
    }
</script>
<style lang="scss" scoped>
    .navbar-container{
        background-color: #20A0FF;
        color: white;
        height:50px;
        line-height:50px;
        .navbar-right{
            display:inline-block;
            float: right;
        }
    }
</style>