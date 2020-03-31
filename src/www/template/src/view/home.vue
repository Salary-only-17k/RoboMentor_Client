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
                <router-view class="router-view" v-on:main="Data_In" v-bind:data="Data" :message="message"></router-view>
            </div>
        </el-main>
    </el-container>
</template>

<script>
    const socket = new WebSocket("ws://"+window.location.host+"/message");
    export default {
        name: 'App',
        data() {
            return {
                Data:false,
                message:{}
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
                    socket.onmessage = function (message){
                        that.message = JSON.parse(message.data.toString());
                    }
                }
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
          height: 110px;
          position: fixed;
          top: 0;
          left: 0;
          right: 0;
          z-index: 1000;
          .header-hd{
              width: 100%;
              height: 60px;
              background-color: #202124;
              .hd-box{
                  width: 1200px;
                  height: 60px;
                  font-size: 0;
                  margin: 0 auto;
                  .hd-box-item{
                      width: auto;
                      height: 60px;
                      display: inline-block;
                      vertical-align: top;
                      .logo{
                          width: 150px;
                          height: 60px;
                          background: url("../assets/images/logo.png") no-repeat center left;
                          background-size: 95%;
                      }
                      &.right{
                          text-align: right;
                          font-size: 12px;
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
  }
</style>
