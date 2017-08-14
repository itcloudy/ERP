<template>
    <div>
        <pagination 
        @pageInfoChange="pageInfoChange"
        :pageSize="countriesData.pageSize" 
        :currentPage="countriesData.currentPage"
        :total="countriesData.total"/>
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
              label="国家">
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
            countriesData:{
                countryList:[],//tree视图数据
                pageSize:20,//每页数量
                total:0,//总数量
                currentPage:1,//当前页
            },
        }
    },
    methods:{
        getCountries(limit,offset){
            this.$ajax.get("/address/country/?limit="+limit +"&offset="+offset).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.countriesData.countryList = data["countries"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.countriesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.countriesData.pageSize = pageSize;
            this.countriesData.currentPage = currentPage;
            this.getCountries(pageSize,(currentPage-1)*pageSize)
        }
    },
    components: {
        Pagination,
    },
    created:function(){
        this.$nextTick(function(){
            this.getCountries(this.countriesData.pageSize,this.countriesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.countriesData.total/this.countriesData.pageSize > 1
        }
    }
      
    }
</script>