<template>
    <el-container>
        <div class="main" v-if="Data">
            <div class="main-box">
                <el-card class="robot-box-card">
                    <div class="title">机器人基础配置</div>
                    <el-form ref="form" :model="Data" label-width="150px">
                        <el-form-item label="机器人名称">
                            <el-input v-model="Data.config.robot_name" placeholder="请输入机器人名称" autocomplete="off" maxlength="18" style="width: 320px;"></el-input>
                            <div class="el-input__span"><i class="iconfont icon-feedback_fill"></i>用于区分账号下不同的机器人，修改后实时更新到机器人设备。</div>
                        </el-form-item>
                    </el-form>
                </el-card>
                <el-card class="robot-box-card">
                    <div class="title">机器人网络信息</div>
                    <el-form ref="form" :model="Data" label-width="150px">
                        <el-form-item label="机器人外网IP地址">
                            <el-input v-model="Data.config.robot_net_ip" placeholder="请输入机器人IP地址" autocomplete="off" maxlength="50" style="width: 240px;" disabled></el-input>
                            <div class="el-input__span"><i class="iconfont icon-feedback_fill"></i>系统自动获取，无需手动设置，请勿自行修改，否则导致机器人故障。</div>
                        </el-form-item>
                        <el-form-item label="机器人内网IP地址">
                            <el-input v-model="Data.config.robot_ip" placeholder="请输入机器人IP地址" autocomplete="off" maxlength="50" style="width: 240px;" disabled></el-input>
                            <div class="el-input__span"><i class="iconfont icon-feedback_fill"></i>系统自动获取，无需手动设置，请勿自行修改，否则导致机器人故障。</div>
                        </el-form-item>
                        <el-form-item label="机器人MAC地址">
                            <el-input v-model="Data.config.robot_mac" placeholder="请输入机器人MAC地址" autocomplete="off" maxlength="50" style="width: 300px;" disabled></el-input>
                            <div class="el-input__span"><i class="iconfont icon-feedback_fill"></i>系统自动获取，无需手动设置，请勿自行修改，否则导致机器人故障。</div>
                        </el-form-item>
                    </el-form>
                </el-card>
                <el-card class="robot-box-card">
                    <div class="title">机器人主控开发板配置</div>
                    <el-form ref="form" :model="Data" label-width="150px">
                        <el-form-item label="主控开发板串口">
                            <el-input v-model="Data.config.robot_board.port" placeholder="请输入串口地址" autocomplete="off" style="width: 300px;"></el-input>
                            <el-input v-model="Data.config.robot_board.rate" placeholder="请输入波特率" autocomplete="off" style="width: 120px;"></el-input>
                            <el-input v-model="Data.config.robot_board.bits" placeholder="请输入数据位" autocomplete="off" style="width: 120px;"></el-input>
                        </el-form-item>
                        <el-form-item label="">
                            <div class="el-input__span"><i class="iconfont icon-feedback_fill"></i>硬件设备的串口地址、波特率、数据位，请确保机器人主控开发板已经连接到机器人设备，修改后重启机器人设备生效。</div>
                        </el-form-item>
                    </el-form>
                </el-card>
                <el-card class="robot-box-card">
                    <el-form ref="form" :model="Data" label-width="150px">
                        <el-form-item>
                            <el-button type="danger" :loading="ButtonStatus" @click="onSubmitRobot">保存设置</el-button>
                        </el-form-item>
                    </el-form>
                </el-card>
            </div>
        </div>
    </el-container>
</template>

<script>
    import{GetHomeRobot,SetHomeRobotSubmit} from "../../api/index";

    export default {
        name: 'App',
        data() {
            return {
                Data:false,
                ButtonStatus:false
            }
        },
        created(){
            this.init();
        },
        mounted() {

        },
        methods: {
            init(){
                GetHomeRobot().then(res=>{
                    if(res.data.code === -1){
                        this.$router.push({path: '/'});
                    }else{
                        this.Data = res.data.data;
                        this.$emit("main",this.Data);
                    }
                });
            },
            onSubmitRobot(){
                if(this.Data.config.robot_name === ""){
                    this.$message({
                        message: '机器人名称不能为空，请补充完整',
                        type: 'warning'
                    });
                }else{
                    this.ButtonStatus = true;
                    SetHomeRobotSubmit(this.Data.config.robot_name, this.Data.config.robot_board.port, this.Data.config.robot_board.rate, this.Data.config.robot_board.bits).then(res=>{
                        if(res.data.code === -1) {
                            this.$router.push({path: '/'});
                        }else if(res.data.code === 0){
                            this.$message({
                                message: '机器人配置保存成功',
                                type: 'success'
                            });
                        }else{
                            this.$message.error(res.data.msg);
                        }
                        this.ButtonStatus = false;
                    });
                }
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
            .robot-box-card{
                width: 100%;
                background-color: #333843;
                border: 0;
                box-shadow: 0 2px 2px 0 rgba(0,0,0,.1);
                margin-bottom: 30px;
                .title{
                    width: 100%;
                    height: 40px;
                    line-height: 40px;
                    font-size: 14px;
                    color: #D6D7DA;
                    margin-bottom: 15px;
                }
            }
        }
    }
</style>