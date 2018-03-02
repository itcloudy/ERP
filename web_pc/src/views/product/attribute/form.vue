<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="attributeForm" :rules="attributeFormRules"  :model="attributeForm" label-width="80px">
                <el-form-item label="属性名称" prop="Name">
                    <el-input v-model="attributeForm.Name"></el-input>
                </el-form-item>
                <el-form-item label="属性编码" prop="Code">
                    <el-input v-model="attributeForm.Code"></el-input>
                </el-form-item>
                <el-form-item label="创建规格">
                    <el-switch on-text="是" off-text="否" v-model="attributeForm.CreatVariant"></el-switch>
                </el-form-item>
            </el-form>
        </div>
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
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                attributeForm:{},
                NewProductAttributeForm:{
                    Name:"",
                    ID:"",
                    Code:"",
                    CreatVariant:true,
                },
               attributeFormRules:{
                    Name:[
                        { required: true, message: '请输入属性名称', trigger: 'blur' }
                    ],
                    Code:[
                        { required: true, message: '请输入属性编码', trigger: 'blur' }
                    ],
                     
                }
                
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                this.$refs['attributeForm'].validate((valid) => {
                    if (valid) {
                        if (this.attributeForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_ATTRIBUTE+this.attributeForm.ID ,this.attributeForm).then(response=>{
                                let {code,msg,attributeID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/attribute/detail/"+attributeID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_ATTRIBUTE,this.attributeForm).then(response=>{
                                let {code,msg,attributeID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/attribute/detail/"+attributeID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                });
            },
            getProductAttributeInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.attributeForm.ID = id;
                    this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE+this.attributeForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.attributeForm = data["attribute"];
                                this.provinceList = [this.attributeForm.Province]
                                this.countryList = [this.attributeForm.Country]
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.attributeForm = this.NewProductAttributeForm;
                }
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/attribute");
                }else if ("form"==type){
                    this.$router.push("/product/attribute/form/"+id);
                }
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