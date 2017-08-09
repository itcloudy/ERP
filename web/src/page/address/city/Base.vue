<template>
    <div>
        <city-tree v-show="showTree" @changeViewType="changeViewType" :cityList="cityList"></city-tree>
        <city-form v-show="showForm" @changeViewType="changeViewType" :city="city"></city-form>
    </div>
    
</template>

<script>
    import  {default as cityTree} from './Tree';
    import  {default as cityForm} from './Form';
    export default {
        data() {
            
            return {
                showTree:true,
                showForm:false,
                cityList:[],
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
                       this.cityList = data["cities"]
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