<template>
    <el-row>
        <el-col :span="3">
            <Sidebar/>
        </el-col>
        <el-col :span="3">
            <Navbar/>
            <Main/>
        </el-col>
        
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
        }
    }
</script>
<style lang="scss" scope>
    
</style>