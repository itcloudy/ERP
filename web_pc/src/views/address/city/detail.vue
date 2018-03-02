<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="cityForm" :model="cityForm" :inline="true"  class="form-read-only">
                <el-form-item label="所属国家">
                    <span>{{cityForm.Country.Name}}</span>
                </el-form-item>
                <el-form-item label="所属省份">
                    <span>{{cityForm.Province.Name}}</span>
                </el-form-item>
                <el-form-item label="城市名称">
                    <span>{{cityForm.Name}}</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';   
    import  {SERVER_ADDRESS_CITY} from '@/server_address';       
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
                cityForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getCityInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.cityForm.ID = id;
                this.$ajax.get(SERVER_ADDRESS_CITY+this.cityForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.cityForm = data["city"];
                            this.access = data["access"];
                        }
                    });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/address/city");
                }else if ("form"==type){
                    this.$router.push("/address/city/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/address/city/form/"+this.cityForm.ID);
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