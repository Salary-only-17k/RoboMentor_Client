<template>
    <el-container>
        <div class="main" v-if="Data">
            <div class="main-box">
                <el-card class="system-box-card">
                    <div class="system-box">
                        <div class="chart-box">
                            <div class="title">实时性能趋势图</div>
                            <dashboard5-chart6 :cData="ChartData"/>
                        </div>
                    </div>
                </el-card>
            </div>
            <div class="main-footer">
                <el-card class="footer-box-card">
                    <div class="footer-item">
                        <el-form ref="form" label-width="80px">
                            <el-form-item label="系统信息">{{Data.goos}} {{Data.goarch}} {{Data.version}}</el-form-item>
                            <el-form-item label="设备名">{{Data.host_name}}</el-form-item>
                            <el-form-item label="版本信息">RoboMentor SDK {{Data.config.robot_version}}</el-form-item>
                            <el-form-item label="运行目录">{{Data.config.robot_path}}</el-form-item>
                            <el-form-item label="Mac地址">{{Data.config.robot_mac}}</el-form-item>
                            <el-form-item label="Ip地址">{{Data.config.robot_ip}}</el-form-item>
                            <el-form-item class="end" label="授权编号">{{Data.config.token}}</el-form-item>
                        </el-form>
                    </div>
                    <div class="footer-item">
                        <div class="wechat">
                            <img src="../../assets/images/qrcode_for_gh_a2c92d370ff8_258.jpg">
                        </div>
                        <div class="wechat-tips">微信扫码 关注官方公众号</div>
                    </div>
                </el-card>
            </div>
        </div>
    </el-container>
</template>

<script>
    import{GetHomeIndex} from "../../api/index";
    import Vue from 'vue';
    import ChartJs from 'vue-chartjs';

    Vue.component('dashboard5-chart6', {
        extends: ChartJs.Line,
        props:["cData"],
        watch:{
            cData:{
                handler(val){
                    this.renderChart({
                        labels:val.labels,
                        datasets:[
                            {
                                label: 'CPU使用率',
                                borderColor: '#F56C6C',
                                pointBackgroundColor: 'white',
                                borderWidth: 1,
                                pointBorderColor: 'white',
                                data: val.data1
                            },{
                                label: '内存使用率',
                                borderColor: '#409EFF',
                                pointBackgroundColor: 'white',
                                pointBorderColor: 'white',
                                borderWidth: 1,
                                data: val.data2
                            }
                        ]
                    }, {
                        scales: {
                            yAxes: [{
                                ticks: {
                                    min: 0,
                                    max: 100
                                },
                                display: true
                            }]
                        },
                        responsive: false,
                        maintainAspectRatio: false
                    })
                    this._data._chart.resize();
                },
                deep:true
            }
        }
    });

    export default {
        name: 'App',
        extends: ChartJs.Line,
        data() {
            return {
                Data:false,
                SocketStatus:true,
                Timer:false,
                ChartData:{
                    labels:[],
                    data1:[],
                    data2:[]
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
                    if(Object.keys(this.message).length !== 0){
                        if(this.SocketStatus && this.message.message_type === "system"){
                            this.SocketCallback(this.message.system_message);
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
                GetHomeIndex().then(res=>{
                    if(res.data.code === -1){
                        this.$router.push({path: '/'});
                    }else{
                        this.Data = res.data.data;
                        this.$emit("main",this.Data);
                    }
                });
            },
            SocketCallback(data){
                if(this.ChartData.labels.length === 10){
                    this.ChartData.labels.shift();
                    this.ChartData.data1.shift();
                    this.ChartData.data2.shift();
                    this.ChartData.labels.push(data.time);
                    var sum = 0;
                    for (var i in data.cpu) {
                        sum += data.cpu[i];
                    }
                    this.ChartData.data1.push(parseFloat(sum).toFixed(2));
                    this.ChartData.data2.push(parseFloat(data.memory).toFixed(2));
                }else{
                    this.ChartData.labels.push(data.time);
                    var sum = 0;
                    for (var i in data.cpu) {
                        sum += data.cpu[i];
                    }
                    this.ChartData.data1.push(parseFloat(sum).toFixed(2));
                    this.ChartData.data2.push(parseFloat(data.memory).toFixed(2));
                }
            }
        },
        beforeDestroy(){
            if(this.Timer) {
                clearInterval(this.Timer);
                this.Timer = false;
            }
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
            margin: 0 auto;
            .system-box-card{
                width: 100%;
                background-color: #333843;
                border: 0;
                box-shadow: 0 2px 2px 0 rgba(0,0,0,.1);
                margin-bottom: 30px;
                .system-box{
                    width: 100%;
                    color: #555d71;
                    .chart-box{
                        background-color: #30333d;
                        padding: 20px;
                        border-radius: 6px;
                        .title{
                            width: 100%;
                            height: 40px;
                            line-height: 40px;
                            font-size: 14px;
                            color: #D6D7DA;
                        }
                    }
                }
            }
        }
        .main-footer{
            width: 1200px;
            min-height: 200px;
            margin: 0 auto;
            .footer-box-card{
                width: 100%;
                background-color: #333843;
                border: 0;
                box-shadow: 0 2px 2px 0 rgba(0,0,0,.1);
                margin-bottom: 30px;
                font-size: 0;
                .footer-item{
                    width: 50%;
                    display: inline-block;
                    vertical-align: top;
                    color: #D6D7DA;
                    font-size: 12px;
                    .el-form-item{
                        position: relative;
                        margin-bottom: 10px;
                        .el-form-item__content{
                            font-size: 12px;
                        }
                        &:after {
                            width: 100%;
                            content: " ";
                            position: absolute;
                            left: 80px;
                            bottom: 0;
                            right: 0;
                            height: 1px;
                            border-bottom: 1px solid rgba(0,0,0,.2);
                            color: rgba(0,0,0,.1);
                            transform-origin: 0 100%;
                            transform: scaleY(.5);
                        }
                        &.end{
                            user-select:text;
                            margin-bottom: 0;
                            &:after {
                                width: 0;
                            }
                        }
                    }
                    .wechat{
                        width: 160px;
                        height: 220px;
                        margin: 0 auto;
                        padding-top: 60px;
                        margin-bottom: 20px;
                        img{
                            width: 160px;
                            height: 160px;
                        }
                    }
                    .wechat-tips{
                        width: 200px;
                        height: 35px;
                        text-align: center;
                        line-height: 35px;
                        font-size: 12px;
                        margin: 0 auto;
                        background-color: #30333d;
                        border-radius: 17.5px;
                    }
                }
            }
        }
    }
</style>