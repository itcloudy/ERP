<template>
    <div>
        <form-top  :Update="access.Update" :Create="access.Create" 
        @formSave="formSave"
        :edit="true"
        @changeView="changeView"/>
        <div v-loading="loading">
            <el-form :inline="true" ref="uomForm" :model="uomForm"  :rules="uomFormRules"label-width="80px">
                 <el-form-item label="所属类别" prop="Category">
                    <el-select
                        v-model="uomForm.Category.ID"
                        :name="uomForm.Category.Name"
                        filterable
                        remote
                        placeholder="请选择类别"
                        :remote-method="getUomCategList">
                        <el-option
                            v-for="item in categoryList"
                            :key="item.ID"
                            :label="item.Name"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="单位名称" prop="Name">
                    <el-input v-model="uomForm.Name"></el-input>
                </el-form-item>
                 <el-form-item label="类型"  prop="Type">
                    <el-select
                        v-model="uomForm.Type"
                        :name="uomTypes[uomForm.Type['label']]"
                        placeholder="请选择类型">
                        <el-option
                            v-for="item in uomTypes"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
               
                <el-form-item label="舍入精度" >
                    <el-input v-model="uomForm.Rounding"></el-input>
                </el-form-item>
                <el-form-item label="符号位置">
                    <el-switch on-text="前置" off-text="后置" v-model="uomForm.Symbol"></el-switch>
                </el-form-item>
                <el-form-item label="比率" v-show="uomForm.Type == 'smaller'">
                    <el-input v-model="uomForm.Factor"></el-input>
                </el-form-item>
                <el-form-item label="更大比率"  v-show="uomForm.Type == 'bigger'">
                    <el-input v-model="uomForm.Factor"></el-input>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import  {default as FormTop} from '@/views/common/FormTop'; 
    import  {SERVER_PRODUCT_UOM,SERVER_PRODUCT_UOM_CATEG} from '@/server_address'; 
    import {validateObjectID} from '@/utils/validators';       
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
                uomForm:{
                    Name:"",
                    ID:"",
                },
                NewUomForm:{
                    Name:"",
                    ID:"",
                    Category:{
                        Name:"",
                        ID:"",
                    },
                    Type:"reference",
                    Factor:1,
                    FactorInv:1,
                    Rounding:0.01,
                    Symbol:false
                },
                uomTypes:[
                     {value:"reference",label:"参考单位"},
                     {value:"bigger",label:"大于参考单位"},
                     {value:"smaller",label:"小于参考单位"},
                ],
                categoryList:[],
                uomFormRules:{
                    Name:[
                        { required: true, message: '请输入单位名称', trigger: 'blur' }
                    ],
                    Category:[
                        { required: true, message: '请选择类别', validator: validateObjectID, trigger: 'blur' }
                    ],
                    Type:[
                        { required: true, message: '请选择类型', trigger: 'blur' }
                    ],
                }
            }
        },
        components:{
           FormTop
        },
        methods:{
            formSave(){
                this.$refs['uomForm'].validate((valid) => {
                    if (valid) {
                        if (this.uomForm.ID >0){
                            this.$ajax.put(SERVER_PRODUCT_UOM + this.uomForm.ID ,this.uomForm).then(response=>{
                                let {code,msg,uomID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/uom/detail/"+uomID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }else{
                            this.$ajax.post(SERVER_PRODUCT_UOM,this.uomForm).then(response=>{
                                let {code,msg,uomID} = response.data;
                                if(code=='success'){
                                    this.$message({ message:msg, type: 'success' });
                                    this.$router.push("/product/uom/detail/"+uomID);
                                }else{
                                    this.$message({ message:msg, type: 'error' });
                                }
                            });
                        }
                    }
                });
            },
            getUomInfo(){
                this.loadging = true;
                let id  = this.$route.params.id;
                if (id!='new'){
                    this.uomForm.ID = id;
                    this.$ajax.get(SERVER_PRODUCT_UOM+this.uomForm.ID).then(response=>{
                            this.loadging = false;
                            let {code,msg,data} = response.data;
                            if(code=='success'){
                                this.uomForm = data["uom"];
                                this.categoryList = [this.uomForm.Category]
                                this.access = data["access"];
                            }
                        });
                }else{
                    this.uomForm = this.NewUomForm;
                }
            },
            getUomCategList(query){
                this.$ajax.get(SERVER_PRODUCT_UOM_CATEG,{
                    params:{
                        offset:0,
                        limit:20,
                        name:query,
                    }
                }).then(response=>{
                    let {code,msg,data} = response.data;
                    if(code=='success'){
                        this.categoryList = data["uomcategs"];
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
        },
        created:function(){
            this.getUomInfo();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'getUomInfo'
        },
         
    }
</script>