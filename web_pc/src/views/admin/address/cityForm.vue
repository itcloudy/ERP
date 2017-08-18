<template>
    <div v-if="edit">
        <el-form ref="cityForm" :model="cityForm" label-width="80px">
            <el-form-item label="所属国家">
                <el-select
                    v-model="cityForm.Country.Name"
                    filterable
                    remote
                    placeholder="请输入国家"
                    :remote-method="getCountryList"
                    :loading="loading">
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
    <div v-else>
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
</template>
<script>
         
    import { mapState } from 'vuex';
    export default {
        data() {
            return {
                edit:true,
                loading: false,
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
        method:{
            getCityInfo(){

            },
            getCountryList(query){
                this.loading = true;
                console.log(query);
                this.$ajax.get("/address/country",{
                    params:{
                        offset:0,
                        limit:20
                    }
                }).then(response=>{
                    this.loading = false;
                    if(code=='success'){
                        this.countryList = data["countries"];
                    }
                });
            }
        }
    }
</script>