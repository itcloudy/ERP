<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="uomcategForm" :model="uomcategForm" :inline="true"  class="form-read-only">
                <el-form-item label="类别名称">
                    <span>{{uomcategForm.Name}}</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';   
    import  {SERVER_PRODUCT_UOM_CATEG} from '@/server_address';    
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
                uomcategForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getUomCategInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.uomcategForm.ID = id;
                this.$ajax.get(SERVER_PRODUCT_UOM_CATEG+this.uomcategForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.uomcategForm = data["uomcateg"];
                            this.access = data["access"];
                        }
                    });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/uomcateg");
                }else if ("form"==type){
                    this.$router.push("/product/uomcateg/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/product/uomcateg/form/"+this.uomcategForm.ID);
            },
        },
        created:function(){
            this.getUomCategInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getUomCategInfo'
        },
         
    }
</script>