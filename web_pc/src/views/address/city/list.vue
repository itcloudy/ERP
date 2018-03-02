<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">地址管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address/city' }">城市</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="citiesData.pageSize" 
            :currentPage="citiesData.currentPage"
            :total="citiesData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="citiesData.cityList"
                @row-dblclick = "goCityDetail"
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
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {SERVER_ADDRESS_CITY} from '@/server_address'; 
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            citiesData:{
                cityList:[],//tree视图数据
                pageSize:20,//每页数量
                total:0,//总数量
                currentPage:1,//当前页
            },
            loading: false,
            access:{
                Create:false,
                Update:false,
                Read:false,
                Unlink:false,
            },
        }
    },
    methods:{
        getCities(limit,offset){
            this.loading = true;
            this.$ajax.get(SERVER_ADDRESS_CITY,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                    this.loading = false;
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.citiesData.cityList = data["cities"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.citiesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.citiesData.pageSize = pageSize;
            this.citiesData.currentPage = currentPage;
            this.getCities(pageSize,(currentPage-1)*pageSize)
        },
        goCityDetail(row, event){
            this.$router.push("/address/city/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/address/city/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getCities(this.citiesData.pageSize,this.citiesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.citiesData.total/this.citiesData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
    
    
</style>
