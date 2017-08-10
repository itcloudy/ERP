<template>
    <div>
        <button @click="changeView('form')">Form</button>

        <el-table
            ref="multipleTable"
            :data="citiesData.cityList"
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
              prop="Country.Name"
              label="所属国家">
            </el-table-column>
            <el-table-column
              prop="Province.Name"
              label="所属省份">
            </el-table-column>
            <el-table-column
              prop="Name"
              label="名称">
            </el-table-column>
        </el-table>
        <pagination 
        @pageInfoChange="pageInfoChange"
        :pageSize="citiesData.pageSize" 
        :currentPage="citiesData.currentPage"
        :total="citiesData.total"/> 
      
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
    props:["citiesData"],
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