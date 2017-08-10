<template>
    <div>
        <button @click="changeView('form')">Form</button>

        <el-table
            ref="multipleTable"
            :data="countriesData.countryList"
            style="width: 100%">
            <el-table-column
              type="selection"
              width="55">
            </el-table-column>
            <el-table-column
              prop="ID"
              label="ID">
            </el-table-column>
            <el-table-column
              prop="Name"
              label="名称">
            </el-table-column>
        </el-table>
        <pagination 
        @pageInfoChange="pageInfoChange"
        :pageSize="countriesData.pageSize" 
        :currentPage="countriesData.currentPage"
        :total="countriesData.total"/> 
      
    </div>
</template>

<script>
  import  {default as Pagination} from '../../global/Pagination';
  import { mapState } from 'vuex';
  
  export default {
    data() {
      return {
        treeViewHeight: this.$store.state.windowHeight-100,
      }
    },
    components: {
           Pagination,
    },
    props:["countriesData"],
    methods:{
      changeView(type){
        this.$emit("changeViewType",type);
      },
      pageInfoChange(pageSize,currentPage){
        this.$emit("pageInfoChange",pageSize,currentPage);
      }
    }
     
  }
</script>