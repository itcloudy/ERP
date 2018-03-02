<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="templateForm" :model="templateForm" :rules="templateFormRules"  label-width="80px">
                 
                <el-form-item label="款式名称" prop="Name">
                    <el-input v-model="templateForm.Name"></el-input>
                </el-form-item>
                <el-form-item label="产品编码" prop="DefaultCode">
                    <el-input v-model="templateForm.DefaultCode"></el-input>
                </el-form-item>
                <el-form-item label="产品类别" prop="Category">
                    <el-select
                        v-model="templateForm.Category.ID"
                        :name="templateForm.Category.Name"
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
                        v-model="templateForm.ProductType"
                        :name="productTypes[templateForm.ProductType['label']]"
                        placeholder="请输入产品类型">
                        <el-option
                            v-for="item in productTypes"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="规格创建方式"  prop="ProductMethod">
                    <el-select
                        v-model="templateForm.ProductMethod"
                        :name="productMethods[templateForm.ProductMethod['label']]"
                        placeholder="请输入规格创建方式">
                        <el-option
                            v-for="item in productMethods"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                
                <el-form-item label="描述">
                    <el-input v-model="templateForm.Description"></el-input>
                </el-form-item>
                <el-form-item label="销售描述">
                    <el-input v-model="templateForm.DescriptionSale"></el-input>
                </el-form-item>
                <el-form-item label="采购描述">
                    <el-input v-model="templateForm.DescriptionPurchase"></el-input>
                </el-form-item>
                <el-form-item label="成本价格">
                    <el-input v-model="templateForm.StandardPrice"></el-input>
                </el-form-item>
                <el-form-item label="标准重量">
                    <el-input v-model="templateForm.StandardWeight"></el-input>
                </el-form-item>
                <el-form-item label="第一销售单位" prop="FirstSaleUom">
                    <el-select
                        v-model="templateForm.FirstSaleUom.ID"
                        :name="templateForm.FirstSaleUom.Name"
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
                 <el-form-item label="第一采购单位" prop="FirstPurchaseUom">
                    <el-select
                        v-model="templateForm.FirstPurchaseUom.ID"
                        :name="templateForm.FirstPurchaseUom.Name"
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
                <el-form-item label="第二销售单位"  v-if="multiUnit">
                    <el-select
                        v-model="templateForm.SecondSaleUom.ID"
                        :name="templateForm.SecondSaleUom.Name"
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
               
                <el-form-item label="第二采购单位"  v-if="multiUnit">
                    <el-select
                        v-model="templateForm.SecondPurchaseUom.ID"
                        :name="templateForm.SecondPurchaseUom.Name"
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
                    <el-switch on-text="是" off-text="否" v-model="templateForm.SaleOk"></el-switch>
                </el-form-item>
                <el-form-item label="有效">
                    <el-switch on-text="是" off-text="否" v-model="templateForm.Active"></el-switch>
                </el-form-item>
                <el-form-item label="代售品">
                    <el-switch on-text="是" off-text="否" v-model="templateForm.Rental"></el-switch>
                </el-form-item>
                
            </el-form>
            <el-tabs  v-model="tab_active" type="card" @tab-click="handleClick">
                <el-tab-pane label="属性明细" name="attribute_lines">
                    <div class="form-tab-top-info">
                        <el-tooltip class="item" effect="dark" content="点击会先保存基本信息" placement="top-start">
                            <el-button type="primary" size="small" @click="showAddAttributeLineForm">添加属性明细</el-button><span></span>
                        </el-tooltip>
                    </div>
                    <el-table
                        ref="caseTable"
                        :data="attributeLines"
                        style="width: 100%">
                        <el-table-column
                            prop="ID"
                            label="ID">
                        </el-table-column>
                        <el-table-column
                            label="属性">
                            <template scope="scope">
                                <div slot="reference" class="name-wrapper">
                                    <el-tag>{{ scope.row.Attribute.Name }}</el-tag>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column
                            label="属性值">
                            <template scope="scope">
                                <span slot="reference" class="values-wrapper" :key="index" v-for="(attValue,index) in scope.row.AttributeValues">
                                    <el-tag>{{ attValue.Name }}</el-tag>
                                </span>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作">
                            <template scope="scope">
                                <el-button
                                type="info"
                                size="mini"
                                @click="updateProductAttributeLine(scope.$index, scope.row)">修改</el-button>
                                <el-button
                                type="danger"
                                size="mini"
                                @click="deleteProductAttributeLine(scope.$index, scope.row)">删除</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="图片" name="template_images">
                </el-tab-pane>
                <el-tab-pane label="产品规格" name="template_products">
                </el-tab-pane>
            </el-tabs>
        </div>
        <AttributeLineDialog v-if="dialogFormVisible"
        :dialogFormVisible="dialogFormVisible"
        :form="attributeLineForm"
        @dialogFormVisibleEn="dialogFormVisibleEn"/>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop';
    import  {default as AttributeLineDialog} from '@/views/product/template/attribute_line_dialog'; 
    import  {SERVER_PRODUCT_TEMPLATE,SERVER_PRODUCT_CATEGORY,SERVER_PRODUCT_UOM,
        SERVER_PRODUCT_ATTRIBUTE,SERVER_PRODUCT_ATTRIBUTE_VALUE,SERVER_PRODUCT_ATTRIBUTE_LINE} from '@/server_address';         
    import { mapState } from 'vuex';
    import {validateObjectID} from '@/utils/validators';
    export default {
        data() {
            return {
                dialogFormVisible:false,
                loadging:false,
                tab_active:"attribute_lines",
                access:{
                    Create:false,
                    Update:false,
                    Read:false,
                    Unlink:false,
                },
                templateForm:{},
                NewTemplateForm:{
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
                    IsProductVariant:true,
                    FirstSaleUom:{ID:"",Name:""},
                    SecondSaleUom:{ID:"",Name:""},
                    FirstPurchaseUom:{ID:"",Name:""},
                    SecondPurchaseUom:{ID:"",Name:""},
                    AttributeLines:[],
                    ProductVariants:[],
                    VariantCount:0,
                    Barcode:"",
                    DefaultCode:"",
                    ProductType:"stock",
                    ProductMethod:"hand"
                },
                categoryList:[],
                uomList:[],
                productTypes:[
                    {value:"stock",label:"库存商品"},
                    {value:"consume",label:"消耗品"},
                    {value:"service",label:"服务"}
                ],
                productMethods:[
                    {value:"hand",label:"手动"},
                    {value:"auto",label:"自动"}
                ],
                attributeLines:[],
               
                templateFormRules:{
                    Name:[
                        { required: true, message: '请输入款式名称', trigger: 'blur' }
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
                        { required: true, message: '请输入款式编码', trigger: 'blur' }
                    ],
                    ProductType:[
                        { required: true, message: '请选择产品类型', trigger: 'blur' }
                    ],
                    ProductMethod:[
                        { required: true, message: '请选择规格创建方式', trigger: 'blur' }
                    ],
                },
                attributeLineForm:{
                    ID:'',
                    ProductTemplate:{ID:"",Name:""},
                    Attribute:{ID:"",Name:""},
                    AttributeValues:[]
                }
            }
        },
        components:{
           FormTop,
           AttributeLineDialog
        },
        methods:{
            updateProductAttributeLine(index,row){
                this.attributeLineForm.Attribute = row.Attribute;
                this.attributeLineForm.AttributeValues = row.AttributeValues;
                this.attributeLineForm.ID = row.ID;
                this.attributeLineForm.ProductTemplate.ID = this.templateForm.ID;
                this.attributeLineForm.ProductTemplate.Name = this.templateForm.Name;
                this.dialogFormVisible = true;
            },
            deleteProductAttributeLine(index,row){
                this.$ajax.delete(SERVER_PRODUCT_ATTRIBUTE_LINE +row.ID).then(response=>{
                let {code,msg,lineID} = response.data;
                if ('success' == code){
                     this.$message({ message:msg, type: 'success' });
                }else{
                    this.$message({ message:msg, type: 'error' });
                }
            });
            },
            dialogFormVisibleEn(){
                this.dialogFormVisible = false;
            },
            showAddAttributeLineForm(){
                if (this.templateForm.ID<1){
                    this.$message({ message:"请先保存", type: 'warning' });
                }else{
                    this.attributeLineForm.ProductTemplate.ID = this.templateForm.ID;
                    this.attributeLineForm.ProductTemplate.Name = this.templateForm.Name;
                    this.dialogFormVisible = true;
                }
                 
            },
            formSave(){
                 this.$refs['templateForm'].validate((valid) => {
                    if (valid) {
                        if (this.templateForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_TEMPLATE+this.templateForm.ID ,this.templateForm).then(response=>{
                                let {code,msg,templateID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/template/detail/"+templateID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_TEMPLATE,this.templateForm).then(response=>{
                                let {code,msg,templateID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/template/detail/"+templateID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                });
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
                                this.categoryList= [this.templateForm.Category];
                                this.attributeLines = this.templateForm.attributeLines;
                                this.getInitUomList();
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.templateForm = this.NewTemplateForm;
                }
            },
            getInitUomList(){
                if (this.templateForm){
                    let uomDict = {};
                    if (this.templateForm.FirstPurchaseUom && this.templateForm.FirstPurchaseUom.ID >0){
                        uomDict[this.templateForm.FirstPurchaseUom.ID] = this.templateForm.FirstPurchaseUom;
                    }
                    if (this.templateForm.FirstSaleUom && this.templateForm.FirstSaleUom.ID >0){
                        uomDict[this.templateForm.FirstSaleUom.ID] = this.templateForm.FirstSaleUom;
                    }
                    if (this.templateForm.SecondPurchaseUom && this.templateForm.SecondPurchaseUom.ID >0){
                        uomDict[this.templateForm.SecondPurchaseUom.ID] = this.templateForm.SecondPurchaseUom;
                    }
                    if (this.templateForm.SecondSaleUom && this.templateForm.SecondSaleUom.ID >0){
                        uomDict[this.templateForm.SecondSaleUom.ID] = this.templateForm.SecondSaleUom;
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
                    this.$router.push("/product/template");
                }else if ("form"==type){
                    this.$router.push("/product/template/form/"+id);
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
        computed:{
            ...mapState({
                multiUnit: state => state.multiUnit
            })
        }
         
    }
</script>
<style lang="scss" scoped>
    .values-wrapper{
        padding: 0 2px;
    }
    
</style>
