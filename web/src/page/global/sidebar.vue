<template>
    <div class="left-sidebar" :style="sidebarStyles">
        <div class="left-top">
             golangERP 
        </div>
        <div class="left-main">
        <el-row class="tac">
            <el-col :span="24">
                <el-menu default-active="2" theme="dark">
                <el-submenu index="1">
                    <template slot="title"><i class="el-icon-message"></i>导航一</template>
                    <el-menu-item-group>
                        <template slot="title">分组一</template>
                        <el-menu-item index="1-1">选项1</el-menu-item>
                        <el-menu-item index="1-2">选项2</el-menu-item>
                    </el-menu-item-group>
                    <el-menu-item-group title="分组2">
                        <el-menu-item index="1-3">选项3</el-menu-item>
                    </el-menu-item-group>
                    <el-submenu index="1-4">
                        <template slot="title">选项4</template>
                        <el-menu-item index="1-4-1">选项1</el-menu-item>
                    </el-submenu>
                </el-submenu>
                <el-menu-item index="2"><i class="el-icon-menu"></i>导航二</el-menu-item>
                <el-menu-item index="3"><i class="el-icon-setting"></i>导航三</el-menu-item>
                </el-menu>
            </el-col>
        </el-row>
  </div>
    </div>
</template>
<script>
    export default{
        name:"sidebar",
        data(){
            return{
                sidebarStyles: this.$store.state.windowStyles.leftSidebarStyles,
                
            }
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