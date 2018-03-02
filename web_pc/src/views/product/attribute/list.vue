<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product' }">产品管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/product/attribute' }">属性</el-breadcrumb-item>
        </el-breadcrumb>
       <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
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
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {SERVER_PRODUCT_ATTRIBUTE} from '@/server_address';        
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
        getAttributes(limit,offset){
            this.loading = true;
            this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                    this.loading = false;
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
            this.getAttributes(pageSize,(currentPage-1)*pageSize)
        },
        goProductAttributeDetail(row, event){
            this.$router.push("/product/attribute/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/product/attribute/form/new");
        }
         
    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getAttributes(this.attributesData.pageSize,this.attributesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.attributesData.total/this.attributesData.pageSize > 1
        }
    }
      
    }
</script>
<style lang="scss" scoped>
     
    
</style>
