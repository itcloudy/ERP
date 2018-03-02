<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="productForm" :model="productForm" :rules="productFormRules"  label-width="80px">
                
                <el-form-item label="产品款式" prop="Template">
                    <el-select
                        v-model="productForm.Template.ID"
                        :name="productForm.Template.Name"
                        filterable
                        remote
                        placeholder="请输入产品款式"
                        :remote-method="getProductTemplateList">
                        <el-option
                            v-for="item in templateList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="规格名称" prop="Name">
                    <el-input v-model="productForm.Name"></el-input>
                </el-form-item>
                <el-form-item label="产品编码" prop="DefaultCode">
                    <el-input v-model="productForm.DefaultCode"></el-input>
                </el-form-item>
                <el-form-item label="产品类别" prop="Category">
                    <el-select
                        v-model="productForm.Category.ID"
                        :name="productForm.Category.Name"
                        filterable
                        remote
                        placeholder="请输入产品类别"
                        :remote-method="getProductCategoryList">
                        <el-option
                            v-for="item in categoryList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="产品类型" prop="ProductType">
                    <el-select
                        v-model="productForm.ProductType"
                        :name="productTypes[productForm.ProductType['label']]"
                        placeholder="请输入产品类型">
                        <el-option
                            v-for="item in productTypes"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                 
                <el-form-item label="描述">
                    <el-input v-model="productForm.Description"></el-input>
                </el-form-item>
                <el-form-item label="销售描述">
                    <el-input v-model="productForm.DescriptionSale"></el-input>
                </el-form-item>
                <el-form-item label="采购描述">
                    <el-input v-model="productForm.DescriptionPurchase"></el-input>
                </el-form-item>
                <el-form-item label="成本价格">
                    <el-input v-model="productForm.StandardPrice"></el-input>
                </el-form-item>
                <el-form-item label="标准重量">
                    <el-input v-model="productForm.StandardWeight"></el-input>
                </el-form-item>
                <el-form-item label="第一销售单位" prop="FirstSaleUom">
                    <el-select
                        v-model="productForm.FirstSaleUom.ID"
                        :name="productForm.FirstSaleUom.Name"
                        filterable
                        remote
                        placeholder="请输入第一销售单位"
                        :remote-method="getProductUomList">
                        <el-option
                            v-for="item in uomList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第二销售单位">
                    <el-select
                        v-model="productForm.SecondSaleUom.ID"
                        :name="productForm.SecondSaleUom.Name"
                        filterable
                        remote
                        placeholder="请输入第二销售单位"
                        :remote-method="getProductUomList">
                        <el-option
                            v-for="item in uomList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第一采购单位" prop="FirstPurchaseUom">
                    <el-select
                        v-model="productForm.FirstPurchaseUom.ID"
                        :name="productForm.FirstPurchaseUom.Name"
                        filterable
                        remote
                        placeholder="请输入第一采购单位"
                        :remote-method="getProductUomList">
                        <el-option
                            v-for="item in uomList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第二采购单位">
                    <el-select
                        v-model="productForm.SecondPurchaseUom.ID"
                        :name="productForm.SecondPurchaseUom.Name"
                        filterable
                        remote
                        placeholder="请输入第二采购单位"
                        :remote-method="getProductUomList">
                        <el-option
                            v-for="item in uomList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="可销售">
                    <el-switch on-text="是" off-text="否" v-model="productForm.SaleOk"></el-switch>
                </el-form-item>
                <el-form-item label="有效">
                    <el-switch on-text="是" off-text="否" v-model="productForm.Active"></el-switch>
                </el-form-item>
                <el-form-item label="代售品">
                    <el-switch on-text="是" off-text="否" v-model="productForm.Rental"></el-switch>
                </el-form-item>
                
            </el-form>
            
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';
    import  {SERVER_PRODUCT_TEMPLATE,SERVER_PRODUCT_PRODUCT,SERVER_PRODUCT_CATEGORY,SERVER_PRODUCT_UOM,
        SERVER_PRODUCT_ATTRIBUTE,SERVER_PRODUCT_ATTRIBUTE_VALUE,SERVER_PRODUCT_ATTRIBUTE_LINE} from '@/server_address';         
    import { mapState } from 'vuex';
    import {validateObjectID} from '@/utils/validators';
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
                productTypes:[
                    {value:"stock",label:"库存商品"},
                    {value:"consume",label:"消耗品"},
                    {value:"service",label:"服务"}
                ],
                productMethods:[
                    {value:"hand",label:"手动"},
                    {value:"auto",label:"自动"}
                ],
                productForm:{},
                NewProductForm:{
                    Template:{ID:"",Name:""},
                    Name:"",
                    ID:"",
                    Description:"",
                    DescriptionSale:"",
                    DescriptionPurchase:"",
                    Rental:false,
                    Category:{ID:"",Name:""},
                    Price:0,
                    StandardPrice:0,
                    StandardWeight:0,
                    SaleOk:true,
                    Active:true,
                    FirstSaleUom:{ID:"",Name:""},
                    SecondSaleUom:{ID:"",Name:""},
                    FirstPurchaseUom:{ID:"",Name:""},
                    SecondPurchaseUom:{ID:"",Name:""},
                    DefaultCode:"",
                    ProductType:"stock",
                },
                categoryList:[],
                templateList:[],
                uomList:[],
                productFormRules:{
                    Template:[
                        { required: true, message: '请输入款式',validator: validateObjectID, trigger: 'blur' }
                    ],
                    Name:[
                        { required: true, message: '请输入规格名称', trigger: 'blur' }
                    ],
                    Category:[
                        { required: true, message: '请选择类别',validator: validateObjectID, trigger: 'blur' }
                    ],
                    FirstSaleUom:[
                        { required: true, message: '请选择第一销售单位',validator: validateObjectID, trigger: 'blur' }
                    ],
                    FirstPurchaseUom:[
                        { required: true, message: '请选择第一采购单位',validator: validateObjectID, trigger: 'blur' }
                    ],
                    DefaultCode:[
                        { required: true, message: '请输入编码', trigger: 'blur' }
                    ],
                },
                
            }
        },
        components:{
           FormTop,
        },
        methods:{
            formSave(){
                 this.$refs['productForm'].validate((valid) => {
                    if (valid) {
                        if (this.productForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_PRODUCT+this.productForm.ID ,this.productForm).then(response=>{
                                let {code,msg,productID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/product/detail/"+productID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_PRODUCT,this.productForm).then(response=>{
                                let {code,msg,productID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/product/detail/"+productID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                });
            },
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
                                this.categoryList= [this.productForm.Category];
                                this.attributeLines = this.productForm.attributeLines;
                                this.getInitUomList();
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.productForm = this.NewProductForm;
                }
            },
            getInitUomList(){
                if (this.productForm){
                    let uomDict = {};
                    if (this.productForm.FirstPurchaseUom && this.productForm.FirstPurchaseUom.ID >0){
                        uomDict[this.productForm.FirstPurchaseUom.ID] = this.productForm.FirstPurchaseUom;
                    }
                    if (this.productForm.FirstSaleUom && this.productForm.FirstSaleUom.ID >0){
                        uomDict[this.productForm.FirstSaleUom.ID] = this.productForm.FirstSaleUom;
                    }
                    if (this.productForm.SecondPurchaseUom && this.productForm.SecondPurchaseUom.ID >0){
                        uomDict[this.productForm.SecondPurchaseUom.ID] = this.productForm.SecondPurchaseUom;
                    }
                    if (this.productForm.SecondSaleUom && this.productForm.SecondSaleUom.ID >0){
                        uomDict[this.productForm.SecondSaleUom.ID] = this.productForm.SecondSaleUom;
                    }
                    for (let key in uomDict) {
                        this.uomList.push(uomDict[key]);
                    }
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
            getProductTemplateList(query){
                this.$ajax.get(SERVER_PRODUCT_TEMPLATE,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.templateList = data["templates"];
                    }
                });
            },
            getProductUomList(query){
                this.$ajax.get(SERVER_PRODUCT_UOM,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.uomList = data["uoms"];
                    }
                });
            },
            changeView(type,id){
                if ("list"==type){
                    this.$router.push("/product/product");
                }else if ("form"==type){
                    this.$router.push("/product/product/form/"+id);
                }
            },
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
</style>
