<template>
    <div>
        <city-tree v-show="showTree" @changeViewType="changeViewType" :citiesData="citiesData"></city-tree>
        <city-form v-show="showForm" @changeViewType="changeViewType" :city="city"></city-form>
    </div>
    
</template>

<script>
    import  {default as cityTree} from './Tree';
    import  {default as cityForm} from './Form';
    export default {
        data() {
            
            return {
                showTree:true,//显示tree视图
                showForm:false,//显示form视图
                citiesData:{
                    cityList:[],//tree视图数据
                    pageSize:20,//每页数量
                    total:0,//总数量
                    currentPage:1,//当前页
                },
                
                city:{}

            }
        },
        components: {
            cityTree,
            cityForm,
        },
        methods:{
            changeViewType(type){
                if ('form'==type){
                    this.showTree = false
                    this.showForm = true;
                }else if ('tree'== type){
                    this.showTree = true
                    this.showForm = false;
                }else{
                    this.showTree = true
                    this.showForm = false;
                }
            },
            getCities(){
                this.$ajax.get("/address/city/").then(response=>{
                   let {code,msg,data} = response.data;
                   if(code=='success'){
                       this.citiesData.cityList = data["cities"]
                   }
                });
            }
        },
        created:function(){
            this.$nextTick(function(){
                this.getCities();
            });
        }

      
    }
</script>