<template>
    <div class="block">
        <mt-header title="城市管理" fixed>
            <router-link to="/" slot="left">
                <mt-button icon="back" @click="goBack">返回</mt-button>
            </router-link>
        </mt-header>
        <section id="content-section">
            <template v-for="(city,index) in citiesData.cityList" >
                <mt-cell :key="index" :title="city.Name" :to="'/admin/address/city/' + city.ID" is-link></mt-cell>
                <div :key="index">
                    <table>
                        <tbody>
                            <tr><td>国家：</td><td>{{city.Country.Name}}</td></tr>
                            <tr><td>省份：</td><td>{{city.Province.Name}}</td></tr>
                        </tbody>
                    </table>
                </div>
            </template>
            
        </section>

    </div>
</template>
<script>
    import  {SERVER_ADDRESS_CITY} from '@/server_address'; 
    import { Toast } from 'mint-ui';
    export default {
    data() {
        return {
            citiesData:{
                cityList:[],//tree视图数据
                pageSize:20,//每页数量
                total:0,//总数量
                currentPage:1,//当前页
            },
        };
    },
    methods:{
        goBack(){
            this.$router.go(-1);
        },
        getCities(limit,offset){
            this.$ajax.get(SERVER_ADDRESS_CITY,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.citiesData.cityList= data["cities"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.citiesData.total = paginator.totalCount;
                    }
                }
            });
        },
    },
    created:function(){
        this.$nextTick(function(){
            this.getCities(this.citiesData.pageSize,this.citiesData.currentPage-1);
        });
    },
    
  }
</script>
<style lang="scss" scope>
   
</style>

