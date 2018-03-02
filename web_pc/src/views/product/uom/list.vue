<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product/uom' }">单位</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="uomsData.pageSize" 
            :currentPage="uomsData.currentPage"
            :total="uomsData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="uomsData.uomList"
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
                <el-table-column
                prop="Category.Name"
                label="单位类别">
                </el-table-column>
                <el-table-column
                prop="Factor"
                label="比率">
                </el-table-column>
                <el-table-column
                prop="FactorInv"
                label="更大比率">
                </el-table-column>
                <el-table-column
                prop="Rounding"
                label="舍入精度">
                </el-table-column>
                <el-table-column
                prop="Type"
                label="类型">
                </el-table-column>
                <el-table-column
                prop="Symbol"
                label="符号位置">
                </el-table-column>

            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {SERVER_PRODUCT_UOM} from '@/server_address';
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            uomsData:{
                uomList:[],//tree视图数据
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
            this.$ajax.get(SERVER_PRODUCT_UOM,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                    this.loading = false;
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.uomsData.uomList = data["uoms"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.uomsData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.uomsData.pageSize = pageSize;
            this.uomsData.currentPage = currentPage;
            this.getProductUoms(pageSize,(currentPage-1)*pageSize)
        },
        goProductUomDetail(row, event){
            this.$router.push("/product/uom/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/product/uom/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getProductUoms(this.uomsData.pageSize,this.uomsData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.uomsData.total/this.uomsData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
    
    
</style>
