<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form ref="uomForm" :model="uomForm" :inline="true"  class="form-read-only">
                <el-form-item label="单位名称">
                    <span>{{uomForm.Name}}</span>
                </el-form-item>
                <el-form-item label="所属类别">
                    <span>{{uomForm.Category.Name}}</span>
                </el-form-item>
                <el-form-item label="比率"  v-if="uomForm.Type =='smaller'">
                    <span>{{uomForm.Factor}}</span>
                </el-form-item>
                <el-form-item label="更大比率"  v-if="uomForm.Type =='bigger'">
                    <span>{{uomForm.FactorInv}}</span>
                </el-form-item>
                <el-form-item label="舍入精度">
                    <span>{{uomForm.Rounding}}</span>
                </el-form-item>
                <el-form-item label="类型">
                    <span>{{uomForm.Type}}</span>
                </el-form-item>
                <el-form-item label="符号位置">
                    <span v-if="uomForm.Symbol">前</span ><span v-else>后</span>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';   
    import  {SERVER_PRODUCT_UOM} from '@/server_address';       
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
                uomForm:{}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getProductUomInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                this.uomForm.ID = id;
                this.$ajax.get(SERVER_PRODUCT_UOM+this.uomForm.ID).then(response=>{
                        this.loadging = false;
                        let {code,msg,data} = response.data;
                        if(code=='success'){
                            this.uomForm = data["uom"];
                            this.access = data["access"];
                        }
                    });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/uom");
                }else if ("form"==type){
                    this.$router.push("/product/uom/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/product/uom/form/"+this.uomForm.ID);
            },
        },
        created:function(){
            this.getProductUomInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getProductUomInfo'
        },
         
    }
</script>