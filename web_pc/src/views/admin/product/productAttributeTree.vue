<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/admin' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/admin/product' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/admin/product/attribute' }">产品属性</el-breadcrumb-item>
        </el-breadcrumb>
        <div>
            <TreeTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="attributesData.pageSize" 
            :currentPage="attributesData.currentPage"
            :total="attributesData.total"/>
            <el-table
                ref="multipleTable"
                :data="attributesData.attributeList"
                @row-dblclick = "goProductAttributeDetail"
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
                label="属性">
                </el-table-column>
                <el-table-column
                prop="Code"
                label="属性编码">
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '../global/Pagination';
    import  {default as TreeTop} from '../global/TreeTop'; 
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            attributesData:{
                attributeList:[],//tree视图数据
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
            serverUrlPath:"/product/attribute"
        }
    },
    methods:{
        getProductAttributes(limit,offset){
            this.$ajax.get(this.serverUrlPath,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.attributesData.attributeList = data["attributes"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.attributesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.attributesData.pageSize = pageSize;
            this.attributesData.currentPage = currentPage;
            this.getProductAttributes(pageSize,(currentPage-1)*pageSize)
        },
        goProductAttributeDetail(row, event){
            this.$router.push("/admin/product/attribute/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/admin/product/attribute/new");
        }
    },
    components: {
        Pagination,
        TreeTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getProductAttributes(this.attributesData.pageSize,this.attributesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.attributesData.total/this.attributesData.pageSize > 1
        }
    }
      
    }
</script>