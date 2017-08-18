<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        :edit="edit"
        @formSave="formSave"
        @changeView="changeView"/>
        <div v-if="edit"  v-loading="loading">
            <el-form ref="cityForm" :model="cityForm" label-width="80px">
                <el-form-item label="所属国家">
                    <el-select
                        v-model="cityForm.Country.Name"
                        filterable
                        remote
                        placeholder="请输入国家"
                        :remote-method="getCountryList"
                        >
                        <el-option
                            v-for="item in countryList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.Name">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="所属省份">
                    <span>{{cityForm.Province.Name}}</span>
                </el-form-item>
                <el-form-item label="城市名称">
                    <span>{{cityForm.Name}}</span>
                </el-form-item>
        
            </el-form>
        </div>
        <div v-else  v-loading="loading">
            <el-form ref="cityForm" :model="cityForm" inline label-width="80px">
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
    import  {default as FormTop} from '../global/FormTop';         
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
                    Name:"123",
                    ID:0,
                    Province:{
                        Name:"456",
                        ID:0,
                    },
                    Country:{
                        Name:"",
                        ID:0,
                    }
                },
                countryList:[{Name:"2342",ID:12},{Name:"2334342",ID:124}]
            }
        },
        components:{
           FormTop
        },
        methods:{
            getCityInfo(){
                this.loadging = true;
                 this.cityForm.ID = this.$route.params.id;
                 this.$ajax.get("/address/city/"+this.cityForm.ID).then(response=>{
                     this.loadging = false;
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.cityForm = data["city"];
                        this.access = data["access"];
                    }
                });
            },
            getCountryList(query){
                console.log(query);
                this.$ajax.get("/address/country",{
                    params:{
                        offset:0,
                        limit:20
                    }
                }).then(response=>{
                    console.log(response);
                    if(code=='success'){
                        this.countryList = data["countries"];
                    }
                });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/admin/address/city");
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