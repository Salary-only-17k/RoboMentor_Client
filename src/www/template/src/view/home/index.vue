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
        padding-top: 140px;
        .main-box{
            width: 1200px;
            min-height: 600px;
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
    }
</style>