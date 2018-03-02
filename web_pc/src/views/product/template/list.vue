<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product/template' }">款式</el-breadcrumb-item>
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
                label="产品分类">
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
                label="可销售">
                    <template scope="scope">
                        {{getYesOrNO(scope.row.SaleOK)}}
                    </template>
                </el-table-column>
                <el-table-column
                label="有效">
                    <template scope="scope">
                        {{getYesOrNO(scope.row.Active)}}
                    </template>
                </el-table-column>
                <el-table-column
                label="是规格产品">
                    <template scope="scope">
                        {{getYesOrNO(scope.row.IsProductVariant)}}
                    </template>
                </el-table-column>
                <el-table-column
                prop="FirstSaleUom.Name"
                label="第一销售单位">
                </el-table-column>
                <el-table-column v-if="multiUnit"
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
                label="产品类型" >
                    <template scope="scope">
                        <span>{{getProductTypeDict(scope.row.ProductType)}}</span>
                    </template>
                </el-table-column>
                <el-table-column
                label="规格创建方式">
                    <template scope="scope" >
                        <span>{{getProductMethodDict(scope.row.ProductMethod)}}</span>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
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
            productTypeDict:{"stock":"库存商品","consume":"消耗品","service":"服务"},
            productMethodDict:{"hand":"手动","auto":"自动"}
        }
    },
    methods:{
        getProductTypeDict(value){
            return this.productTypeDict[value];
        },
        getProductMethodDict(value){
            return this.productMethodDict[value];
        },
        getYesOrNO(boolVar){
                if (boolVar){
                    return "是";
                }else{
                    return "否";
                }
            },
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
            this.$router.push("/product/template/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/product/template/form/new");
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
        },
        ...mapState({
               multiUnit: state => state.multiUnit
        })
    }
      
    }
</script>
<style lang="scss" scoped>
    
    
</style>
