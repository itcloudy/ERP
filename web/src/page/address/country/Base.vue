<template>
    <div>
        <country-tree v-show="showTree" @changeViewType="changeViewType" @pageInfoChange="pageInfoChange" :countriesData="countriesData"></country-tree>
        <country-form v-show="showForm" @changeViewType="changeViewType" :country="country"></country-form>
    </div>
    
</template>

<script>
    import  {default as countryTree} from './Tree';
    import  {default as countryForm} from './Form';
    export default {
        data() {
            
            return {
                showTree:true,//显示tree视图
                showForm:false,//显示form视图
                countriesData:{
                    countryList:[],//tree视图数据
                    pageSize:20,//每页数量
                    total:0,//总数量
                    currentPage:1,//当前页
                },
                
                country:{}

            }
        },
        components: {
            countryTree,
            countryForm,
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
            getCountries(limit,offset){
                this.$ajax.get("/address/country/?limit="+limit +"&offset="+offset).then(response=>{
                   let {code,msg,data} = response.data;
                   if(code=='success'){
                        this.countriesData.countryList = data["countries"];
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
            }
        },
        created:function(){
            this.$nextTick(function(){
                this.getCountries(this.countriesData.pageSize,this.countriesData.currentPage-1);
            });
        }

      
    }
</script>