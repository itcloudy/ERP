<template>
    <div class="login" :style="winSize">
        <el-row>
            <el-col :span="24">
                <div class="content">
                    <el-form label-position="left" label-width="0px" class="loginform" 
                    :style="formOffset"
                    :model='data'
                    :rules="rule_data"
                    ref='loginData'>
                        <h3 class="title"> <span>系统登录</span></h3>
                        <el-form-item
                                prop='username'>
                            <el-input type="text" auto-complete="off" placeholder="请输入账号"
                                      v-model='data.username'></el-input>
                        </el-form-item>

                        <el-form-item
                                prop='password'>
                            <el-input type="password" auto-complete="off" placeholder="请输入密码"
                                      v-model='data.password'
                                      @keyup.native.enter="onLogin('loginData')"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click='onLogin("loginData")'>登录</el-button>
                        </el-form-item>
                    </el-form>
                </div>
            </el-col>
        </el-row>
    </div>
</template>
<script>
    import {axiosAjaxLogin} from '../../api/global';
    export default {
        name: 'login',
        data() {
            return {
                winSize: {
                    width: '',
                    height: ''
                },
                formOffset: {
                    position: 'absolute',
                    left: '',
                    top: ''
                },
                login_actions: {
                    disabled: false
                },
                data: {
                    username: '',
                    password: '',
                    token: ''
                },
                rule_data:{
                    username: [{
                        validator:(rule, value, callback)=>{
                            if (value === '') {
                                callback(new Error('请输入用户名'));
                            } else {
                                if(/^[a-zA-Z0-9_-]{6,16}$/.test(value)){
                                    callback();
                                }else{
                                    callback(new Error('用户名至少6位,由大小写字母和数字,-,_,组成'));
                                }
                            }
                        },
                        trigger: 'blur'
                    }],
                    password: [{
					validator:(rule, value, callback)=>{
						if (value === '') {
							callback(new Error('请输入密码'));
						} else {
						    if(!(/^[a-zA-Z0-9_-]{6,16}$/.test(value))){
								callback(new Error('密码至少6位,由大小写字母和数字,-,_组成'));
							}else{
								callback();
                            }

						}
                    },
                    trigger: 'blur'
                }]
                },
            }
        },
        methods: {
            // 设置窗口大小
            setSize() {
                this.winSize.width = window.innerWidth + 'px';
                this.winSize.height = window.innerHeight +'px' ;
                this.formOffset.left = (parseInt(this.winSize.width) / 2 - 175) + 'px';
                this.formOffset.top = (parseInt(this.winSize.height) / 2 - 178) + 'px';
            },
            onLogin(ref){
                this.$refs[ref].validate((valid)=>{
                    console.log(valid);
                    if(valid){
                        let params = {
                            username: this.data.username,
                            password: this.data.password
                        };
                        this.$ajax.post('/login',params);
                        //登录成功跳转到首页
                        this.$router.push('/');
                    }
                });
            },
        },
        created() {
            this.setSize();
        },
        mounted() {
        }
    }
</script>
<style lang="scss" scoped>
.login {
    background : #1F2D3D;

    .loginform {
        box-shadow            : 0 0px 8px 0 rgba(0, 0, 0, 0.06), 0 1px 0px 0 rgba(0, 0, 0, 0.02);
        -webkit-border-radius : 5px;
        border-radius         : 5px;
        -moz-border-radius    : 5px;
        background-clip       : padding-box;
        margin-bottom         : 3rem;
        background-color      : #F9FAFC;
        border                : 2px solid #8492A6;
    }

    .title {
        margin      : 0px auto 40px auto;
        text-align  : center;
        color       : red;
        font-weight : normal;
        font-size   : 2rem;

        span {
            cursor : pointer;

            &.active {
                font-weight : bold;
                font-size   : 18px;
            }
        }
    }

    .loginform {
        width   : 350px;
        padding : 35px 35px 15px 35px;
    }
}
</style>
