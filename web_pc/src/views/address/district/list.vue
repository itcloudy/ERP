<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">地址管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address/district' }">区县</el-breadcrumb-item>
        </el-breadcrumb>
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="districtsData.pageSize" 
            :currentPage="districtsData.currentPage"
            :total="districtsData.total"/>
            <el-table
                ref="multipleTable"
                :data="districtsData.districtList"
                @row-dblclick = "goDistrictDetail"
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
                label="区县">
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {SERVER_ADDRESS_DISTRICT} from '@/server_address';        
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            districtsData:{
                districtList:[],//tree视图数据
                pageSize:20,//每页数量
                total:0,//总数量
                currentPage:1,//当前页
            },
             access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
            },
        
        }
    },
    methods:{
        getDistricts(limit,offset){
            this.$ajax.get(SERVER_ADDRESS_DISTRICT,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.districtsData.districtList = data["districts"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.districtsData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.districtsData.pageSize = pageSize;
            this.districtsData.currentPage = currentPage;
            this.getDistricts(pageSize,(currentPage-1)*pageSize)
        },
        goDistrictDetail(row, event){
            this.$router.push("/address/district/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/address/district/form/new");
        }
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getDistricts(this.districtsData.pageSize,this.districtsData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.districtsData.total/this.districtsData.pageSize > 1
        }
    }
      
    }
</script>