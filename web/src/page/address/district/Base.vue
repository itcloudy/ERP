<template>
    <div>
        <district-tree v-show="showTree" @changeViewType="changeViewType" @pageInfoChange="pageInfoChange" :districtsData="districtsData"></district-tree>
        <district-form v-show="showForm" @changeViewType="changeViewType" :district="district"></district-form>
    </div>
    
</template>

<script>
    import  {default as districtTree} from './Tree';
    import  {default as districtForm} from './Form';
    export default {
        data() {
            
            return {
                showTree:true,//显示tree视图
                showForm:false,//显示form视图
                districtsData:{
                    districtList:[],//tree视图数据
                    pageSize:20,//每页数量
                    total:0,//总数量
                    currentPage:1,//当前页
                },
                
                district:{}

            }
        },
        components: {
            districtTree,
            districtForm,
        },
        methods:{
            changeViewType(type){
                if ('form'==type){
                    this.showTree = false
                    this.showForm = true;
                }else if ('tree'== type){
                    this.showTree = true
                    this.showForm = false;
                }else{
                    this.showTree = true
                    this.showForm = false;
                }
            },
            getDistricts(limit,offset){
                this.$ajax.get("/address/district/?limit="+limit +"&offset="+offset).then(response=>{
                   let {code,msg,data} = response.data;
                   if(code=='success'){
                        this.districtsData.districtList = data["districts"];
                        let paginator = data.paginator;
                        if (paginator){
                            this.districtsData.total = paginator.totalCount;

                        }
                       
                   }
                });
            },
            pageInfoChange(pageSize,currentPage){
                this.districtsData.pageSize = pageSize;
                this.districtsData.currentPage = currentPage;
                this.getDistricts(pageSize,(currentPage-1)*pageSize)
            }
        },
        created:function(){
            this.$nextTick(function(){
                this.getDistricts(this.districtsData.pageSize,this.districtsData.currentPage-1);
            });
        }

      
    }
</script>