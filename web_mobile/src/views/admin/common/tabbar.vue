<template>
    <mt-tabbar v-model="tabSelected">
        <mt-tab-item id="admin-menu"  >
            后台菜单
        </mt-tab-item>
        <mt-tab-item id="订单">
            订单
        </mt-tab-item>
        <mt-tab-item id="发现">
            发现
        </mt-tab-item>
        <mt-tab-item id="我的">
            我的
        </mt-tab-item>
    </mt-tabbar>

</template>
<script>
    import { Toast } from 'mint-ui';
    export default {
    data() {
        return {
            tabSelected: '',
        };
    },
    methods:{
        goPage(path){
            this.$router.push(path);
        }
    },
    watch: {
        tabSelected: function (val, oldVal) {
            console.log(val)
        }
    
    }
    }
</script>
<style lang="scss" scope>
</style>
