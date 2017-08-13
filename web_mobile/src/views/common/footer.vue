<template>
    <section id="bottom-guide">
        <div>footer</div>
    </section>
</template>
<script>
    import {mapState} from 'vuex'
    export default {
    	data(){
            return{
                
            }
        },
         
    }
</script>
<style lang="scss" scoped>
    #bottom-guide{
        background-color: #fff;
        position: fixed;
        z-index: 100;
        bottom:0;
    }
</style>