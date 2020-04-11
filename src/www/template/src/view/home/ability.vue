<template>
    <el-container>
        <div class="main">
            <div class="main-box">
                <el-card class="empty-box-card">没有找到相关数据</el-card>
            </div>
        </div>
    </el-container>
</template>

<script>
    import{GetHomeIndex} from "../../api/index";

    export default {
        name: 'App',
        data() {
            return {
                Data:false
            }
        },
        created(){
            this.init();
        },
        mounted() {

        },
        methods: {
            init(){
                GetHomeIndex().then(res=>{
                    if(res.data.code === -1){
                        this.$router.push({path: '/'});
                    }else{
                        this.Data = res.data.data;
                        this.$emit("main",this.Data);
                    }
                });
            }
        },
        beforeDestroy(){

        }
    }
</script>

<style lang="scss" scoped>
    .main{
        width: 100%;
        padding-top: 130px;
        .main-box{
            width: 1200px;
            min-height: 600px;
            margin: 0 auto;

        }
    }
</style>