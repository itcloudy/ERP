<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        :edit="edit"
        @formSave="formSave"
        @changeView="changeView"/>
        <div v-if="edit"  v-loading="loading">
            <el-form :inline="true"  ref="attributeForm" :rules="attributeFormRules" :model="attributeForm" label-width="80px">
                <el-form-item label="属性名称">
                    <el-input v-model="attributeForm.Name"></el-input>
                </el-form-item>
                <el-form-item label="属性编码">
                    <el-input v-model="attributeForm.Code"></el-input>
                </el-form-item>
                        
            </el-form>
        </div>
        <div v-else  v-loading="loading">
            <el-form ref="attributeForm" :model="attributeForm" :inline="true"  label-width="80px">
               
                <el-form-item label="属性名称">
                    <span>{{attributeForm.Name}}</span>
                </el-form-item>
                <el-form-item label="属性编码">
                    <span>{{attributeForm.Code}}</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '../global/FormTop';         
    import { mapState } from 'vuex';
    export default {
        data() {
            return {
                edit:false,
                loadging:false,
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                attributeForm:{
                    Name:"",
                    ID:0,
                     
                },
                NewAttributeForm:{
                    Name:"",
                    ID:0,
                    
                },
                attributeFormRules:{
                    Name:[
                        { required: true, message: '属性名称', trigger: 'blur' },
                    ]
                }
            }
        },
        components:{
           FormTop
        },
        methods:{
            getAttributeInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;

                if (id!='new'){
                    this.attributeForm.ID = id;
                    this.$ajax.get("/product/attribute/"+this.attributeForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.attributeForm = data["attribute"];
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.edit = true;
                    this.attributeForm = this.NewAttributeForm;
                }
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/admin/product/attribute");
                }else if ("form"==type){
                    this.$router.push("/admin/product/attribute/"+id);
                }
            },
            formEdit(){
                this.edit = true;
            },
            formSave(){
                 this.$refs["attributeForm"].validate((valid) => {
                    if (valid) {
                        this.edit = false;
                        // 判断是采用post(创建)还是put(更新)
                        if (this.attributeForm.ID>0){
                            this.$ajax.put("/admin/product/attribute/",this.attributeForm).then(response=>{
                                 console.log(response);
                                 let {code,msg,data} = response.data;
                                 
                            });
                        }else{
                            this.$ajax.post("/admin/product/attribute/",this.attributeForm).then(response=>{
                                 console.log(response);
                                 let {code,msg,data} = response.data;
                                 
                            });
                        }
                    } else {
                        this.$message({  message:"存在错误信息或信息不完整",   type: 'error' });
                        return false;
                    }
                });
                
            }
        },
        created:function(){
            this.getAttributeInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getAttributeInfo'
        },
         
    }
</script>