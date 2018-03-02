<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/sale' }">销售管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/sale/order' }">销售订单</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="ordersData.pageSize" 
            :currentPage="ordersData.currentPage"
            :total="ordersData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="ordersData.orderList"
                @row-dblclick = "goOrderDetail"
                style="width: 100%">
                <el-table-column
                type="selection"
                width="55">
                </el-table-column>
                <el-table-column
                prop="Name"
                label="订单号">
                </el-table-column>
                <el-table-column
                label="客户">
                    <template scope="scope">
                        {{scope.row.Partner.Name}}
                    </template>
                </el-table-column>
                <el-table-column
                label="公司">
                    <template scope="scope">
                        {{scope.row.Company.Name}}
                    </template>
                </el-table-column>
                <el-table-column
                label="业务员">
                    <template scope="scope">
                        {{scope.row.SalesMan.Name}}
                    </template>
                </el-table-column>
                <el-table-column
                label="收货地址">
                    <template scope="scope">
                        {{scope.row.Country.Name}}{{scope.row.Province.Name}}{{scope.row.City.Name}}{{scope.row.District.Name}}{{scope.row.Street}}
                    </template>
                </el-table-column>
                <el-table-column
                label="状态">
                    <template scope="scope">
                        {{scope.row.State}}
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {SERVER_SALE_ORDER} from '@/server_address'; 
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            ordersData:{
                orderList:[],//tree视图数据
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
        getOrders(limit,offset){
            this.loading = true;
            this.$ajax.get(SERVER_SALE_ORDER,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                    this.loading = false;
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.ordersData.orderList = data["orders"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.ordersData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.ordersData.pageSize = pageSize;
            this.ordersData.currentPage = currentPage;
            this.getOrders(pageSize,(currentPage-1)*pageSize)
        },
        goOrderDetail(row, event){
            this.$router.push("/sale/order/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/sale/order/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getOrders(this.ordersData.pageSize,this.ordersData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.ordersData.total/this.ordersData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
    
    
</style>
