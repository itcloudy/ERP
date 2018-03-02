<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="valueForm" :model="valueForm" :rules="valueFormRules" label-width="80px">
                <el-form-item label="属性" prop="Attribute">
                    <el-select
                        v-model="valueForm.Attribute.ID"
                        :name="valueForm.Attribute.Name"
                        filterable
                        remote
                        placeholder="请输入属性"
                        :remote-method="getAttributeList">
                        <el-option
                            v-for="item in attributeList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="属性值" prop="Name">
                    <el-input v-model="valueForm.Name"></el-input>
                </el-form-item>
        
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';    
    import  {SERVER_PRODUCT_ATTRIBUTE,SERVER_PRODUCT_ATTRIBUTE_VALUE} from '@/server_address';             
    import { mapState } from 'vuex';
    import {validateObjectID} from '@/utils/validators';
    export default {
        data() {
            return {
                loadging:false,
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                valueForm:{},
                NewValueForm:{
                    Name:"",
                    ID:"",
                    Attribute:{
                        Name:"",
                        ID:"",
                    },
                },
                attributeList:[],
                valueFormRules:{
                    Name:[
                        { required: true, message: '请输入属性值名称', trigger: 'blur' }
                    ],
                    Attribute:[
                        { required: true, message: '请选择属性',validator: validateObjectID, trigger: 'blur' }
                    ],
                     
                }
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                this.$refs['valueForm'].validate((valid) => {
                    if (valid) {
                        if (this.valueForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_ATTRIBUTE_VALUE+this.valueForm.ID ,this.valueForm).then(response=>{
                                let {code,msg,attributeValueID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/attributevalue/detail/"+attributeValueID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_ATTRIBUTE_VALUE,this.valueForm).then(response=>{
                                let {code,msg,attributeValueID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/attributevalue/detail/"+attributeValueID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                });
            },
            getAttributeValueInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.valueForm.ID = id;
                    this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE_VALUE+this.valueForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.valueForm = data["attributeValue"];
                                this.attributeList = [this.valueForm.Attribute];
                            }
                        });
                }else{
                    this.valueForm = this.NewValueForm;
                }
            },
            getAttributeList(query){
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
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/attributevalue");
                }else if ("form"==type){
                    this.$router.push("/product/attributevalue/form/"+id);
                }
            },
        },
        created:function(){
            this.getAttributeValueInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getAttributeValueInfo'
        },
         
    }
</script>