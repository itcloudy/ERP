<template>
    <div>
        <button @click="changeView('form')">Form</button>
        <pagination 
        @pageInfoChange="pageInfoChange"
        :pageSize="districtsData.pageSize" 
        :currentPage="districtsData.currentPage"
        :total="districtsData.total"/> 
        <el-table
            ref="multipleTable"
            :data="districtsData.districtList"
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
              prop="City.Name"
              label="所属城市">
            </el-table-column>
            <el-table-column
              prop="Name"
              label="名称">
            </el-table-column>
        </el-table>
        <pagination 
         v-if="showBottomPagitator"
        @pageInfoChange="pageInfoChange"
        :pageSize="districtsData.pageSize" 
        :currentPage="districtsData.currentPage"
        :total="districtsData.total"/> 
      
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
    props:["districtsData"],
    methods:{
      changeView(type){
        this.$emit("changeViewType",type);
      },
      pageInfoChange(pageSize,currentPage){
        this.$emit("pageInfoChange",pageSize,currentPage);
      }
    },
    computed:{
      showBottomPagitator:function(){
        return this.districtsData.total/this.districtsData.pageSize > 1
      }
    }
     
  }
</script>