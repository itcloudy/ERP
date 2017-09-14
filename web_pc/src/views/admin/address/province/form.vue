<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="provinceForm" :model="provinceForm" label-width="80px">
                <el-form-item label="所属国家">
                    <el-select
                        v-model="provinceForm.Country.ID"
                        :name="provinceForm.Country.Name"
                        filterable
                        remote
                        placeholder="请输入国家"
                        :remote-method="getCountryList">
                        <el-option
                            v-for="item in countryList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="省份名称">
                    <el-input v-model="provinceForm.Name"></el-input>
                </el-form-item>
        
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/admin/common/FormTop';      
    import  {SERVER_ADDRESS_COUNTRY,SERVER_ADDRESS_PROVINCE} from '@/server_address';           
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
                provinceForm:{
                    Name:"",
                    ID:0,
                    Country:{
                        Name:"",
                        ID:0,
                    }
                },
                NewProvinceForm:{
                    Name:"",
                    ID:0,
                    
                    Country:{
                        Name:"",
                        ID:0,
                    }
                },
                countryList:[],
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                if (this.provinceForm.ID >0){
                    this.$ajax.put(SERVER_ADDRESS_PROVINCE+this.provinceForm.ID ,this.provinceForm).then(response=>{
                        let {code,msg,provinceID} = response.data;
                        if(code=='success'){
                            this.$message({ message:msg, type: 'success' });
                            this.$router.push("/admin/address/province/detail/"+provinceID);
                        }else{
                            this.$message({ message:msg, type: 'error' });
                        }
                    });
                }else{
                    this.$ajax.post(SERVER_ADDRESS_PROVINCE,this.provinceForm).then(response=>{
                        let {code,msg,provinceID} = response.data;
                        if(code=='success'){
                            this.$message({ message:msg, type: 'success' });
                            this.$router.push("/admin/address/province/detail/"+provinceID);
                        }else{
                            this.$message({ message:msg, type: 'error' });
                        }
                    });
                }
            },
            getProvinceInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.provinceForm.ID = id;
                    this.$ajax.get(SERVER_ADDRESS_PROVINCE+this.provinceForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.provinceForm = data["province"];
                                this.countryList = [this.provinceForm.Country]
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.provinceForm = this.NewProvinceForm;
                }
            },
            getCountryList(query){
                this.$ajax.get(SERVER_ADDRESS_COUNTRY,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.countryList = data["countries"];
                    }
                });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/admin/address/province");
                }else if ("form"==type){
                    this.$router.push("/admin/address/province/form/"+id);
                }
            },
        },
        created:function(){
            this.getProvinceInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getProvinceInfo'
        },
         
    }
</script>