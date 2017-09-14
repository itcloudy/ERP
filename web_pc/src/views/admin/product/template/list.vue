<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/admin' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/admin/address' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/admin/product/template' }">款式</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="templatesData.pageSize" 
            :currentPage="templatesData.currentPage"
            :total="templatesData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="templatesData.templateList"
                @row-dblclick = "goProductTemplateDetail"
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
                label="款式名称">
                </el-table-column>
                <el-table-column
                prop="Category.Name"
                label="所属省份">
                </el-table-column>
                <el-table-column
                prop="Rental"
                label="代售品">
                </el-table-column>
                <el-table-column
                prop="Price"
                label="款式价格">
                </el-table-column>
                <el-table-column
                prop="StandardPrice"
                label="成本价格">
                </el-table-column>
                <el-table-column
                prop="StandardWeight"
                label="标准重量">
                </el-table-column>
                <el-table-column
                prop="SaleOk"
                label="可销售">
                </el-table-column>
                <el-table-column
                prop="Active"
                label="有效">
                </el-table-column>
                <el-table-column
                prop="IsProductVariant"
                label="是规格产品">
                </el-table-column>
                <el-table-column
                prop="FirstSaleUom.Name"
                label="第一销售单位">
                </el-table-column>
                <el-table-column
                prop="SecondSaleUom.Name"
                label="第二销售单位">
                </el-table-column>
                <el-table-column
                prop="VariantCount"
                label="产品规格数量">
                </el-table-column>
                 
                <el-table-column
                prop="DefaultCode"
                label="产品编码">
                </el-table-column>
                <el-table-column
                prop="ProductType"
                label="产品类型">
                </el-table-column>
                <el-table-column
                prop="ProductMethod"
                label="规格创建方式">
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/admin/common/Pagination';
    import  {default as ListTop} from '@/views/admin/common/ListTop'; 
    import  {SERVER_PRODUCT_TEMPLATE} from '@/server_address'; 
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            templatesData:{
                templateList:[],//tree视图数据
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
        getProductTemplates(limit,offset){
            this.loading = true;
            this.$ajax.get(SERVER_PRODUCT_TEMPLATE,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                    this.loading = false;
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.templatesData.templateList = data["templates"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.templatesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.templatesData.pageSize = pageSize;
            this.templatesData.currentPage = currentPage;
            this.getProductTemplates(pageSize,(currentPage-1)*pageSize)
        },
        goProductTemplateDetail(row, event){
            this.$router.push("/admin/product/template/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/admin/product/template/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getProductTemplates(this.templatesData.pageSize,this.templatesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.templatesData.total/this.templatesData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
    
    
</style>
