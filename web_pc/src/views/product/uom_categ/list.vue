<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product/uomcateg' }">单位类别</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="uomcategsData.pageSize" 
            :currentPage="uomcategsData.currentPage"
            :total="uomcategsData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="uomcategsData.uomcategList"
                @row-dblclick = "goProductUomDetail"
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
                label="名称">
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {SERVER_PRODUCT_UOM_CATEG} from '@/server_address';
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            uomcategsData:{
                uomcategList:[],//tree视图数据
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
        getProductUoms(limit,offset){
            this.loading = true;
            this.$ajax.get(SERVER_PRODUCT_UOM_CATEG,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
            }).then(response=>{
                this.loading = false;
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.uomcategsData.uomcategList = data["uomcategs"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.uomcategsData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.uomcategsData.pageSize = pageSize;
            this.uomcategsData.currentPage = currentPage;
            this.getProductUoms(pageSize,(currentPage-1)*pageSize)
        },
        goProductUomDetail(row, event){
            this.$router.push("/product/uomcateg/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/product/uomcateg/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getProductUoms(this.uomcategsData.pageSize,this.uomcategsData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.uomcategsData.total/this.uomcategsData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
    
    
</style>
