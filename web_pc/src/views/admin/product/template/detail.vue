<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="templateForm" :model="templateForm" class="form-read-only">
                <el-form-item label="款式名称">
                    <span>{{templateForm.Name}}</span>
                </el-form-item>
                <el-form-item label="产品编码">
                    <span>{{templateForm.DefaultCode}}</span>
                </el-form-item>
                <el-form-item label="产品类别">
                    <span v-if="templateForm.Category">{{templateForm.Category.Name}}</span v-else><span>暂未定</span>
                </el-form-item>
                <el-form-item label="可销售">
                    <span v-if="templateForm.SaleOk">是</span ><span v-else>否</span>
                </el-form-item>
                <el-form-item label="有效">
                    <span v-if="templateForm.Active">是</span ><span v-else>否</span>
                </el-form-item>
                <el-form-item label="代售品">
                    <span v-if="templateForm.Rental">是</span ><span v-else>否</span>
                </el-form-item>
                <el-form-item label="描述">
                    <span>{{templateForm.Description}}</span>
                </el-form-item>
                <el-form-item label="销售描述">
                    <span>{{templateForm.DescriptionSale}}</span>
                </el-form-item>
                <el-form-item label="采购描述">
                    <span>{{templateForm.DescriptionPurchase}}</span>
                </el-form-item>
                <el-form-item label="成本价格">
                    <span>{{templateForm.StandardPrice}}</span>
                </el-form-item>
                <el-form-item label="标准重量">
                    <span>{{templateForm.StandardWeight}}</span>
                </el-form-item>
                <el-form-item label="第一销售单位">
                    <span v-if="templateForm.SecondSaleUom">{{templateForm.SecondSaleUom.Name}}</span v-else><span>暂未定</span>
                </el-form-item>
                <el-form-item label="第二销售单位">
                    <span v-if="templateForm.SecondSaleUom">{{templateForm.SecondSaleUom.Name}}</span v-else><span>暂未定</span>
                </el-form-item>
                <el-form-item label="第一采购单位">
                    <span v-if="templateForm.FirstPurchaseUom">{{templateForm.FirstPurchaseUom.Name}}</span v-else><span>暂未定</span>
                </el-form-item>
                <el-form-item label="第二采购单位">
                    <span v-if="templateForm.SecondPurchaseUom">{{templateForm.SecondPurchaseUom.Name}}</span v-else><span>暂未定</span>
                </el-form-item>
                <el-form-item label="产品类型">
                    <span>{{templateForm.ProductType}}</span>
                </el-form-item>
                <el-form-item label="规格创建方式">
                    <span>{{templateForm.ProductMethod}}</span>
                </el-form-item>
                
            </el-form>
            <el-tabs  v-model="tab_active" type="card" @tab-click="handleClick">
                <el-tab-pane label="属性明细" name="attribute_lines">
                    <el-table
                        ref="caseTable"
                        :data="attributeLines"
                        style="width: 100%">
                        <el-table-column
                            prop="Name"
                            label="属性">
                        </el-table-column>
                        <el-table-column
                            prop="Values"
                            label="属性值">
                        </el-table-column>
                        <el-table-column label="操作">
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="图片" name="template_images">
                </el-tab-pane>
                <el-tab-pane label="产品规格" name="template_products">
                </el-tab-pane>
            </el-tabs>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/admin/common/FormTop'; 
    import  {SERVER_PRODUCT_TEMPLATE} from '@/server_address';         
    import { mapState } from 'vuex';
    export default {
        data() {
            return {
                loadging:false,
                tab_active:"attribute_lines",
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                templateForm:{},
               

            }
        },
        components:{
           FormTop
        },
        methods:{
            showAddAttributeLineDailog(){
                this.formSave();
            },
             
            getProductTemplateInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.templateForm.ID = id;
                   
                    this.$ajax.get(SERVER_PRODUCT_TEMPLATE+this.templateForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.templateForm = data["template"];
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.templateForm = this.NewTemplateForm;
                }
            },
            
            
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/admin/product/template");
                }else if ("form"==type){
                    this.$router.push("/admin/product/template/form/"+id);
                }
            },
        },
        created:function(){
            this.getProductTemplateInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getProductTemplateInfo'
        },
         
    }
</script>
<style lang="scss" scoped>
    
    
</style>
