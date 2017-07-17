<template>
    <div class="right-main" :style="mainContentStyles">
        <div class="main-container" >
            main
        </div>
    </div>
</template>
<script>
    export default{
        name:"main",
        data(){
            return{
                mainContentStyles: this.$store.state.windowStyles.rightMainStyles
            }
        },
    }
</script>
<style lang="scss" scoped>
    .right-main{
        background-color: rgb(240,238,238);
        .main-container{
            padding:0.5rem;
        }
    }
    
</style>