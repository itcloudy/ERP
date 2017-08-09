<template>
   <div>
   <button @click="changeView('tree')">Tree</button>
   <p>23424</p>
   </div>
</template>

<script>
    export default {
    data() {
      return {
      }
    },
    props:["city"],
    methods:{
      changeView(type){
        this.$emit("changeViewType",type);
      }
    }
     
  }
</script>