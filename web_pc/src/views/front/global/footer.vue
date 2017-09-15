<template>
    <footer class="main-footer">
        <div >
            <b>Version</b> 1.0.1
            
            <strong>Copyright &copy; 2017-2018 <a href="http://hechihan.com">golangERP</a>.</strong> All rights
        reserved.
        </div>
  </footer>
</template>
<script>
    export default{
        name:"navbar",
        data(){
            return{
                
            }
        }
    }
</script>
<style lang="scss" scoped>
    .main-footer{
        position: fixed;
        bottom: 0;
        background: #fff;
        padding: 15px;
        color: #444;
        border-top: 1px solid #d2d6de;
        width: 100%;
        display: flex;
        justify-content: center;

    }
</style>
