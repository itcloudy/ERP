<template>
    <el-row>
        <div v-show="leftShow" class="el-col" :class="'el-col-'+leftCol">
            <Sidebar/>
        </div>
        <div class="el-col" :class="'el-col-'+rightCol">
            <Navbar/>
            <Main/>
        </div>
        
    </el-row>
</template>
<script>
    import  {default as Sidebar} from './sidebar';
    import  {default as Navbar} from './navbar';
    import  {default as Main} from './main';
    export default{
        name:"layout",
        components:{
            Sidebar,
            Navbar,
            Main
        },
        data(){
            return {
                leftCol:3,//左侧菜单拦
            }
        },
        computed:{
            // 右侧内容拦
            rightCol:function(){
                return 24 - this.leftCol
            },
            leftShow:function(){
                if (this.leftCol > 0)
                    return true
                else
                    return false
            }
        }
    }
</script>
<style lang="scss" scope>
    
</style>