<template>
    <el-container>
        <div class="main">
            <div class="main-box">
                <div class="tools-box">
                    <el-row :gutter="15">
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">远程控制</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="远程控制" @click="onTools(1)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">远程消息</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="远程消息" @click="onTools(2)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">串口调试</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="串口调试" @click="onTools(3)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">陀螺仪</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="陀螺仪" @click="onTools(4)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">总线舵机</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="总线舵机" @click="onTools(5)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">PWM</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="PWM" @click="onTools(6)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                    </el-row>
                </div>
            </div>
        </div>

        <el-dialog :title="toolsTitle" :visible.sync="dialogTools" width="1000px" :before-close="HandleClose" :closeOnClickModal="false" :closeOnPressEscape="false">
            <el-form ref="form" :model="serial" label-width="0px" v-if="toolsIndex === 3">
                <el-form-item label="">
                    <el-input v-model="serial.port" placeholder="请输入串口地址" autocomplete="off" style="width: 300px;"></el-input>
                    <el-input v-model="serial.rate" placeholder="波特率" autocomplete="off" style="width: 120px;"></el-input>
                    <el-input v-model="serial.bits" placeholder="数据位" autocomplete="off" style="width: 100px;"></el-input>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>硬件设备的串口地址、波特率、数据位，例如：/dev/tty.RoboMentors。</span>
                </el-form-item>
                <el-form-item label="">
                    <el-input type="textarea" v-model="serial.content" placeholder="请输入要发送的串口数据" autocomplete="off" rows="10" resize="none" style="width: 100%;resize: none;"></el-input>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>根据串口协议，填写要发送的数据。</span>
                </el-form-item>
                <el-form-item>
                    <el-button type="danger" :loading="ButtonStatus" @click="onSubmitSerial">发送串口数据</el-button>
                </el-form-item>
                <el-form-item label="">
                    <el-switch v-model="serial.Switch" active-color="#F56C6C" inactive-color="#30333d"></el-switch>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>是否开启串口数据监听，开启后，可以实时获取串口发送回来的数据信息。</span>
                </el-form-item>
                <el-form-item label="" v-if="serial.Switch">
                    <el-input type="textarea" v-model="serial.ReturnContent" placeholder="这里会实时显示串口发来的数据信息" autocomplete="off" rows="12" resize="none" style="width: 100%;resize: none;" disabled></el-input>
                </el-form-item>
            </el-form>
        </el-dialog>
    </el-container>
</template>

<script>
    import{GetHomeTools,SetHomeToolsSerial} from "../../api/index";

    export default {
        name: 'App',
        data() {
            return {
                Data:false,
                toolsTitle:"",
                dialogTools:false,
                toolsIndex:0,
                ButtonStatus:false,
                serial:{
                    port:"",
                    rate:"",
                    bits:"",
                    content:"",
                    Switch:false,
                    ReturnContent:"",
                    ButtonStatus:false
                }
            }
        },
        created(){
            this.init();
        },
        mounted() {

        },
        methods: {
            init(){
                GetHomeTools().then(res=>{
                    if(res.data.code === -1){
                        this.$router.push({path: '/'});
                    }else{
                        this.Data = res.data.data;
                        this.$emit("main",this.Data);
                    }
                });
            },
            onTools(tools){
                if( tools === 3){
                    this.toolsTitle = "串口调试";
                    this.dialogTools = true;
                    this.toolsIndex = tools;
                }else{
                    this.$message({
                        message: '功能被锁，请联系管理员',
                        type: 'warning'
                    });
                }
            },
            HandleClose(done){
                done();
                this.toolsIndex = 0;
                this.ButtonStatus = false;
                this.serial = {port:"", rate:"", bits:"", content:"", Switch:false, ReturnContent:"", ButtonStatus:false};
            },
            onSubmitSerial(){
                if(this.serial.port === "" || this.serial.rate === "" || this.serial.bits === "" || this.serial.content === ""){
                    this.$message({
                        message: '串口信息不完整，请补充完整',
                        type: 'warning'
                    });
                }else{
                    this.ButtonStatus = true;
                    SetHomeToolsSerial(this.serial).then(res=>{
                        if(res.data.code === -1) {
                            this.$router.push({path: '/'});
                        }else if(res.data.code === 0){
                            this.$message({
                                message: '串口数据发送成功',
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
        padding-top: 140px;
        .main-box{
            width: 1200px;
            min-height: 600px;
            margin: 0 auto;
            .tools-box{
                width: 100%;
                .grid-content{
                    width: 100%;
                    height: 200px;
                    line-height: 35px;
                    background-color: #333843;
                    border-radius: 6px;
                    border: 0;
                    box-shadow: 0 2px 2px 0 rgba(0,0,0,.1);
                    margin-bottom: 30px;
                    .hd{
                        width: 100%;
                        height: 100px;
                        background-color: #222222;
                        border-radius: 6px 6px 0 0;
                        text-align: center;
                        line-height: 100px;
                        margin-bottom: 35px;
                    }
                    .button-box{
                        width: 100%;
                        height: 35px;
                        text-align: center;
                        .tools-button{
                            width: 100px;
                            height: 35px;
                            line-height: 35px;
                            padding: 0;
                            font-size: 12px;
                        }
                    }
                }
            }
        }
    }
</style>