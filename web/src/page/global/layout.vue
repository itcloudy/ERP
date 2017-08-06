<template>
    <el-row>
        <div v-show="showLeftSidebar" class="el-col" :class="'el-col-'+splitCol.leftCol">
            <Sidebar/>
        </div>
        <div class="el-col left-sidebar-container" :class="'el-col-'+splitCol.rightCol">
            <transition name="sidebar">
                <Navbar/>
            </transition>
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
                // showLeftSidebar: this.$store.state.showLeftSidebar,
            }
        },
        computed:{
            // 右侧内容拦
            splitCol:function(){
                if (this.showLeftSidebar){
                    return {
                        leftCol:3,
                        rightCol:21,
                    }
                }else{
                     return {
                        leftCol:0,
                        rightCol:24,
                    }
                }
            },
            showLeftSidebar:function(){
                return  this.$store.state.showLeftSidebar
            }
        },
        methods:{
            setWindowHeight(){
                let height = window.innerHeight ;
                this.$store.dispatch("GLOBAL_SET_WINDOW_HRIGHT",{height:height});
            }
        },
        created(){
            this.setWindowHeight()
        },
    }
</script>
<style lang="scss" scope>
    .left-sidebar-container{
        .sidebar-enter-active, .sidebar-leave-active {
            transition: opacity .5s
        }
        .sidebar-enter, .sidebar-leave-to  {
            opacity: 0
        }
    }
</style>