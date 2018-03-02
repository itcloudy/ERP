<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="categoryForm" :model="categoryForm" :rules="categoryFormRules" label-width="80px">
                <el-form-item label="上级分类">
                    <el-select
                        v-model="categoryForm.Parent.ID"
                        :name="categoryForm.Parent.Name"
                        filterable
                        remote
                        placeholder="请输入上级分类"
                        :remote-method="getProductCategoryList">
                        <el-option
                            v-for="item in categoryList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                 
                <el-form-item label="分类名称" prop="Name">
                    <el-input v-model="categoryForm.Name"></el-input>
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
                categoryForm:{  },
                NewCategoryForm:{
                    Name:"",
                    ID:"",
                    Parent:{
                        Name:"",
                        ID:"",
                    },
                },
                categoryList:[],
                categoryFormRules:{
                    Name:[
                        { required: true, message: '请输入产品分类名称', trigger: 'blur' }
                    ]
                }
                
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                this.$refs['categoryForm'].validate((valid) => {
                    if (valid) {
                        if (this.categoryForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_CATEGORY+this.categoryForm.ID ,this.categoryForm).then(response=>{
                                let {code,msg,categoryID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/category/detail/"+categoryID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_CATEGORY,this.categoryForm).then(response=>{
                                let {code,msg,categoryID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/category/detail/"+categoryID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                });
            
            },
            getProductCategoryInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.categoryForm.ID = id;
                    this.$ajax.get(SERVER_PRODUCT_CATEGORY+this.categoryForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.categoryForm = data["category"];
                                this.categoryList = [this.categoryForm.Parent];
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.categoryForm = this.NewCategoryForm;
                }
            },
            getProductCategoryList(query){
                this.$ajax.get(SERVER_PRODUCT_CATEGORY,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.categoryList = data["categories"];
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