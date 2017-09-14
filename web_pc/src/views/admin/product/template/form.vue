<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="templateForm" :model="templateForm" label-width="80px">
                 
                <el-form-item label="款式名称">
                    <el-input v-model="templateForm.Name"></el-input>
                </el-form-item>
        
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/admin/common/FormTop'; 
    import  {SERVER_PRODUCT_TEMPLATE} from '@/server_address';         
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
                templateForm:{
                    Name:"",
                    ID:"",
                    
                },
                NewTemplateForm:{
                    Name:"",
                    ID:"",
                    
                },
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                if (this.templateForm.ID >0){
                    this.$ajax.put(SERVER_PRODUCT_TEMPLATE+this.templateForm.ID ,this.templateForm).then(response=>{
                        let {code,msg,templateID} = response.data;
                        if(code=='success'){
                            this.$message({ message:msg, type: 'success' });
                            this.$router.push("/admin/product/template/detail/"+templateID);
                        }else{
                            this.$message({ message:msg, type: 'error' });
                        }
                    });
                }else{
                    this.$ajax.post(SERVER_PRODUCT_TEMPLATE,this.templateForm).then(response=>{
                        let {code,msg,templateID} = response.data;
                        if(code=='success'){
                            this.$message({ message:msg, type: 'success' });
                            this.$router.push("/admin/product/template/detail/"+templateID);
                        }else{
                            this.$message({ message:msg, type: 'error' });
                        }
                    });
                }
            },
            getCityInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.templateForm.ID = id;
                    this.$ajax.get(SERVER_PRODUCT_TEMPLATE+this.templateForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.templateForm = data["template"];
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.templateForm = this.NewTemplateForm;
                }
            },
             
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/admin/product/template");
                }else if ("form"==type){
                    this.$router.push("/admin/product/template/form/"+id);
                }
            },
        },
        created:function(){
            this.getCityInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getCityInfo'
        },
         
    }
</script>