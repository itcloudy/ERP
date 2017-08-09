<template>
    <div>
        <button @click="changeViewType('form')">Form</button>
         <button @click="changeViewType('tree')">Tree</button>
        <city-tree v-if="showTree" v-on:changeViewType="changeViewType" cityList="cityList"></city-tree>
        <city-form v-if="showForm" v-on:changeViewType="changeViewType" city="city"></city-form>
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