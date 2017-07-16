<template>
    <div class="left-sidebar" :style="leftSidebarHeight">
        <div class="left-top ">
             golangERP 
        </div>
        <div class="left-main"></div>
    </div>
</template>
<script>
    export default{
        name:"sidebar",
        data(){
            return{
                leftSidebarHeight:{
                    height:"",
                }
            }
        },
        methods:{
            setleftSidebarHeight(){
                this.leftSidebarHeight.height = window.innerHeight +'px' ;
            }
        },
        created(){
            this.setleftSidebarHeight()
        }
    }
</script>
<style lang="scss" scoped>
    .left-sidebar{
        background-color: #528B8B;
        color: white;
        .left-top{
            background-color: #20A0FF;
            height:3rem;
            line-height:3rem;
            font-size:150%;
            text-align:center;
        }
    }
</style>