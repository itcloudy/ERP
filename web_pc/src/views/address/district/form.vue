<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="districtForm" :rules="districtFormRules"  :model="districtForm" label-width="80px">
                <el-form-item label="所属国家"  prop="Country">
                    <el-select
                        v-model="districtForm.Country.ID"
                        :name="districtForm.Country.Name"
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
                <el-form-item label="所属省份"  prop="Province">
                    <el-select
                        v-model="districtForm.Province.ID"
                        :name="districtForm.Province.Name"
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
                <el-form-item label="所属城市"  prop="City">
                    <el-select
                        v-model="districtForm.City.ID"
                        :name="districtForm.City.Name"
                        filterable
                        remote
                        placeholder="请输入城市"
                        :remote-method="getCityList">
                        <el-option
                            v-for="item in cityList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="区县名称"  prop="Name">
                    <el-input v-model="districtForm.Name"></el-input>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';    
    import  {SERVER_ADDRESS_CITY,SERVER_ADDRESS_COUNTRY,SERVER_ADDRESS_PROVINCE,SERVER_ADDRESS_DISTRICT} from '@/server_address';             
    import { mapState } from 'vuex';
    import {validateObjectID} from '@/utils/validators';
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
                districtForm:{ },
                NewDistrictForm:{
                    Name:"",
                    ID:"",
                    City:{
                        Name:"",
                        ID:"",
                    },
                    Province:{
                        Name:"",
                        ID:"",
                    },
                    Country:{
                        Name:"",
                        ID:"",
                    }
                },
                cityList:[],
                provinceList:[],
                countryList:[],
                districtFormRules:{
                    Name:[
                        { required: true, message: '请输入区县名称', trigger: 'blur' }
                    ],
                    City:[
                        { required: true, message: '请选择城市',validator: validateObjectID, trigger: 'blur' }
                        ],
                    Province:[
                        { required: true, message: '请选择省份',validator: validateObjectID, trigger: 'blur' }
                    ],
                    Country:[
                        { required: true, message: '请选择国家',validator: validateObjectID, trigger: 'blur' }
                    ]
                }
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                this.$refs['districtForm'].validate((valid) => {
                    if (valid) {
                        if (this.districtForm.ID >0){
                            this.$ajax.put(SERVER_ADDRESS_DISTRICT+this.districtForm.ID ,this.districtForm).then(response=>{
                                let {code,msg,districtID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/address/district/detail/"+districtID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_ADDRESS_DISTRICT,this.districtForm).then(response=>{
                                let {code,msg,districtID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/address/district/detail/"+districtID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                });
            },
            getDistrictInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.districtForm.ID = id;
                    this.$ajax.get(SERVER_ADDRESS_DISTRICT+this.districtForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.districtForm = data["district"];
                                this.provinceList = [this.districtForm.Province]
                                this.countryList = [this.districtForm.Country]
                                this.cityList = [this.districtForm.City]
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.districtForm = this.NewDistrictForm;
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
            getProvinceList(query){
                this.$ajax.get(SERVER_ADDRESS_PROVINCE,{
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
            getCityList(query){
                this.$ajax.get(SERVER_ADDRESS_CITY,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.cityList = data["cities"];
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