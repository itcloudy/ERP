<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product/category' }">类别</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="catetoriesData.pageSize" 
            :currentPage="catetoriesData.currentPage"
            :total="catetoriesData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="catetoriesData.categoryList"
                @row-dblclick = "goProductCategoryDetail"
                style="width: 100%">
                <el-table-column
                type="selection"
                width="55">
                </el-table-column>
                <el-table-column
                prop="Parent.Name"
                label="上级分类">
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
    import  {SERVER_PRODUCT_CATEGORY} from '@/server_address';  
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            catetoriesData:{
                categoryList:[],//tree视图数据
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
        getProductCategories(limit,offset){
            this.loading = true;
            this.$ajax.get(SERVER_PRODUCT_CATEGORY,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                    this.loading = false;
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.catetoriesData.categoryList = data["categories"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.catetoriesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.catetoriesData.pageSize = pageSize;
            this.catetoriesData.currentPage = currentPage;
            this.getProductCategories(pageSize,(currentPage-1)*pageSize)
        },
        goProductCategoryDetail(row, event){
            this.$router.push("/product/category/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/product/category/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getProductCategories(this.catetoriesData.pageSize,this.catetoriesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.catetoriesData.total/this.catetoriesData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
    
    
</style>
