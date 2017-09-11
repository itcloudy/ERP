<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="cityForm" :model="cityForm" label-width="80px">
                <el-form-item label="所属国家">
                    <el-select
                        v-model="cityForm.Country.ID"
                        :name="cityForm.Country.Name"
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
                 <el-form-item label="所属省份">
                    <el-select
                        v-model="cityForm.Province.ID"
                        :name="cityForm.Province.Name"
                        filterable
                        remote
                        placeholder="请输入省份"
                        :remote-method="getProvinceList">
                        <el-option
                            v-for="item in provinceList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="城市名称">
                    <el-input v-model="cityForm.Name"></el-input>
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
                provinceList:[],
                countryList:[],
                
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                if (this.cityForm.ID >0){
                    this.$ajax.put("/address/city",this.cityForm).then(response=>{
                        let {code,msg,cityID} = response.data;
                        if(code=='success'){
                            this.$message({ message:msg, type: 'success' });
                            this.$router.push("/admin/address/city/detail/"+cityID);
                        }else{
                            this.$message({ message:msg, type: 'error' });
                        }
                    });
                }else{
                    this.$ajax.post("/address/city",this.cityForm).then(response=>{
                        let {code,msg,cityID} = response.data;
                        if(code=='success'){
                            this.$message({ message:msg, type: 'success' });
                            this.$router.push("/admin/address/city/detail/"+cityID);
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
                    this.cityForm.ID = id;
                    this.$ajax.get("/address/city/"+this.cityForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.cityForm = data["city"];
                                this.provinceList = [this.cityForm.Province]
                                this.countryList = [this.cityForm.Country]
                                this.access = data["access"];
                            }
                        });
                }else{
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
            getProvinceList(query){
                this.$ajax.get("/address/province",{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.provinceList = data["provinces"];
                    }
                });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/admin/address/city");
                }else if ("form"==type){
                    this.$router.push("/admin/address/city/form/"+id);
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