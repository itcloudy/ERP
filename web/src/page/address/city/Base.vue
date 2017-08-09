<template>
    <div>
       
         
        <city-tree v-if="showTree" @changeViewType="changeViewType" cityList="cityList"></city-tree>
        <city-form v-if="showForm" @changeViewType="changeViewType" city="city"></city-form>
        <div>{{showTree}}</div>
        <div>{{showForm}}</div>
        <div>{{cityTree}}</div>
        <div>{{city}}</div>
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
          }
        }
      
    }
</script>