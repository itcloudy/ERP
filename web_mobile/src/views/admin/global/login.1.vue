<template>
    <div>
        <div class="page-login">
            <head-top :head-title="'登录'" />
        </div>
        <el-form :model="loginForm" id="login-form" :rules="formRules" ref="loginForm">  
            
            <el-form-item   prop="username">
                <el-input v-model="loginForm.username" placeholder="请输入账号"></el-input>
            </el-form-item>
            <el-form-item   prop="password">  
                <el-input v-model="loginForm.password" placeholder="请输入密码"></el-input>
            </el-form-item>
            
        </el-form>
        <div class="login_container" @click="mobileLogin">登录</div>
    </div>
</template>
<script>
    import {default as headTop} from '../../common/header'
    export default {
    	data(){
            return{
                loginForm:{
                    username:"admin",
                    password:"cloudy"
                },
                formRules:{
                    username:[
                        { required: true, message: '请输入账号', trigger: 'blur' },
                    ],
                    password:[
                        { required: true, message: '请输入密码', trigger: 'blur' },
                    ],
                }
            }
        },
        props: ['headTitle'],
        components: {
            headTop,
        },
         
    }
</script>
<style lang="scss" scoped>
    @import '../../../styles/mixin';
    .page-title{
       
        position: fixed;
        z-index: 100;
        top: 0;  
        h3{
            text-align:center;
        }      
    }
    #login-form{
        margin-top:3rem;
    }
    .login_container{
        margin: 0 .5rem 1rem;
        @include fontSizeColor(.7rem, #fff);
        background-color: #4cd964;
        padding: .5rem 0;
        border: 1px;
        border-radius: 0.15rem;
        text-align: center;
    }
</style>