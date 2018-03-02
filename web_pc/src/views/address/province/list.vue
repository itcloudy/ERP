<template>
    <div>
         <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">地址管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address/province' }">省份</el-breadcrumb-item>
        </el-breadcrumb>
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="provincesData.pageSize" 
            :currentPage="provincesData.currentPage"
            :total="provincesData.total"/>
            <el-table
                ref="multipleTable"
                :data="provincesData.provinceList"
                 @row-dblclick = "goProvinceDetail"
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
    </div>
</template>
<script>
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {SERVER_ADDRESS_PROVINCE} from '@/server_address';        
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
            access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
            },
            
        }
    },
    methods:{
        getProvinces(limit,offset){
            this.$ajax.get(SERVER_ADDRESS_PROVINCE,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.provincesData.provinceList = data["provinces"];
                    this.access = data["access"];
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
        },
        goProvinceDetail(row, event){
            this.$router.push("/address/province/detail/"+row.ID);
        },
         changeCreateForm(){
            this.$router.push("/address/province/form/new");
        }
    },
    components: {
        Pagination,
        ListTop
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