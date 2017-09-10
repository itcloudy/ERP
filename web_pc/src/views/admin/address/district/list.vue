<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/admin' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/admin/address' }">地址管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/admin/address/district' }">区县</el-breadcrumb-item>
        </el-breadcrumb>
        <div>
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
    import  {default as Pagination} from '@/views/admin/common/Pagination';
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
            serverUrlPath:"/address/district"
        }
    },
    methods:{
        getDistricts(limit,offset){
            this.$ajax.get(this.serverUrlPath,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.districtsData.districtList = data["districts"];
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
            this.$router.push("/admin/address/district/"+row.ID);
        }
    },
    components: {
        Pagination,
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