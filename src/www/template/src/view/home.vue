<template>
    <el-container>
        <el-main>
            <div class="header" v-if="Data">
                <div class="header-hd">
                    <div class="hd-box">
                        <div class="hd-box-item">
                          <div class="logo"></div>
                        </div>
                        <div class="hd-box-item right">
                            <div class="hd-tools-item">
                                <div class="switch-status"><i class="iconfont icon-coordinates_fill"></i> {{serviceStatus}}</div>
                            </div>
                            <div class="hd-tools-item">
                                <a href="https://www.robomentor.cn" target="_blank" class="home-run"><i class="iconfont icon-homepage_fill"></i> 官方网站</a>
                            </div>
                            <div class="hd-tools-item">
                                <a href="https://www.robomentor.cn/#/shop" target="_blank" class="home-run"><i class="iconfont icon-service_fill"></i> 硬件商城</a>
                            </div>
                            <div class="hd-tools-item">
                                <a href="http://wiki.robomentor.cn/#/developer" target="_blank" class="home-run"><i class="iconfont icon-document_fill"></i> 开发者社区</a>
                            </div>
                            <div class="hd-tools-item">
                                <a href="http://wiki.robomentor.cn" target="_blank" class="home-run"><i class="iconfont icon-document_fill"></i> 开发文档</a>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="header-nav">
                    <div class="header-nav-box">
                        <el-row :gutter="15">
                            <el-col :span="3">
                                <router-link to="/" exact>
                                    <div class="grid-content"><i class="iconfont icon-homepage_fill"></i> 控制台首页</div>
                                </router-link>
                            </el-col>
                            <el-col :span="3">
                                <router-link to="/robot">
                                    <div class="grid-content"><i class="iconfont icon-integral_fill"></i> 机器人设置</div>
                                </router-link>
                            </el-col>
                            <el-col :span="3">
                                <router-link to="/skill">
                                    <div class="grid-content"><i class="iconfont icon-flashlight_fill"></i> 技能管理</div>
                                </router-link>
                            </el-col>
                            <el-col :span="3">
                                <router-link to="/ability">
                                    <div class="grid-content"><i class="iconfont icon-manage_fill"></i> 能力管理</div>
                                </router-link>
                            </el-col>
                            <el-col :span="3">
                                <router-link to="/tools">
                                    <div class="grid-content"><i class="iconfont icon-setup_fill"></i> 调试工具</div>
                                </router-link>
                            </el-col>
                        </el-row>
                    </div>
                </div>
            </div>
            <div class="container">
                <router-view class="router-view" v-on:main="Data_In" :message="message"></router-view>
            </div>
            <div class="error-box" v-if="Data && !Data.config.robot_auth.status">
                <div class="error-content">
                    <div class="icon"><i class="iconfont icon-message_fill"></i></div>
                    <div class="title">机器人设备尚未联网，请先进行联网设置</div>
                    <div class="describe">RoboMentorSDK涉及到很多云端机器人能力，需要在网络环境下使用，请按照开发手册进行网络的设置，联网成功后请重新机器人设备。</div>
                    <div class="button">
                        <a href="http://wiki.robomentor.cn/1579502" target="_blank">查看联网设置</a>
                    </div>
                </div>
            </div>
        </el-main>
    </el-container>
</template>

<script>
    const socket = new WebSocket("ws://"+window.location.hostname+":8888/message");
    export default {
        name: 'App',
        data() {
            return {
                Data:false,
                message:{},
                serviceStatus:"正在诊断服务"
            }
        },
        provide(){
            return{
                MessagesEmpty:this.MessagesEmpty
            }
        },
        created(){

        },
        methods: {
            Data_In(data){
                if(!this.Data){
                    this.Data = data;
                    this.onSocket();
                }
            },
            onSocket(){
                let that = this;
                socket.onopen = function(){
                    that.serviceStatus = "服务连接正常";
                    socket.onmessage = function (message){
                        that.message = JSON.parse(message.data.toString());
                    };
                };
                socket.onerror = function () {
                    that.serviceStatus = "服务连接异常";
                };
                socket.onclose = function () {
                    that.serviceStatus = "服务连接中断";
                };
            },
            MessagesEmpty(){
                this.message = {};
            },
        },
        beforeDestroy(){

        }
    }
</script>

<style lang="scss" scoped>
  .el-main{
      padding: 0;
      .header{
          width: 100%;
          height: 100px;
          position: fixed;
          top: 0;
          left: 0;
          right: 0;
          z-index: 1000;
          .header-hd{
              width: 100%;
              height: 50px;
              background-color: #202124;
              .hd-box{
                  width: 1200px;
                  height: 50px;
                  font-size: 0;
                  margin: 0 auto;
                  .hd-box-item{
                      width: auto;
                      height: 50px;
                      display: inline-block;
                      vertical-align: top;
                      .logo{
                          width: 150px;
                          height: 50px;
                          background: url("../assets/images/logo.png") no-repeat center left;
                          background-size: 95%;
                      }
                      &.right{
                          float: right;
                          font-size: 12px;
                          .hd-tools-item{
                              min-width: 120px;
                              height: 30px;
                              display: inline-block;
                              vertical-align: top;
                              border-radius: 15px;
                              padding: 0 10px;
                              background-color: #30333d;
                              margin-left: 10px;
                              margin-top: 10px;
                              line-height: 30px;
                              text-align: center;
                              font-size: 12px;
                              color: #606266;
                              cursor: pointer;
                              a{
                                  color: #606266;
                                  &:hover{
                                      color: #D6D7DA;
                                  }
                              }
                              &:hover{
                                  color: #D6D7DA;
                              }
                          }
                      }
                  }
              }
          }
          .header-nav{
              width: 100%;
              height: 50px;
              background-color: #333843;
              box-shadow: 0 1px 6px rgba(0,0,0,.05);
              .header-nav-box{
                  width: 1200px;
                  height: 50px;
                  margin: 0 auto;
                  a{
                      width: 100%;
                      .grid-content{
                          width: auto;
                          height: 35px;
                          line-height: 35px;
                          cursor: pointer;
                          background-color: #30333d;
                          border-radius: 17.5px;
                          padding: 0 15px;
                          margin-top: 7.5px;
                          text-align: center;
                      }
                      &.router-link-active,&.router-link-active:hover,&:hover{
                          .grid-content{
                              background-color: #409EFF;
                              color: #ffffff;
                          }
                      }
                  }
              }
          }
      }
      .error-box{
          width: 100%;
          height: 100%;
          position: fixed;
          z-index: 1000;
          left: 0;
          bottom: 0;
          right: 0;
          top: 0;
          background: #30333d;
          .error-content{
              width: 600px;
              height: 300px;
              background-color: #333843;
              margin: 0 auto;
              padding: 20px;
              margin-top: 250px;
              .icon{
                  width: 80px;
                  height: 80px;
                  margin: 0 auto;
                  text-align: center;
                  line-height: 80px;
                  .iconfont{
                      font-size: 24px;
                      color: #f56c6c;
                  }
              }
              .title{
                  width: 100%;
                  height: 35px;
                  line-height: 35px;
                  font-size: 14px;
                  text-align: center;
                  margin-bottom: 15px;
              }
              .describe{
                  width: 500px;
                  height: 48px;
                  line-height: 24px;
                  margin: 0 auto;
                  color: #555d71;
                  margin-bottom: 15px;
              }
              .button{
                  width: 100%;
                  height: 35px;
                  a{
                      width: 120px;
                      height: 35px;
                      border-radius: 17.5px;
                      padding: 0 10px;
                      background-color: #30333d;
                      line-height: 35px;
                      text-align: center;
                      font-size: 12px;
                      color: #606266;
                      cursor: pointer;
                      display: block;
                      margin: 0 auto;
                      font-weight: 400;
                      &:hover{
                          color: #ffffff;
                          background-color: #f56c6c;;
                      }
                  }
              }
          }
      }
  }
</style>
