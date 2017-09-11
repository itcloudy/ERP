<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="provinceForm" :model="provinceForm" :inline="true"  class="form-read-only">
                <el-form-item label="所属国家">
                    <span v-if="provinceForm.Country">{{provinceForm.Country.Name}}</span>
                    <span v-else>未知</span>
                </el-form-item>
                <el-form-item label="城市名称">
                    <span>{{provinceForm.Name}}</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/admin/common/FormTop';         
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
                provinceForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getCityInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.provinceForm.ID = id;
                this.$ajax.get("/address/province/"+this.provinceForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.provinceForm = data["province"];
                            this.access = data["access"];
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
            formEdit(){
                 this.$router.push("/admin/address/province/form/"+this.provinceForm.ID);
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