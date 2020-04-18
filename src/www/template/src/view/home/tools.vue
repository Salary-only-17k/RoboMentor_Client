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
                                <div class="hd">tcp/socket通讯</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="tcp/socket通讯" @click="onTools(4)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">陀螺仪</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="陀螺仪" @click="onTools(5)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">总线舵机</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="总线舵机" @click="onTools(6)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">PWM</div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" data-tools="PWM" @click="onTools(7)">打开工具</el-button>
                                </div>
                            </div>
                        </el-col>
                    </el-row>
                </div>
            </div>
        </div>

        <el-dialog :title="toolsTitle" :visible.sync="dialogTools" width="1000px" :before-close="HandleClose" :closeOnClickModal="false" :closeOnPressEscape="false">

            <div class="camera-box" v-if="toolsIndex === 1">
                <div class="no" v-if="!camera.Status">请在机器人主程序中开启摄像头</div>
                <img :src="camera.ReturnContent" v-else>
            </div>

            <el-form ref="form" :model="remote" label-width="0px" v-if="toolsIndex === 2">
                <el-form-item label="">
                    <el-input type="textarea" v-model="remote.content" placeholder="请输入要发送的消息数据" autocomplete="off" rows="10" resize="none" style="width: 100%;resize: none;"></el-input>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>根据自定义的消息通讯协议，填写要发送的数据。</span>
                </el-form-item>
                <el-form-item>
                    <el-button type="danger" :loading="remote.ButtonStatus" @click="onSubmitRemote">发送消息</el-button>
                </el-form-item>
            </el-form>

            <el-form ref="form" :model="serial" label-width="0px" v-if="toolsIndex === 3">
                <el-form-item label="">
                    <el-input v-model="serial.port" placeholder="请输入串口地址" autocomplete="off" style="width: 300px;"></el-input>
                    <el-input v-model="serial.rate" placeholder="波特率" autocomplete="off" style="width: 120px;"></el-input>
                    <el-input v-model="serial.bits" placeholder="数据位" autocomplete="off" style="width: 100px;"></el-input>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>硬件设备的串口地址、波特率、数据位，例如：/dev/ttyACM0。</span>
                </el-form-item>
                <el-form-item label="">
                    <el-input type="textarea" v-model="serial.content" placeholder="请输入要发送的串口数据" autocomplete="off" rows="10" resize="none" style="width: 100%;resize: none;"></el-input>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>根据串口协议，填写要发送的数据。</span>
                </el-form-item>
                <el-form-item>
                    <el-button type="danger" :loading="serial.ButtonStatus" @click="onSubmitSerial">发送串口数据</el-button>
                </el-form-item>
                <el-form-item label="">
                    <el-switch v-model="serial.Switch" active-color="#F56C6C" inactive-color="#30333d"></el-switch>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>是否开启串口数据监听，开启后，可以实时获取串口发送回来的数据信息。</span>
                </el-form-item>
                <el-form-item label="" v-if="serial.Switch">
                    <el-input type="textarea" v-model="serial.ReturnContent" placeholder="这里会实时显示串口发来的数据信息" autocomplete="off" rows="12" resize="none" style="width: 100%;resize: none;"></el-input>
                </el-form-item>
            </el-form>

            <el-form ref="form" :model="tcp" label-width="0px" v-if="toolsIndex === 4">
                <el-form-item label="">
                    <el-input v-model="tcp.ip" placeholder="请输入通讯IP地址" autocomplete="off" style="width: 300px;"></el-input>
                    <el-input v-model="tcp.port" placeholder="写入端口" autocomplete="off" style="width: 160px;"></el-input>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>tcp/socket通讯配置，包含IP、写入端口。</span>
                </el-form-item>
                <el-form-item label="">
                    <el-input type="textarea" v-model="tcp.content" placeholder="请输入要发送的tcp/socket数据，为空则视为只读数据" autocomplete="off" rows="10" resize="none" style="width: 100%;resize: none;"></el-input>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>根据tcp/socket协议，填写要发送的数据。</span>
                </el-form-item>
                <el-form-item>
                    <el-button type="danger" :loading="tcp.ButtonStatus" @click="onSubmitTcp">发送数据</el-button>
                </el-form-item>
                <el-form-item label="">
                    <el-switch v-model="tcp.Switch" active-color="#F56C6C" inactive-color="#30333d"></el-switch>
                    <span class="el-input__span"><i class="iconfont icon-feedback_fill"></i>是否开启串口数据监听，开启后，可以实时获取串口发送回来的数据信息。</span>
                </el-form-item>
                <el-form-item label="" v-if="tcp.Switch">
                    <el-input type="textarea" v-model="tcp.ReturnContent" placeholder="这里会实时显示tcp/socket发来的数据信息" autocomplete="off" rows="12" resize="none" style="width: 100%;resize: none;"></el-input>
                </el-form-item>
            </el-form>
        </el-dialog>
    </el-container>
</template>

<script>
    import{GetHomeTools,SetHomeToolsSerialSubmit,SetHomeToolsRemoteSubmit,SetHomeToolsTcpSubmit} from "../../api/index";

    export default {
        name: 'App',
        data() {
            return {
                Data:false,
                toolsTitle:"",
                dialogTools:false,
                toolsIndex:0,
                SocketStatus:true,
                camera:{
                    Status:false,
                    ReturnContent:""
                },
                remote:{
                    content:"",
                    ButtonStatus:false
                },
                serial:{
                    port:"/dev/ttyACM0",
                    rate:"115200",
                    bits:"8",
                    content:"",
                    Switch:false,
                    ReturnContent:"",
                    ButtonStatus:false
                },
                tcp:{
                    ip:"192.168.1.8",
                    port:"40923",
                    content:"",
                    Switch:false,
                    ReturnContent:"",
                    ButtonStatus:false
                }
            }
        },
        inject:['MessagesEmpty'],
        props:{
            message:{
                type:Object,
                default:""
            }
        },
        watch:{
            message:{
                handler() {
                    if(Object.keys(this.message).length !== 0) {
                        if (this.SocketStatus) {
                            this.SocketCallback(this.message);
                        }
                    }
                },
                deep: true
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
                if( tools === 1) {
                    this.toolsTitle = "远程控制";
                    this.dialogTools = true;
                    this.toolsIndex = tools;
                }else if(tools === 2){
                    this.toolsTitle = "远程消息";
                    this.dialogTools = true;
                    this.toolsIndex = tools;
                }else if(tools === 3){
                    this.toolsTitle = "串口调试";
                    this.dialogTools = true;
                    this.toolsIndex = tools;
                }else if(tools === 4){
                    this.toolsTitle = "tcp/socket通讯";
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
                this.serial = {port:"/dev/ttyACM0", rate:"115200", bits:"8", content:"", Switch:false, ReturnContent:"", ButtonStatus:false};
                this.remote = {content:"", ButtonStatus:false}
            },
            onSubmitSerial(){
                if(this.serial.port === "" || this.serial.rate === "" || this.serial.bits === "" || this.serial.content === ""){
                    this.$message({
                        message: '串口信息不完整，请补充完整',
                        type: 'warning'
                    });
                }else{
                    this.serial.ButtonStatus = true;
                    SetHomeToolsSerialSubmit(this.serial).then(res=>{
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
                        this.serial.ButtonStatus = false;
                    });
                }
            },
            onSubmitRemote(){
                if(this.remote.content === ""){
                    this.$message({
                        message: '消息内容不完整，请补充完整',
                        type: 'warning'
                    });
                }else{
                    this.remote.ButtonStatus = true;
                    SetHomeToolsRemoteSubmit(this.remote).then(res=>{
                        if(res.data.code === -1) {
                            this.$router.push({path: '/'});
                        }else if(res.data.code === 0){
                            this.$message({
                                message: '远程消息发送成功',
                                type: 'success'
                            });
                        }else{
                            this.$message.error(res.data.msg);
                        }
                        this.remote.ButtonStatus = false;
                    });
                }
            },
            onSubmitTcp(){
                if(this.tcp.ip === "" && this.tcp.content === "" && this.tcp.port === "") {
                    this.$message({
                        message: '通讯信息不完整，请补充完整',
                        type: 'warning'
                    });
                }else{
                    this.tcp.ButtonStatus = true;
                    SetHomeToolsTcpSubmit(this.tcp).then(res=>{
                        if(res.data.code === -1) {
                            this.$router.push({path: '/'});
                        }else if(res.data.code === 0){
                            this.$message({
                                message: '通讯数据发送成功',
                                type: 'success'
                            });
                        }else{
                            this.$message.error(res.data.msg);
                        }
                        this.tcp.ButtonStatus = false;
                    });
                }
            },
            SocketCallback(data){
                if(data.message_type === "serial_message" && this.serial.Switch){
                    this.serial.ReturnContent = data.serial_message.content;
                }

                if(data.message_type === "tcp_message" && this.tcp.Switch){
                    this.tcp.ReturnContent = data.tcp_message.content;
                }

                if(data.message_type === "camera_message"){
                    this.camera.Status = true;
                    this.camera.ReturnContent = data.camera_message.content;
                }
            },
        },
        beforeDestroy(){
            this.SocketStatus = false;
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
                            height: 30px;
                            line-height: 30px;
                            padding: 0;
                            font-size: 10px;
                            border-radius: 15px;
                            background-color: #30333d;
                            border: 0;
                            color: #D6D7DA;
                            &:hover{
                                color: #ffffff;
                                background-color: #f56c6c;
                            }
                        }
                    }
                }
            }
        }
    }
    .camera-box{
        width: 100%;
        height: 540px;
        background-color: #000000;
        position: relative;
        .no{
            width: 100%;
            height: 540px;
            position: absolute;
            top: 0;
            left: 0;
            text-align: center;
            line-height: 540px;
            font-size: 12px;
        }
        img{
            width: 100%;
            height: 540px;
        }
    }
</style>