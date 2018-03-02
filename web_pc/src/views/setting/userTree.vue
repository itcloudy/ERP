<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">地址管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address/user' }">城市</el-breadcrumb-item>
        </el-breadcrumb>
        <div>
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="usersData.pageSize" 
            :currentPage="usersData.currentPage"
            :total="usersData.total"/>
            <el-table
                ref="multipleTable"
                :data="usersData.userList"
                @row-dblclick = "goUserDetail"
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
                label="登录名">
                </el-table-column>
                <el-table-column
                prop="NameZh"
                label="中文名称">
                </el-table-column>
                <el-table-column
                prop="Email"
                label="邮箱">
                </el-table-column>
                <el-table-column
                prop="Mobile"
                label="手机号码">
                </el-table-column>
                <el-table-column
                prop="Tel"
                label="电话">
                </el-table-column>
                <el-table-column
                prop="IsAdmin"
                label="系统管理员">
                </el-table-column>
                .<el-table-column
                prop="Active"
                label="有效">
                </el-table-column>
                <el-table-column
                prop="Qq"
                label="QQ">
                </el-table-column>
                <el-table-column
                prop="WeChat"
                label="微信">
                </el-table-column>
                 <el-table-column
                prop="IsBackground"
                label="后台用户">
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
            usersData:{
                userList:[],//tree视图数据
                pageSize:20,//每页数量
                total:0,//总数量
                currentPage:1,//当前页
            },
            serverUrlPath:"/setting/user"
        }
    },
    methods:{
        getUsers(limit,offset){
            this.$ajax.get(this.serverUrlPath,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.usersData.userList = data["users"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.usersData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.usersData.pageSize = pageSize;
            this.usersData.currentPage = currentPage;
            this.getUsers(pageSize,(currentPage-1)*pageSize)
        },
        goUserDetail(row, event){
            this.$router.push("/setting/user/"+row.ID);
        }
    },
    components: {
        Pagination,
    },
    created:function(){
        this.$nextTick(function(){
            this.getUsers(this.usersData.pageSize,this.usersData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.usersData.total/this.usersData.pageSize > 1
        }
    }
      
    }
</script>