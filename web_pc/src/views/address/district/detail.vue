<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="districtForm" :model="districtForm" :inline="true"  class="form-read-only">
                <el-form-item label="所属国家">
                    <span v-if="districtForm.Country">{{districtForm.Country.Name}}</span>
                    <span v-else>未知</span>
                </el-form-item>
                <el-form-item label="所属省份">
                    <span v-if="districtForm.Province">{{districtForm.Province.Name}}</span>
                    <span v-else>未知</span>
                </el-form-item>
                <el-form-item label="所属城市">
                    <span v-if="districtForm.City">{{districtForm.City.Name}}</span>
                    <span v-else>未知</span>
                </el-form-item>
                <el-form-item label="区县名称">
                    <span>{{districtForm.Name}}</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop'; 
    import  {SERVER_ADDRESS_CITY,SERVER_ADDRESS_COUNTRY,SERVER_ADDRESS_PROVINCE,SERVER_ADDRESS_DISTRICT} from '@/server_address';        
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
                districtForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getDistrictInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.districtForm.ID = id;
                this.$ajax.get(SERVER_ADDRESS_DISTRICT+this.districtForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.districtForm = data["district"];
                            this.access = data["access"];
                        }
                    });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/address/district");
                }else if ("form"==type){
                    this.$router.push("/address/district/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/address/district/form/"+this.districtForm.ID);
            },
        },
        created:function(){
            this.getDistrictInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getDistrictInfo'
        },
         
    }
</script>