<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="valueForm" :model="valueForm" :inline="true"  class="form-read-only">
                <el-form-item label="属性">
                    <span>{{valueForm.Attribute.Name}}</span>
                </el-form-item>
                <el-form-item label="属性值">
                    <span>{{valueForm.Name}}</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';   
    import  {SERVER_PRODUCT_ATTRIBUTE_VALUE} from '@/server_address';              
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
                valueForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getAttributeValueInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.valueForm.ID = id;
                this.$ajax.get(SERVER_PRODUCT_ATTRIBUTE_VALUE+this.valueForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.valueForm = data["attributeValue"];
                            this.access = data["access"];
                        }
                    });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/attributevalue");
                }else if ("form"==type){
                    this.$router.push("/product/attributevalue/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/product/attributevalue/form/"+this.valueForm.ID);
            },
        },
        created:function(){
            this.getAttributeValueInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getAttributeValueInfo'
        },
         
    }
</script>