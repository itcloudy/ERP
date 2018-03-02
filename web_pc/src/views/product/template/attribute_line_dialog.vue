<template>
    <el-dialog :title="form.ProductTemplate.Name +'<属性明细>'" :visible="dialogFormVisible" :show-close="false">
        <el-form :model="form" ref="attributeLineDialogForm" :inline="true" :rules="valueFormRules">
            <el-form-item label="属性"  prop="Attribute">
                <el-select
                    v-model="form.Attribute.ID"
                    :name="form.Attribute.Name"
                    filterable
                    remote
                    placeholder="请选择属性"
                    :remote-method="getProductAttributeList">
                    <el-option
                        v-for="item in attributeList"
                        :key="item.ID"
                        :label="item.Name"
                        :value="item.ID">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="属性值"   prop="AttributeValues">
                <el-select
                    v-model="form.AttributeValues"
                    filterable
                    remote
                    multiple
                    placeholder="请选择属性"
                    :selected="selectedValues"
                    :remote-method="getProductAttributeValueList">
                    <el-option
                        v-for="item in attributeValueList"
                        :key="item.ID"
                        :label="item.Name"
                        :value="item.ID">
                    </el-option>
                </el-select>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisibleEn">取 消</el-button>
            <el-button type="primary" @click="saveForm">确 定</el-button>
        </div>
    </el-dialog>
</template>
<script>
import  {SERVER_PRODUCT_ATTRIBUTE,SERVER_PRODUCT_ATTRIBUTE_VALUE,SERVER_PRODUCT_ATTRIBUTE_LINE} from '@/server_address';         
import {validateObjectID,validateList} from '@/utils/validators';
    export default {
        data() {
            return {
                attributeList:[],
                selectedValues:[],
                attributeValueList:[],
                valueFormRules:{
                    Attribute:[
                        { required: true, message: '请选择属性',validator: validateObjectID, trigger: 'blur' }
                    ],
                    AttributeValues:[
                        { required: true, message: '请选择属性值', validator: validateList,trigger: 'blur' }
                    ]
                }
            };
        },
        props:["form","dialogFormVisible"],
        methods:{
            dialogFormVisibleEn(){
                this.$emit("dialogFormVisibleEn");
            },
            saveForm(){
                this.$refs['attributeLineDialogForm'].validate((valid) => {
                    if (valid) {
                        if (this.form.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_ATTRIBUTE_LINE+this.form.ID ,this.form).then(response=>{
                            console.log(JSON.stringify(response));
                             
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_ATTRIBUTE_LINE,this.form).then(response=>{
                            console.log(JSON.stringify(response));
                             
                            });
                        }
                    }
                });
                this.dialogFormVisibleEn();
            },
            getProductAttributeList(query){
                this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.attributeList = data["attributes"];
                    }
                });
            },
            getProductAttributeValueList(query){
                 this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE_VALUE,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.attributeValueList = data["attributeValues"];
                    }
                });
            },
        },
        created:function(){
            this.attributeList = [this.form.Attribute];
            this.AttributeValues = this.form.AttributeValues;
            this.selectedValues = this.form.AttributeValues.map(function(item){return item.Name});
        }
    }
</script>
