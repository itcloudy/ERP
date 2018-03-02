<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>       
        <div v-loading="loading">
            <el-form ref="countryForm" :model="countryForm" :inline="true"  class="form-read-only">
                <el-form-item label="国家名称">
                    <span>{{countryForm.Name}}</span>
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
                edit:false,
                loadging:false,
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                countryForm:{
                    Name:"",
                    ID:"",
                     
                },
                
            }
        },
        components:{
           FormTop
        },
        methods:{
            getcountryInfo(){
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
                } 
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/address/country");
                }else if ("form"==type){
                    this.$router.push("/address/country/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/address/country/form/"+this.countryForm.ID);
            },
        },
        created:function(){
            this.getcountryInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getcountryInfo'
        },
         
    }
</script>