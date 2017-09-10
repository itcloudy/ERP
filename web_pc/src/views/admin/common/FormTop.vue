<template>
    <div id="form-top">
        <el-button  icon="edit" size="small" type="info" v-if="Update && !edit"  @click="formEdit">编辑</el-button>
        <el-button  icon="edit" size="small" type="info" v-if="edit"  @click="formSave">保存</el-button>
        <el-button  icon="document" size="small" @click="change2Form" v-if="Create && !edit">创建</el-button>
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
        props:["Create","Update","Unlink","edit"],
        methods:{
            change2List(){
                this.$emit("changeView","list","");
            },
            change2Form(){
                this.$emit("changeView","form","new");
            },
            formEdit(){
                 this.$emit("formEdit");
            },
            formSave(){
                this.$emit("formSave");
            }
        }
    }
</script>
<style lang="scss" scope>
    #form-top{
        padding-bottom: 5px;
        border-bottom-color: rgb(170, 170, 170);
        border-bottom-style: solid;
        border-bottom-width: 1px;
        margin-bottom: 10px;
        #view-type{
            float: right;
        }
        
    }
    
</style>
