<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="countryForm" :rules="countryFormRules" :model="countryForm" label-width="80px">
                <el-form-item label="国家名称" prop="Name">
                    <el-input v-model="countryForm.Name"></el-input>
                </el-form-item>
        
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';   
    import  {SERVER_ADDRESS_COUNTRY} from '@/server_address';      
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
                countryForm:{ },
                NewCountryForm:{
                    Name:"",
                    ID:"",   
                },
                countryFormRules:{
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
                this.$refs['countryForm'].validate((valid) => {
                    if (valid) {
                        if (this.countryForm.ID >0){
                            this.$ajax.put(SERVER_ADDRESS_COUNTRY+this.countryForm.ID,this.countryForm).then(response=>{
                                let {code,msg,countryID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/address/country/detail/"+countryID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_ADDRESS_COUNTRY,this.countryForm).then(response=>{
                                let {code,msg,countryID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/address/country/detail/"+countryID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                 });
            },
            getCountryInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.countryForm.ID = id;
                    this.$ajax.get(SERVER_ADDRESS_COUNTRY+this.countryForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.countryForm = data["country"];
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.countryForm = this.NewCountryForm;
                }
            },
            
            
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/address/country");
                }else if ("form"==type){
                    this.$router.push("/address/country/form/"+id);
                }
            },
        },
        created:function(){
            this.getCountryInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getCountryInfo'
        },
         
    }
</script>