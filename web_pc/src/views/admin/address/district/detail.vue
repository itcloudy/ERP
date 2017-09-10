<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        :edit="edit"
        @formSave="formSave"
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
    import  {default as FormTop} from '@/views/admin/common/FormTop';         
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
                cityForm:{
                    Name:"",
                    ID:0,
                    Province:{
                        Name:"",
                        ID:0,
                    },
                    Country:{
                        Name:"",
                        ID:0,
                    }
                },
                NewCityForm:{
                    Name:"",
                    ID:0,
                    Province:{
                        Name:"",
                        ID:0,
                    },
                    Country:{
                        Name:"",
                        ID:0,
                    }
                },
                countryList:[],
                provinceList:[]
            }
        },
        components:{
           FormTop
        },
        methods:{
            getCityInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;

                if (id!='new'){
                    this.cityForm.ID = id;
                    this.$ajax.get("/address/city/"+this.cityForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.cityForm = data["city"];
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.edit = true;
                    this.cityForm = this.NewCityForm;
                }
                
                
            },
            getCountryList(query){
                this.$ajax.get("/address/country",{
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
            getCountryList(query){
                this.$ajax.get("/address/province",{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.provinceList = data["privinces"];
                    }
                });
            },
            changeView(type,id){
                console.log(type);
                if ("list"==type){
                    this.$router.push("/admin/address/city");
                }else if ("form"==type){
                    this.$router.push("/admin/address/city/form/"+id);
                }
            },
            formEdit(){
                this.edit = true;
            },
            formSave(){
                this.edit = false;
            }
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