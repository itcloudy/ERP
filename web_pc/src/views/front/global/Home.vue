<template>
    <div>
        <frontHeader/>
        <frontMain/>
        <frontFooter/>
    </div>
</template>
<script>
    import  {default as frontFooter} from './footer';
    import  {default as frontHeader} from './header';
    import  {default as frontMain} from './main';
    export default{
        name:"frontHome",
        components:{
            frontFooter,
            frontHeader,
            frontMain,
        },
        data(){
            return {
            }
        } 
    }
</script>
<style lang="scss" scope>
 
    
</style>