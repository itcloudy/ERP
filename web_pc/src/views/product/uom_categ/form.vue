<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="uomcategForm" :model="uomcategForm" :rules="uomcategFormRules" label-width="80px">
                <el-form-item label="单位类别名称" prop="Name">
                    <el-input v-model="uomcategForm.Name"></el-input>
                </el-form-item>
        
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';         
    import  {SERVER_PRODUCT_UOM_CATEG} from '@/server_address';
    import {validateObjectID} from '@/utils/validators';
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
                uomcategForm:{ },
                NewUomCategForm:{
                    Name:"",
                    ID:"",
                },
                uomcategFormRules:{
                    Name:[
                        { required: true, message: '请输入单位类别名称', trigger: 'blur' }
                    ]
                }
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                this.$refs['uomcategForm'].validate((valid) => {
                    if (valid) {
                        if (this.uomcategForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_UOM_CATEG+this.uomcategForm.ID ,this.uomcategForm).then(response=>{
                                let {code,msg,uomcategID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/uomcateg/detail/"+uomcategID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_UOM_CATEG,this.uomcategForm).then(response=>{
                                let {code,msg,uomcategID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/uomcateg/detail/"+uomcategID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }  
                    } 
                });
            },
            getUomCategInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.uomcategForm.ID = id;
                    this.$ajax.get(SERVER_PRODUCT_UOM_CATEG+this.uomcategForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.uomcategForm = data["uomcateg"];
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.uomcategForm = this.NewUomCategForm;
                }
            },
             
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/uomcateg");
                }else if ("form"==type){
                    this.$router.push("/product/uomcateg/form/"+id);
                }
            },
        },
        created:function(){
            this.getUomCategInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getUomCategInfo'
        },
         
    }
</script>