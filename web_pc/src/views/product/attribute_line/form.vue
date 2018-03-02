<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="lineForm" :rules="lineFormRules" :model="lineForm" label-width="80px">
                <el-form-item label="属性" prop="Template">
                    <el-select
                        v-model="lineForm.ProductTemplate.ID"
                        :name="lineForm.ProductTemplate.Name"
                        filterable
                        remote
                        placeholder="请输入款式"
                        :remote-method="getTemplateList">
                        <el-option
                            v-for="item in templateList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="属性" prop="Attribute">
                    <el-select
                        v-model="lineForm.Attribute.ID"
                        :name="lineForm.Attribute.Name"
                        filterable
                        remote
                        multiple 
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
                <el-form-item label="属性值" prop="AttributeValues">
                    <el-select
                        v-model="lineForm.AttributeValues"
                        :name="lineForm.AttributeValues"
                        filterable
                        remote
                        placeholder="请输入属性值"
                        :remote-method="getAttributeValueList">
                        <el-option
                            v-for="item in attributeValueList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                 
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';   
    import  {SERVER_PRODUCT_TEMPLATE,SERVER_PRODUCT_ATTRIBUTE,SERVER_PRODUCT_ATTRIBUTE_VALUE,SERVER_PRODUCT_ATTRIBUTE_LINE} from '@/server_address';      
    import { mapState } from 'vuex';
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
                lineForm:{ },
                NewAttributeLineForm:{
                    Name:"",
                    ID:"",  
                    Attribute:{
                        Name:"",
                        ID:"",
                    } ,
                    ProductTemplate:{
                        Name:"",
                        ID:"",
                    } ,
                    AttributeValues:[],
                },
                templateList:[],
                attributeList:[],
                attributeValueList:[],
                lineFormRules:{
                    Name:[
                        { required: true, message: '请输入国家名称', trigger: 'blur' }
                    ] 
                }
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                this.$refs['lineForm'].validate((valid) => {
                    if (valid) {
                        if (this.lineForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_ATTRIBUTE_LINE+this.lineForm.ID,this.lineForm).then(response=>{
                                let {code,msg,lineID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/attribute/line/detail/"+lineID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_ATTRIBUTE_LINE,this.lineForm).then(response=>{
                                let {code,msg,lineID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/attribute/line/detail/"+lineID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                 });
            },
            getAttributeLineInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.lineForm.ID = id;
                    this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE_LINE+this.lineForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.lineForm = data["line"];
                                console.log(JSON.stringify(this.lineForm));
                                this.attributeList = [this.lineForm.Attribute];
                                this.templateList = [this.lineForm.ProductTemplate];
                                console.log(JSON.stringify(this.lineForm.AttributeValues));
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.lineForm = this.NewAttributeLineForm;
                }
            },
            
            getTemplateList(query){
                this.$ajax.get(SERVER_PRODUCT_TEMPLATE,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.templateList = data["templates"];
                    }
                });
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
            getAttributeValueList(query){
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
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/attribute/line");
                }else if ("form"==type){
                    this.$router.push("/product/attribute/line/form/"+id);
                }
            },
        },
        created:function(){
            this.getAttributeLineInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getAttributeLineInfo'
        },
         
    }
</script>