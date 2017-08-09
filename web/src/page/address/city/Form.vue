<template>
   <div>
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