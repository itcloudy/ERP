<template>
   <div class="block">
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10,20,50,100, 200, 300, 400]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total">
    </el-pagination>
  </div>
</template>
<script>
  export default {
    methods: {
      pageInfoChange(pageSize,currentPage){
        this.$emit("pageInfoChange",pageSize,currentPage);
      },
      handleSizeChange(val) {
        // 每页显示数量改变，从第一页开始显示
        this.pageInfoChange(val,1);
      },
      handleCurrentChange(val) {
        this.pageInfoChange(this.pageSize,val);
      }
    },
    props:["pageSize","currentPage","total"],
    data() {
      return {
        
      };
    }
  }
</script>
<style lang="scss" scope>
    .block{
        margin: 5px;
        float:right;
    }
</style>
