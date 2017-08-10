<template>
    <div>
        <province-tree v-show="showTree" @changeViewType="changeViewType" @pageInfoChange="pageInfoChange" :provincesData="provincesData"></province-tree>
        <province-form v-show="showForm" @changeViewType="changeViewType" :province="province"></province-form>
    </div>
    
</template>

<script>
    import  {default as provinceTree} from './Tree';
    import  {default as provinceForm} from './Form';
    export default {
        data() {
            
            return {
                showTree:true,//显示tree视图
                showForm:false,//显示form视图
                provincesData:{
                    provinceList:[],//tree视图数据
                    pageSize:20,//每页数量
                    total:0,//总数量
                    currentPage:1,//当前页
                },
                
                province:{}

            }
        },
        components: {
            provinceTree,
            provinceForm,
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
            getProvinces(limit,offset){
                this.$ajax.get("/address/province/?limit="+limit +"&offset="+offset).then(response=>{
                   let {code,msg,data} = response.data;
                   if(code=='success'){
                        this.provincesData.provinceList = data["provinces"];
                        let paginator = data.paginator;
                        if (paginator){
                            this.provincesData.total = paginator.totalCount;

                        }
                       
                   }
                });
            },
            pageInfoChange(pageSize,currentPage){
                this.provincesData.pageSize = pageSize;
                this.provincesData.currentPage = currentPage;
                this.getProvinces(pageSize,(currentPage-1)*pageSize)
            }
        },
        created:function(){
            this.$nextTick(function(){
                this.getProvinces(this.provincesData.pageSize,this.provincesData.currentPage-1);
            });
        }

      
    }
</script>