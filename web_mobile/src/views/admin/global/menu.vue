<template>
    <div class="block">
        <mt-header title="golangERP" fixed>
            <router-link to="/" slot="left">
                <mt-button icon="back" @click="goBack">返回</mt-button>
            </router-link>
        </mt-header>
        <section id="content-section">
            <template v-for="(topMenu,index) in backgroundMenus">
                <div :key="index" class="top-menus">
                    <div>
                        <mt-button type="danger" disabled size="large">{{topMenu.name}}</mt-button>
                    </div>
                    <template v-for="(menu,index) in topMenu.children">
                        <mt-button :key="index" class="menu" @click="goPage(menu.path)">{{menu.name}}</mt-button>
                    </template>
                </div>
            </template>
        </section>
    </div>
</template>
<script>
    import { Toast } from 'mint-ui';
    import { mapMutations } from 'vuex';
    import { mapState } from 'vuex';
    export default {
    data() {
        return {
        };
    },
    computed:{
        ...mapState({
            backgroundMenus: state => state.backgroundMenus,
        })
    },
    methods:{
        goBack(){
            this.$router.go(-1);
        },
        goPage(path){
            Toast(path);
            this.$router.push(path);
        }
    }
    
  }
</script>
<style lang="scss" scope>
</style>

