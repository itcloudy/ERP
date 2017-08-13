<template>
    <div>
    
        <head-top></head-top>
        <foot-bottom></foot-bottom>
     </div>
</template>
<script>
    import {default as headTop} from '../../common/header'
    import {default as footBottom} from '../../common/footer'
    import {mapState } from 'vuex'
    export default {
    	data(){
            return{
                
            }
        },
        components:{
            headTop,
            footBottom
        }
         
    }
</script>
<style lang="scss" scoped>
     
</style>