<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product/product' }">规格</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="productsData.pageSize" 
            :currentPage="productsData.currentPage"
            :total="productsData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="productsData.productList"
                @row-dblclick = "goProductProductDetail"
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
                label="规格名称">
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {SERVER_PRODUCT_PRODUCT} from '@/server_address'; 
    import { mapState } from 'vuex';
    export default {
        data() {
            return {
                treeViewHeight: this.$store.state.windowHeight-100,
                productsData:{
                    productList:[],//tree视图数据
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
            getYesOrNO(boolVar){
                    if (boolVar){
                        return "是";
                    }else{
                        return "否";
                    }
                },
            getProductProducts(limit,offset){
                this.loading = true;
                this.$ajax.get(SERVER_PRODUCT_PRODUCT,{
                        params:{
                            offset:offset,
                            limit:limit
                        }
                }).then(response=>{
                    this.loading = false;
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.productsData.productList = data["products"];
                        this.access = data["access"];
                        let paginator = data.paginator;
                        if (paginator){
                            this.productsData.total = paginator.totalCount;
                        }
                    }
                });
            },
            pageInfoChange(pageSize,currentPage){
                this.productsData.pageSize = pageSize;
                this.productsData.currentPage = currentPage;
                this.getProductProducts(pageSize,(currentPage-1)*pageSize)
            },
            goProductProductDetail(row, event){
                this.$router.push("/product/product/detail/"+row.ID);
            },
            changeCreateForm(){
                this.$router.push("/product/product/form/new");
            }
            
        },
        components: {
            Pagination,
            ListTop
        },
        created:function(){
            this.$nextTick(function(){
                this.getProductProducts(this.productsData.pageSize,this.productsData.currentPage-1);
            });
        },
        computed:{
            showBottomPagitator:function(){
                return this.productsData.total/this.productsData.pageSize > 1
            }
        }
    }
</script>
<style lang="scss" scoped>
    
    
</style>
