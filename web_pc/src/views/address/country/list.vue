<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '' }">后台首页</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address' }">地址管理</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/address/country' }">国家</el-breadcrumb-item>
        </el-breadcrumb>
        <div>
            <ListTop :Create="access.Create" @changeCreateForm="changeCreateForm" />
            <pagination 
            @pageInfoChange="pageInfoChange"
            :pageSize="countriesData.pageSize" 
            :currentPage="countriesData.currentPage"
             
            :total="countriesData.total"/>
            <el-table
                ref="multipleTable"
                :data="countriesData.countryList"
                @row-dblclick = "goCountryDetail"
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
                label="国家">
                </el-table-column>
            </el-table>
            
        </div>
    </div>
</template>
<script>
    import  {default as ListTop} from '@/views/common/ListTop'; 
    import  {default as Pagination} from '@/views/common/Pagination';
    import  {SERVER_ADDRESS_COUNTRY} from '@/server_address';
    import { mapState } from 'vuex';
    export default {
      data() {
        return {
            treeViewHeight: this.$store.state.windowHeight-100,
            countriesData:{
                countryList:[],//tree视图数据
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
            
        }
    },
    methods:{
        getCountries(limit,offset){
            this.$ajax.get(SERVER_ADDRESS_COUNTRY,{
                    params:{
                        offset:offset,
                        limit:limit
                    }
                }).then(response=>{
                let {code,msg,data} = response.data;
                if(code=='success'){
                    this.countriesData.countryList = data["countries"];
                    this.access = data["access"];
                    let paginator = data.paginator;
                    if (paginator){
                        this.countriesData.total = paginator.totalCount;
                    }
                }
            });
        },
        pageInfoChange(pageSize,currentPage){
            this.countriesData.pageSize = pageSize;
            this.countriesData.currentPage = currentPage;
            this.getCountries(pageSize,(currentPage-1)*pageSize)
        },
        goCountryDetail(row, event){
            this.$router.push("/address/country/detail/"+row.ID);
        },
        changeCreateForm(){
            this.$router.push("/address/country/form/new");
        }

    },
    components: {
        Pagination,
        ListTop
    },
    created:function(){
        this.$nextTick(function(){
            this.getCountries(this.countriesData.pageSize,this.countriesData.currentPage-1);
        });
    },
    computed:{
        showBottomPagitator:function(){
            return this.countriesData.total/this.countriesData.pageSize > 1
        }
    }
      
    }
</script>