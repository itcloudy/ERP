<template>
    <div  class="navbar-container">
        
        <span  @click="showLeftSidebarClick"><img src="/static/pc/images/menu.png" alt="菜单显示切换"></span>
        <div class="navbar-right">
            <el-dropdown trigger="click" @command="handleCommand">
                <el-button type="primary" class="el-dropdown-link">
                    {{userinfo.NameZh}}<i class="el-icon-caret-bottom el-icon--right"></i>
                </el-button>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item command="go_front">进入前台</el-dropdown-item>
                    <el-dropdown-item command="self_info">个人信息</el-dropdown-item>
                    <el-dropdown-item command='logout'>注销</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>
<script>
    import { mapState } from 'vuex';
    import localStore from '../../utils/local_store';
    export default{
        name:"navbar",
        data(){
            return{
                
            }
        },
        methods:{
            showLeftSidebarClick:function(){
                 this.$store.dispatch('GLOBAL_TOGGLE_LEFT_SIDEBAR');
            },
            handleCommand(command){
                let id =  this.userinfo.ID;
                if (command =="self_info"){
                    this.$ajax.get("/user/"+id)
                }else if (command =='logout'){
                    
                    //发送请求到服务器，用于记录登出时间
                    this.$ajax.get('/login/'+id)
                    .then(response=>{
                        let {code,msg,data} = response.data;
                        if (code=='success'){
                            //提示
                            this.$message({
                                message:msg,
                                type: 'success'
                            });
                            this.$router.push('/login');
                        }else{
                            this.$message({
                                message:msg,
                                type: 'error'
                            });
                        }
                        
                    });
                }else if("go_front"==command){
                    this.$router.push("/");
                }
            }
        },
        computed:{
           ...mapState({
               userinfo: state => state.userinfo
           })
        }
    }
</script>
<style lang="scss" scoped>
    .navbar-container{
        background-color: #20A0FF;
        color: white;
        height:50px;
        line-height:50px;
        .navbar-right{
            float: right;
        }
    }
</style>