<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">地址管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address/group' }">城市</el-breadcrumb-item>
        </el-breadcrumb>
        <div>
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="groupsData.pageSize" 
            :currentPage="groupsData.currentPage"
            :total="groupsData.total"/>
            <el-table
                ref="multipleTable"
                :data="groupsData.groupList"
                @row-dblclick = "goGroupDetail"
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
                label="权限组名称">
                </el-table-column>
                <el-table-column
                prop="Category"
                label="权限组分类">
                </el-table-column>
                <el-table-column
                prop="Description"
                label="权限组说明">
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '../common/Pagination';
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            groupsData:{
                groupList:[],//tree视图数据
                pageSize:20,//每页数量
                total:0,//总数量
                currentPage:1,//当前页
            },
            serverUrlPath:"/setting/group"
        }
    },
    methods:{
        getGroups(limit,offset){
            this.$ajax.get(this.serverUrlPath,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.groupsData.groupList = data["groups"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.groupsData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.groupsData.pageSize = pageSize;
            this.groupsData.currentPage = currentPage;
            this.getGroups(pageSize,(currentPage-1)*pageSize)
        },
        goGroupDetail(row, event){
            this.$router.push("/setting/group/"+row.ID);
        }
    },
    components: {
        Pagination,
    },
    created:function(){
        this.$nextTick(function(){
            this.getGroups(this.groupsData.pageSize,this.groupsData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.groupsData.total/this.groupsData.pageSize > 1
        }
    }
      
    }
</script>