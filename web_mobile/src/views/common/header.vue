<template>
    <header id="top-header">
        <section class="title-head" v-if="headTitle">
            <span class="title-text">{{headTitle}}</span>
        </section>
    </header>
</template>

<script>
    import {mapState} from 'vuex'
    export default {
    	data(){
            return{
                
            }
        },
        props: ['headTitle'],
         
    }
</script>
<style lang="scss" scoped>
    @import '../../styles/mixin';
    #top-header{
        background-color: $headBgColor;
        position: fixed;
        z-index: 100;
        top: 0;
        @include widthHeight(100%, 2rem);
    
        .title-head{
            @include allCenter;
            line-height:2rem;
            width: 50%;
            color: #fff;
            text-align: center;
            .title-text{
                @include fontSizeColor(1rem, #fff);
                text-align: center;
                font-weight: bold;
            }
        }
    }
</style>