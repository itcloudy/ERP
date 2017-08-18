<template>
    <div id="form-top">
        <el-button  icon="edit" size="small" type="info" v-if="Update">编辑</el-button>
        <el-button  icon="document" size="small" v-if="Create">创建</el-button>
        <el-button-group id="view-type">
            <el-button size="small" @click="change2List" icon="menu"></el-button>
            
        </el-button-group>
    </div>
</template>
<script>
    
    export default {

        data() {
            return {
                
            };
        },
        props:["Create","Update","Unlink","Tree"],
        methods:{
            change2List(){
                this.$emit("changeView","list","");
            }
        }
    }
</script>
<style lang="scss" scope>
    #form-top{
        padding-left: 15px;
        padding-right: 15px;
        #view-type{
            float: right;
        }
        
    }
    
</style>
