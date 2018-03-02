<template>
    <div id="tree-top" class="block">
        <el-button  icon="document" size="small" @click="changeCreateForm" v-if="Create">创建</el-button>
    </div>
</template>
<script>
    
    export default {

        data() {
            return {
                
            };
        },
        props:["Create"],
        methods:{
            changeCreateForm(){
                this.$emit("changeCreateForm");
            },
           
        }
    }
</script>
<style lang="scss" scope>
    #tree-top{
        float:left;
    }
    
</style>
