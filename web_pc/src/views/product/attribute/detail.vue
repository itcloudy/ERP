<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="attributeForm" :model="attributeForm" :inline="true"  class="form-read-only">
                <el-form-item label="属性名称">
                    <span>{{attributeForm.Name}}</span>
                </el-form-item>
                <el-form-item label="属性编码">
                    <span>{{attributeForm.Code}}</span>
                </el-form-item>
                <el-form-item label="创建规格">
                    <span v-if="attributeForm.CreatVariant">是</span>
                    <span v-else>否</span>
                </el-form-item>
            </el-form>
        </div>
        <el-tabs v-if="attributeForm.Values"  v-model="tab_active" type="card" @tab-click="handleClick">
            <el-tab-pane label="属性值" name="attribute_value">
                <el-table
                    ref="caseTable"
                    :data="attributeForm.Values"
                    style="width: 100%">
                    <el-table-column
                        type="selection"
                        width="55">
                    </el-table-column>
                    <el-table-column
                        prop="Name"
                        label="属性值">
                    </el-table-column>
                    <el-table-column label="操作">
                        <template scope="scope">
                            <el-button
                            type="info"
                            size="mini"
                            @click="updateProductAttributeValue(scope.$index, scope.row)">修改</el-button>
                            <el-button
                            type="danger"
                            size="mini"
                            @click="deleteProductAttributeValue(scope.$index, scope.row)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop'; 
    import  {SERVER_PRODUCT_ATTRIBUTE} from '@/server_address';        
    import { mapState } from 'vuex';
    export default {
        data() {
            return {
                loadging:false,
                tab_active:'attribute_value',
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                attributeForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getProductAttributeInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.attributeForm.ID = id;
                this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE+this.attributeForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.attributeForm = data["attribute"];
                            this.access = data["access"];
                        }
                    });
            },
            deleteProductAttributeValue(index,row){

            },
            updateProductAttributeValue(index,row){

            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/attribute");
                }else if ("form"==type){
                    this.$router.push("/product/attribute/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/product/attribute/form/"+this.attributeForm.ID);
            },
        },
        created:function(){
            this.getProductAttributeInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getProductAttributeInfo'
        },
         
    }
</script>