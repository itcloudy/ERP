<template>
    <div>
        <pagination 
        @pageInfoChange="pageInfoChange"
        :pageSize="provincesData.pageSize" 
        :currentPage="provincesData.currentPage"
        :total="provincesData.total"/>
        <el-table
            ref="multipleTable"
            :data="provincesData.provinceList"
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
              prop="Name"
              label="省份">
            </el-table-column>
        </el-table>
      
    </div>
</template>
<script>
    import  {default as Pagination} from '../global/Pagination';
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            provincesData:{
                provinceList:[],//tree视图数据
                pageSize:20,//每页数量
                total:0,//总数量
                currentPage:1,//当前页
            },
        }
    },
    methods:{
        getProvinces(limit,offset){
            this.$ajax.get("/address/province/?limit="+limit +"&offset="+offset).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.provincesData.provinceList = data["provinces"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.provincesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.provincesData.pageSize = pageSize;
            this.provincesData.currentPage = currentPage;
            this.getProvinces(pageSize,(currentPage-1)*pageSize)
        }
    },
    components: {
        Pagination,
    },
    created:function(){
        this.$nextTick(function(){
            this.getProvinces(this.provincesData.pageSize,this.provincesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.provincesData.total/this.provincesData.pageSize > 1
        }
    }
      
    }
</script>