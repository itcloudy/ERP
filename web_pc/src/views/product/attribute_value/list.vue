<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product/attributevalue' }">属性值</el-breadcrumb-item>
        </el-breadcrumb>
         
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="valuesData.pageSize" 
            :currentPage="valuesData.currentPage"
            :total="valuesData.total"/>
            <el-table
                v-loading.body="loading"
                ref="multipleTable"
                :data="valuesData.valueList"
                @row-dblclick = "goAttributeValueDetail"
                style="width: 100%">
                <el-table-column
                type="selection"
                width="55">
                </el-table-column>
                 <el-table-column
                prop="Attribute.Name"
                label="属性">
                </el-table-column>
                <el-table-column
                prop="Name"
                label="属性值名称">
                </el-table-column>
                <el-table-column label="操作">
                    <template scope="scope">
                        <el-button
                        type="danger"
                        size="mini"
                        @click="deleteProductAttributeValue(scope.$index, scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {SERVER_PRODUCT_ATTRIBUTE_VALUE} from '@/server_address';             
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            valuesData:{
                valueList:[],//tree视图数据
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
        getAttributeValues(limit,offset){
            this.loading = true;
            this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE_VALUE,{
                params:{
                    offset:offset,
                    limit:limit
                }
            }).then(response=>{
                this.loading = false;
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.valuesData.valueList = data["attributeValues"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.valuesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.valuesData.pageSize = pageSize;
            this.valuesData.currentPage = currentPage;
            this.getAttributeValues(pageSize,(currentPage-1)*pageSize)
        },
        goAttributeValueDetail(row, event){
            this.$router.push("/product/attributevalue/detail/"+row.ID);
        },
        deleteProductAttributeValue(index,row){
            this.$ajax.delete(SERVER_PRODUCT_ATTRIBUTE_VALUE +row.ID).then(response=>{
                let {code,msg,attributeValueID} = response.data;
                if ('success' == code){
                     this.$message({ message:msg, type: 'success' });
                }else{
                    this.$message({ message:msg, type: 'error' });
                }
            });
        },
        changeCreateForm(){
            this.$router.push("/product/attributevalue/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getAttributeValues(this.valuesData.pageSize,this.valuesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.valuesData.total/this.valuesData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
     
    
</style>
