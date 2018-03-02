<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="categoryForm" :model="categoryForm" :inline="true"  class="form-read-only">
                <el-form-item label="上级分类">
                    <span>{{categoryForm.Parent.Name}}</span>
                </el-form-item>
                <el-form-item label="分类名称">
                    <span>{{categoryForm.Name}}</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop'; 
     import  {SERVER_PRODUCT_CATEGORY} from '@/server_address';              
    import { mapState } from 'vuex';
    export default {
        data() {
            return {
                loadging:false,
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                categoryForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getProductCategoryInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.categoryForm.ID = id;
                this.$ajax.get(SERVER_PRODUCT_CATEGORY+this.categoryForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.categoryForm = data["category"];
                            this.access = data["access"];
                        }
                    });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/category");
                }else if ("form"==type){
                    this.$router.push("/product/category/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/product/category/form/"+this.categoryForm.ID);
            },
        },
        created:function(){
            this.getProductCategoryInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getProductCategoryInfo'
        },
         
    }
</script>