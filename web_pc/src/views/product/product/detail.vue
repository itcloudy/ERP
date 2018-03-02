<product>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formEdit="formEdit"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="productForm" :model="productForm" class="form-read-only">
                <el-form-item label="款式名称">
                    <span>{{productForm.Name}}</span>
                </el-form-item>
                <el-form-item label="产品编码">
                    <span>{{productForm.DefaultCode}}</span>
                </el-form-item>
                <el-form-item label="产品类别">
                    <span v-if="productForm.Category">{{productForm.Category.Name}}</span><span  v-else>暂未定</span>
                </el-form-item>
                <el-form-item label="可销售">
                    <span v-if="productForm.SaleOk">是</span ><span v-else>否</span>
                </el-form-item>
                <el-form-item label="有效">
                    <span v-if="productForm.Active">是</span ><span v-else>否</span>
                </el-form-item>
                <el-form-item label="代售品">
                    <span v-if="productForm.Rental">是</span ><span v-else>否</span>
                </el-form-item>
                <el-form-item label="描述">
                    <span>{{productForm.Description}}</span>
                </el-form-item>
                <el-form-item label="销售描述">
                    <span>{{productForm.DescriptionSale}}</span>
                </el-form-item>
                <el-form-item label="采购描述">
                    <span>{{productForm.DescriptionPurchase}}</span>
                </el-form-item>
                <el-form-item label="成本价格">
                    <span>{{productForm.StandardPrice}}</span>
                </el-form-item>
                <el-form-item label="标准重量">
                    <span>{{productForm.StandardWeight}}</span>
                </el-form-item>
                <el-form-item label="第一销售单位">
                    <span v-if="productForm.SecondSaleUom">{{productForm.SecondSaleUom.Name}}</span ><span v-else>暂未定</span>
                </el-form-item>
                <el-form-item label="第二销售单位">
                    <span v-if="productForm.SecondSaleUom">{{productForm.SecondSaleUom.Name}}</span ><span v-else>暂未定</span>
                </el-form-item>
                <el-form-item label="第一采购单位">
                    <span v-if="productForm.FirstPurchaseUom">{{productForm.FirstPurchaseUom.Name}}</span ><span v-else>暂未定</span>
                </el-form-item>
                <el-form-item label="第二采购单位">
                    <span v-if="productForm.SecondPurchaseUom">{{productForm.SecondPurchaseUom.Name}}</span ><span v-else>暂未定</span>
                </el-form-item>
                <el-form-item label="产品类型">
                    <span v-if="productForm.ProductType">{{productType[productForm.ProductType]}}</span><span v-else>暂未定</span>
                </el-form-item>
                <el-form-item label="规格创建方式">
                    <span v-if="productForm.ProductMethod">{{productMethod[productForm.ProductMethod]}}</span><span v-else>暂未定</span>
                </el-form-item>
                
            </el-form>
        </div>
    </div>
</product>
<script>
    import  {default as FormTop} from '@/views/common/FormTop'; 
    import  {SERVER_PRODUCT_PRODUCT} from '@/server_address';         
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
                productForm:{},
                productType:{"stock":"库存商品","consume":"消耗品","service":"服务"},
                productMethod:{"hand":"手动","auto":"自动"}
            }
        },
        components:{
           FormTop
        },
        methods:{
            getProductProductInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.productForm.ID = id;
                   
                    this.$ajax.get(SERVER_PRODUCT_PRODUCT+this.productForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.productForm = data["product"];
                                this.attributeLines = this.productForm.attributeLines;
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.productForm = this.NewProductForm;
                }
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/product");
                }else if ("form"==type){
                    this.$router.push("/product/product/form/"+id);
                }
            },
            formEdit(){
                 this.$router.push("/product/product/form/"+this.productForm.ID);
            }
        },
        created:function(){
            this.getProductProductInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getProductProductInfo'
        },
         
    }
</script>
<style lang="scss" scoped>
    .values-wrapper{
        padding: 0 2px;
    }
    
</style>
