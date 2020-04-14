webpackJsonp([1],{"4/hK":function(t,s){},BQy5:function(t,s){},FRBu:function(t,s){},FWHM:function(t,s){},LxZp:function(t,s){},NHnr:function(t,s,e){"use strict";Object.defineProperty(s,"__esModule",{value:!0});var a=e("7+uW"),o={render:function(){var t=this.$createElement,s=this._self._c||t;return s("div",{attrs:{id:"app"}},[s("router-view")],1)},staticRenderFns:[]};var i=e("VU/8")({name:"App"},o,!1,function(t){e("FRBu")},null,null).exports,n=e("/ocq");const l=new WebSocket("ws://"+window.location.hostname+":8888/message");var r={name:"App",data:()=>({Data:!1,message:{},serviceStatus:"正在诊断服务"}),provide(){return{MessagesEmpty:this.MessagesEmpty}},created(){},methods:{Data_In(t){this.Data||(this.Data=t,this.onSocket())},onSocket(){let t=this;l.onopen=function(){t.serviceStatus="服务连接正常",l.onmessage=function(s){t.message=JSON.parse(s.data.toString())}},l.onerror=function(){t.serviceStatus="服务连接异常"},l.onclose=function(){t.serviceStatus="服务连接中断"}},MessagesEmpty(){this.message={}}},beforeDestroy(){}},c={render:function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("el-container",[e("el-main",[t.Data?e("div",{staticClass:"header"},[e("div",{staticClass:"header-hd"},[e("div",{staticClass:"hd-box"},[e("div",{staticClass:"hd-box-item"},[e("div",{staticClass:"logo"})]),t._v(" "),e("div",{staticClass:"hd-box-item right"},[e("div",{staticClass:"hd-tools-item"},[e("div",{staticClass:"switch-status"},[e("i",{staticClass:"iconfont icon-coordinates_fill"}),t._v(" "+t._s(t.serviceStatus))])]),t._v(" "),e("div",{staticClass:"hd-tools-item"},[e("a",{staticClass:"home-run",attrs:{href:"https://www.robomentor.cn",target:"_blank"}},[e("i",{staticClass:"iconfont icon-homepage_fill"}),t._v(" 官方网站")])]),t._v(" "),e("div",{staticClass:"hd-tools-item"},[e("a",{staticClass:"home-run",attrs:{href:"https://www.robomentor.cn/#/developer/edu",target:"_blank"}},[e("i",{staticClass:"iconfont icon-video_fill"}),t._v(" 机甲学院")])]),t._v(" "),e("div",{staticClass:"hd-tools-item"},[e("a",{staticClass:"home-run",attrs:{href:"https://www.robomentor.cn/#/developer/shop",target:"_blank"}},[e("i",{staticClass:"iconfont icon-service_fill"}),t._v(" 硬件商城")])]),t._v(" "),e("div",{staticClass:"hd-tools-item"},[e("a",{staticClass:"home-run",attrs:{href:"http://www.robomentor.cn/#/developer",target:"_blank"}},[e("i",{staticClass:"iconfont icon-document_fill"}),t._v(" 开发者社区")])]),t._v(" "),e("div",{staticClass:"hd-tools-item"},[e("a",{staticClass:"home-run",attrs:{href:"http://wiki.robomentor.cn",target:"_blank"}},[e("i",{staticClass:"iconfont icon-document_fill"}),t._v(" 开发文档")])])])])]),t._v(" "),e("div",{staticClass:"header-nav"},[e("div",{staticClass:"header-nav-box"},[e("el-row",{attrs:{gutter:15}},[e("el-col",{attrs:{span:3}},[e("router-link",{attrs:{to:"/",exact:""}},[e("div",{staticClass:"grid-content"},[e("i",{staticClass:"iconfont icon-homepage_fill"}),t._v(" 控制台首页")])])],1),t._v(" "),e("el-col",{attrs:{span:3}},[e("router-link",{attrs:{to:"/robot"}},[e("div",{staticClass:"grid-content"},[e("i",{staticClass:"iconfont icon-integral_fill"}),t._v(" 机器人设置")])])],1),t._v(" "),e("el-col",{attrs:{span:3}},[e("router-link",{attrs:{to:"/skill"}},[e("div",{staticClass:"grid-content"},[e("i",{staticClass:"iconfont icon-flashlight_fill"}),t._v(" 技能管理")])])],1),t._v(" "),e("el-col",{attrs:{span:3}},[e("router-link",{attrs:{to:"/ability"}},[e("div",{staticClass:"grid-content"},[e("i",{staticClass:"iconfont icon-manage_fill"}),t._v(" 能力管理")])])],1),t._v(" "),e("el-col",{attrs:{span:3}},[e("router-link",{attrs:{to:"/tools"}},[e("div",{staticClass:"grid-content"},[e("i",{staticClass:"iconfont icon-setup_fill"}),t._v(" 调试工具")])])],1)],1)],1)])]):t._e(),t._v(" "),e("div",{staticClass:"container"},[e("router-view",{staticClass:"router-view",attrs:{message:t.message},on:{main:t.Data_In}})],1),t._v(" "),t.Data&&!t.Data.config.robot_auth.status?e("div",{staticClass:"error-box"},[e("div",{staticClass:"error-content"},[e("div",{staticClass:"icon"},[e("i",{staticClass:"iconfont icon-message_fill"})]),t._v(" "),e("div",{staticClass:"title"},[t._v("机器人设备尚未联网，请先进行联网设置")]),t._v(" "),e("div",{staticClass:"describe"},[t._v("RoboMentorSDK涉及到很多云端机器人能力，需要在网络环境下使用，请按照开发手册进行网络的设置，联网成功后请重新机器人设备。")]),t._v(" "),e("div",{staticClass:"button"},[e("a",{attrs:{href:"http://wiki.robomentor.cn/1579502",target:"_blank"}},[t._v("查看联网设置")])])])]):t._e()])],1)},staticRenderFns:[]};var d=e("VU/8")(r,c,!1,function(t){e("bP0V")},"data-v-2d4e0288",null).exports,u=e("mtWM"),h=e.n(u),m=e("zL8q"),p=e.n(m);let v;const f=h.a.create({baseURL:"/"});function b(){return f({url:"common/home/index",method:"GET",data:""})}function _(t){return f({url:"common/home/skill/edit?type="+t,method:"GET",data:""})}f.interceptors.request.use(t=>(t.headers={"Content-Type":"application/json"},t),t=>Promise.reject(t)),f.interceptors.response.use(t=>t,t=>{if(t&&t.response){switch(t.response.status){case 400:case 401:case 403:case 404:case 408:case 500:case 504:break;default:t.message="未知错误"}v=m.Message.error(t.message)}else v=m.Message.error(t.message);return{error:-1,data:""}});var g=e("UlOv");a.default.component("dashboard5-chart6",{extends:g.a.Line,props:["cData"],watch:{cData:{handler(t){this.renderChart({labels:t.labels,datasets:[{label:"CPU使用率",borderColor:"#F56C6C",pointBackgroundColor:"white",borderWidth:1,pointBorderColor:"white",data:t.data1},{label:"内存使用率",borderColor:"#409EFF",pointBackgroundColor:"white",pointBorderColor:"white",borderWidth:1,data:t.data2}]},{scales:{yAxes:[{ticks:{min:0,max:100},display:!0}]},responsive:!1,maintainAspectRatio:!1}),this._data._chart.resize()},deep:!0}}});var C={name:"App",extends:g.a.Line,data:()=>({Data:!1,SocketStatus:!0,Timer:!1,ChartData:{labels:[],data1:[],data2:[]}}),created(){this.init(),this.GetSystemInfo()},mounted(){this.Timer=setInterval(this.GetSystemInfo,2e3)},methods:{init(){b().then(t=>{-1===t.data.code?this.$router.push({path:"/"}):(this.Data=t.data.data,this.$emit("main",this.Data))})},GetSystemInfo(){f({url:"common/home/index/system",method:"GET",data:""}).then(t=>{if(0===t.data.code)if(10===this.ChartData.labels.length){this.ChartData.labels.shift(),this.ChartData.data1.shift(),this.ChartData.data2.shift(),this.ChartData.labels.push(t.data.data.time);var s=0;for(var e in t.data.data.cpu)s+=t.data.data.cpu[e];this.ChartData.data1.push(parseFloat(s).toFixed(2)),this.ChartData.data2.push(parseFloat(t.data.data.memory).toFixed(2))}else{this.ChartData.labels.push(t.data.data.time);s=0;for(var e in t.data.data.cpu)s+=t.data.data.cpu[e];this.ChartData.data1.push(parseFloat(s).toFixed(2)),this.ChartData.data2.push(parseFloat(t.data.data.memory).toFixed(2))}})}},beforeDestroy(){this.Timer&&(clearInterval(this.Timer),this.Timer=!1),this.SocketStatus=!1}},x={render:function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("el-container",[t.Data?a("div",{staticClass:"main"},[a("div",{staticClass:"main-box"},[a("el-card",{staticClass:"system-box-card"},[a("div",{staticClass:"system-box"},[a("div",{staticClass:"chart-box"},[a("div",{staticClass:"title"},[t._v("实时性能趋势图")]),t._v(" "),a("dashboard5-chart6",{attrs:{cData:t.ChartData}})],1)])])],1),t._v(" "),a("div",{staticClass:"main-footer"},[a("el-card",{staticClass:"footer-box-card"},[a("div",{staticClass:"footer-item"},[a("el-form",{ref:"form",attrs:{"label-width":"80px"}},[a("el-form-item",{attrs:{label:"系统信息"}},[t._v(t._s(t.Data.goos)+" "+t._s(t.Data.goarch)+" "+t._s(t.Data.version))]),t._v(" "),a("el-form-item",{attrs:{label:"设备名"}},[t._v(t._s(t.Data.host_name))]),t._v(" "),a("el-form-item",{attrs:{label:"版本信息"}},[t._v("RoboMentor SDK "+t._s(t.Data.config.robot_version))]),t._v(" "),a("el-form-item",{attrs:{label:"运行目录"}},[t._v(t._s(t.Data.config.robot_path))]),t._v(" "),a("el-form-item",{attrs:{label:"Mac地址"}},[t._v(t._s(t.Data.config.robot_mac))]),t._v(" "),a("el-form-item",{attrs:{label:"Ip地址"}},[t._v(t._s(t.Data.config.robot_ip))]),t._v(" "),a("el-form-item",{staticClass:"end",attrs:{label:"授权编号"}},[t._v(t._s(t.Data.config.token))])],1)],1),t._v(" "),a("div",{staticClass:"footer-item"},[a("div",{staticClass:"wechat"},[a("img",{attrs:{src:e("oP5K")}})]),t._v(" "),a("div",{staticClass:"wechat-tips"},[t._v("微信扫码 关注官方公众号")])])])],1)]):t._e()])},staticRenderFns:[]};var k=e("VU/8")(C,x,!1,function(t){e("bf6V")},"data-v-41dcf22a",null).exports,j={name:"App",data:()=>({Data:!1,ButtonStatus:!1}),created(){this.init()},mounted(){},methods:{init(){f({url:"common/home/robot",method:"GET",data:""}).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):(this.Data=t.data.data,this.$emit("main",this.Data))})},onSubmitRobot(){var t,s,e,a;""===this.Data.config.robot_name?this.$message({message:"机器人名称不能为空，请补充完整",type:"warning"}):(this.ButtonStatus=!0,(t=this.Data.config.robot_name,s=this.Data.config.robot_board.port,e=this.Data.config.robot_board.rate,a=this.Data.config.robot_board.bits,f({url:"common/home/robot/submit?name="+t+"&port="+s+"&rate="+e+"&bits="+a,method:"GET",data:""})).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):0===t.data.code?this.$message({message:"机器人配置保存成功",type:"success"}):this.$message.error(t.data.msg),this.ButtonStatus=!1}))}},beforeDestroy(){}},S={render:function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("el-container",[t.Data?e("div",{staticClass:"main"},[e("div",{staticClass:"main-box"},[e("el-card",{staticClass:"robot-box-card"},[e("div",{staticClass:"title"},[t._v("机器人基础配置")]),t._v(" "),e("el-form",{ref:"form",attrs:{model:t.Data,"label-width":"150px"}},[e("el-form-item",{attrs:{label:"机器人名称"}},[e("el-input",{staticStyle:{width:"320px"},attrs:{placeholder:"请输入机器人名称",autocomplete:"off",maxlength:"18"},model:{value:t.Data.config.robot_name,callback:function(s){t.$set(t.Data.config,"robot_name",s)},expression:"Data.config.robot_name"}}),t._v(" "),e("div",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("用于区分账号下不同的机器人，修改后实时更新到机器人设备。")])],1)],1)],1),t._v(" "),e("el-card",{staticClass:"robot-box-card"},[e("div",{staticClass:"title"},[t._v("机器人网络信息")]),t._v(" "),e("el-form",{ref:"form",attrs:{model:t.Data,"label-width":"150px"}},[e("el-form-item",{attrs:{label:"机器人外网IP地址"}},[e("el-input",{staticStyle:{width:"240px"},attrs:{placeholder:"请输入机器人IP地址",autocomplete:"off",maxlength:"50",disabled:""},model:{value:t.Data.config.robot_net_ip,callback:function(s){t.$set(t.Data.config,"robot_net_ip",s)},expression:"Data.config.robot_net_ip"}}),t._v(" "),e("div",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("系统自动获取，无需手动设置，请勿自行修改，否则导致机器人故障。")])],1),t._v(" "),e("el-form-item",{attrs:{label:"机器人内网IP地址"}},[e("el-input",{staticStyle:{width:"240px"},attrs:{placeholder:"请输入机器人IP地址",autocomplete:"off",maxlength:"50",disabled:""},model:{value:t.Data.config.robot_ip,callback:function(s){t.$set(t.Data.config,"robot_ip",s)},expression:"Data.config.robot_ip"}}),t._v(" "),e("div",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("系统自动获取，无需手动设置，请勿自行修改，否则导致机器人故障。")])],1),t._v(" "),e("el-form-item",{attrs:{label:"机器人MAC地址"}},[e("el-input",{staticStyle:{width:"300px"},attrs:{placeholder:"请输入机器人MAC地址",autocomplete:"off",maxlength:"50",disabled:""},model:{value:t.Data.config.robot_mac,callback:function(s){t.$set(t.Data.config,"robot_mac",s)},expression:"Data.config.robot_mac"}}),t._v(" "),e("div",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("系统自动获取，无需手动设置，请勿自行修改，否则导致机器人故障。")])],1)],1)],1),t._v(" "),e("el-card",{staticClass:"robot-box-card"},[e("div",{staticClass:"title"},[t._v("机器人主控开发板配置")]),t._v(" "),e("el-form",{ref:"form",attrs:{model:t.Data,"label-width":"150px"}},[e("el-form-item",{attrs:{label:"主控开发板串口"}},[e("el-input",{staticStyle:{width:"300px"},attrs:{placeholder:"请输入串口地址",autocomplete:"off"},model:{value:t.Data.config.robot_board.port,callback:function(s){t.$set(t.Data.config.robot_board,"port",s)},expression:"Data.config.robot_board.port"}}),t._v(" "),e("el-input",{staticStyle:{width:"120px"},attrs:{placeholder:"请输入波特率",autocomplete:"off"},model:{value:t.Data.config.robot_board.rate,callback:function(s){t.$set(t.Data.config.robot_board,"rate",s)},expression:"Data.config.robot_board.rate"}}),t._v(" "),e("el-input",{staticStyle:{width:"120px"},attrs:{placeholder:"请输入数据位",autocomplete:"off"},model:{value:t.Data.config.robot_board.bits,callback:function(s){t.$set(t.Data.config.robot_board,"bits",s)},expression:"Data.config.robot_board.bits"}})],1),t._v(" "),e("el-form-item",{attrs:{label:""}},[e("div",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("硬件设备的串口地址、波特率、数据位，请确保机器人主控开发板已经连接到机器人设备，修改后重启机器人设备生效。")])])],1)],1),t._v(" "),e("el-card",{staticClass:"robot-box-card"},[e("el-form",{ref:"form",attrs:{model:t.Data,"label-width":"150px"}},[e("el-form-item",[e("el-button",{attrs:{type:"danger",loading:t.ButtonStatus},on:{click:t.onSubmitRobot}},[t._v("保存设置")])],1)],1)],1)],1)]):t._e()])},staticRenderFns:[]};var y=e("VU/8")(j,S,!1,function(t){e("Z45g")},"data-v-fb58ab3c",null).exports,w=e("E5Az"),D=e.n(w),T=(e("CQVp"),e("UM8r"),e("aGTD"),e("U80t"),e("c+I8"),e("jQeI"),e("PsxY"),e("Ev5Y"),e("QSKu"),e("/A6h"),e("OkRY"),e("6S2o"),e("vq+x"),e("Z6qg"),e("RkhK"),e("LxZp"),e("CK2l"),e("7Xsf"),e("Yokd"),e("fo6W"),e("soCA"),e("THjC"),e("Kk9m"),e("U3HU"),e("dxBS"),e("TQy8"),{name:"App",components:{CodeMirror:D.a},data:()=>({Data:!1,dialogBoxStatus:!1,codeType:"",codeData:"",cmOption:{tabSize:4,indentUnit:4,styleActiveLine:!1,lineNumbers:!0,styleSelectedText:!1,line:!0,foldGutter:!0,gutters:["CodeMirror-linenumbers","CodeMirror-foldgutter"],highlightSelectionMatches:{showToken:/\w/,annotateScrollbar:!0},mode:"text/x-go",hintOptions:{completeSingle:!1},keyMap:"sublime",matchBrackets:!0,showCursorWhenSelecting:!0,theme:"monokai",extraKeys:{Ctrl:"autocomplete"}},fullscreenLoading:!1,Progress:0,ProgressText:"准备就绪",saveStatus:!1,runStatus:!1}),created(){this.init()},mounted(){},methods:{init(){f({url:"common/home/skill",method:"GET",data:""}).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):(this.Data=t.data.data,this.$emit("main",this.Data))})},HandleClose(t){t(),this.codeData="",this.codeType=""},editMaster(){this.fullscreenLoading=!0,this.codeType="Master",_(this.codeType).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):0===t.data.code?(this.codeData=t.data.data.code,this.dialogBoxStatus=!0,this.fullscreenLoading=!1):(this.$message.error(t.data.msg),this.fullscreenLoading=!1,this.codeType="")})},editSkill(){this.fullscreenLoading=!0,this.codeType="Skill",_(this.codeType).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):0===t.data.code?(this.codeData=t.data.data.code,this.dialogBoxStatus=!0,this.fullscreenLoading=!1):(this.$message.error(t.data.msg),this.fullscreenLoading=!1,this.codeType="")})},onSave(){var t;this.saveStatus=!0,(t={code:this.codeData,type:this.codeType},f({url:"common/home/skill/save",method:"POST",data:t})).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):0===t.data.code?(this.$message({message:"程序代码保存成功",type:"success"}),this.saveStatus=!1):(this.$message.error(t.data.msg),this.saveStatus=!1)})},onRun(){let t=this;var s;t.runStatus=!0,t.ProgressText="正在编译技能程序",t.Progress=10,(s=t.codeType,f({url:"common/home/skill/build?type="+s,method:"GET",data:""})).then(s=>{-1===s.data.code?t.$router.push({path:"/"}):0===s.data.code?(t.ProgressText="技能程序编译成功",t.Progress=30,setTimeout(function(){var s;t.ProgressText="正在重启机器人程序",t.Progress=55,(s=t.codeType,f({url:"common/home/skill/restart?type="+s,method:"GET",data:""})).then(s=>{-1===s.data.code?t.$router.push({path:"/"}):0===s.data.code?(t.ProgressText="机器人技能程序重启成功",t.Progress=100,t.runStatus=!1):(t.$message.error(s.data.msg),t.runStatus=!1,t.ProgressText="准备就绪",t.Progress=0)})},900)):(t.$message.error(s.data.msg),t.runStatus=!1,t.ProgressText="准备就绪",t.Progress=0)})},addSkill(){}},beforeDestroy(){}}),$={render:function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("el-container",[t.Data?e("div",{directives:[{name:"loading",rawName:"v-loading.fullscreen.lock",value:t.fullscreenLoading,expression:"fullscreenLoading",modifiers:{fullscreen:!0,lock:!0}}],staticClass:"main",attrs:{"element-loading-background":"rgba(0, 0, 0, .8)"}},[e("div",{staticClass:"main-box"},[e("div",{staticClass:"skill-box"},[e("el-row",{attrs:{gutter:15}},[e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd master"},[t._v("机器人主技能程序 "),e("span",[t._v("Golang")])]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger"},on:{click:t.editMaster}},[e("i",{staticClass:"iconfont icon-brush_fill"}),t._v(" 编辑程序")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("测试小技能 "),e("span",[t._v("Python")])]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger"},on:{click:t.editSkill}},[e("i",{staticClass:"iconfont icon-brush_fill"}),t._v(" 编辑技能")]),t._v(" "),e("el-button",{staticClass:"tools-button delete",attrs:{type:"danger"},on:{click:t.editSkill}},[e("i",{staticClass:"iconfont icon-trash_fill"}),t._v(" 删除技能")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content end",on:{click:t.addSkill}},[e("i",{staticClass:"iconfont icon-brush_fill"}),t._v(" 创建机器人技能\n                        ")])])],1)],1)])]):t._e(),t._v(" "),e("el-dialog",{attrs:{visible:t.dialogBoxStatus,width:"1200px","before-close":t.HandleClose,closeOnClickModal:!1,closeOnPressEscape:!1},on:{"update:visible":function(s){t.dialogBoxStatus=s}}},[e("div",{staticClass:"edit-box"},[e("div",{staticClass:"edit-tools"},[e("div",{staticClass:"item"},[e("div",{staticClass:"run-status-box"},[e("div",{staticClass:"title"},[t._v(t._s(t.ProgressText))]),t._v(" "),e("el-progress",{attrs:{percentage:t.Progress,color:"#409eff","text-inside":!0,"show-text":!1}})],1)]),t._v(" "),e("div",{staticClass:"item right"},[e("el-button",{staticClass:"el-button-tools",attrs:{loading:t.saveStatus,type:"danger"},on:{click:t.onSave}},[t.saveStatus?t._e():e("i",{staticClass:"iconfont icon-order_fill"}),t._v("保存程序")]),t._v(" "),e("el-button",{staticClass:"el-button-tools run",attrs:{loading:t.runStatus,type:"danger"},on:{click:t.onRun}},[t.runStatus?t._e():e("i",{staticClass:"iconfont icon-play_fill"}),t._v("运行程序")])],1)]),t._v(" "),e("codemirror",{ref:"codemirrorBox",attrs:{options:t.cmOption},model:{value:t.codeData,callback:function(s){t.codeData=s},expression:"codeData"}})],1)])],1)},staticRenderFns:[]};var z=e("VU/8")(T,$,!1,function(t){e("BQy5")},"data-v-44ab01dd",null).exports,P={name:"App",data:()=>({Data:!1}),created(){this.init()},mounted(){},methods:{init(){b().then(t=>{-1===t.data.code?this.$router.push({path:"/"}):(this.Data=t.data.data,this.$emit("main",this.Data))})}},beforeDestroy(){}},R={render:function(){var t=this.$createElement,s=this._self._c||t;return s("el-container",[s("div",{staticClass:"main"},[s("div",{staticClass:"main-box"},[s("el-card",{staticClass:"empty-box-card"},[this._v("没有找到相关数据")])],1)])])},staticRenderFns:[]};var B=e("VU/8")(P,R,!1,function(t){e("w2vp")},"data-v-83e0f2f6",null).exports,E={name:"App",data:()=>({Data:!1,toolsTitle:"",dialogTools:!1,toolsIndex:0,SocketStatus:!0,remote:{content:"",ButtonStatus:!1},serial:{port:"/dev/ttyACM0",rate:"115200",bits:"8",content:"",Switch:!1,ReturnContent:"",ButtonStatus:!1},tcp:{ip:"192.168.1.8",port:"40923",content:"",Switch:!1,ReturnContent:"",ButtonStatus:!1}}),inject:["MessagesEmpty"],props:{message:{type:Object,default:""}},watch:{message:{handler(){0!==Object.keys(this.message).length&&this.SocketStatus&&this.SocketCallback(this.message)},deep:!0}},created(){this.init()},mounted(){},methods:{init(){f({url:"common/home/tools",method:"GET",data:""}).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):(this.Data=t.data.data,this.$emit("main",this.Data))})},onTools(t){3===t?(this.toolsTitle="串口调试",this.dialogTools=!0,this.toolsIndex=t):2===t?(this.toolsTitle="远程消息",this.dialogTools=!0,this.toolsIndex=t):4===t?(this.toolsTitle="tcp/socket通讯",this.dialogTools=!0,this.toolsIndex=t):this.$message({message:"功能被锁，请联系管理员",type:"warning"})},HandleClose(t){t(),this.toolsIndex=0,this.serial={port:"/dev/ttyACM0",rate:"115200",bits:"8",content:"",Switch:!1,ReturnContent:"",ButtonStatus:!1},this.remote={content:"",ButtonStatus:!1}},onSubmitSerial(){var t;""===this.serial.port||""===this.serial.rate||""===this.serial.bits||""===this.serial.content?this.$message({message:"串口信息不完整，请补充完整",type:"warning"}):(this.serial.ButtonStatus=!0,(t=this.serial,f({url:"common/home/tools/serial/submit",method:"POST",data:t})).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):0===t.data.code?this.$message({message:"串口数据发送成功",type:"success"}):this.$message.error(t.data.msg),this.serial.ButtonStatus=!1}))},onSubmitRemote(){var t;""===this.remote.content?this.$message({message:"消息内容不完整，请补充完整",type:"warning"}):(this.remote.ButtonStatus=!0,(t=this.remote,f({url:"common/home/tools/remote/submit",method:"POST",data:t})).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):0===t.data.code?this.$message({message:"远程消息发送成功",type:"success"}):this.$message.error(t.data.msg),this.remote.ButtonStatus=!1}))},onSubmitTcp(){var t;""===this.tcp.ip&&""===this.tcp.content&&""===this.tcp.port?this.$message({message:"通讯信息不完整，请补充完整",type:"warning"}):(this.tcp.ButtonStatus=!0,(t=this.tcp,f({url:"common/home/tools/tcp/submit",method:"POST",data:t})).then(t=>{-1===t.data.code?this.$router.push({path:"/"}):0===t.data.code?this.$message({message:"通讯数据发送成功",type:"success"}):this.$message.error(t.data.msg),this.tcp.ButtonStatus=!1}))},SocketCallback(t){"serial_message"===t.message_type&&this.serial.Switch&&(this.serial.ReturnContent=t.serial_message.content),"tcp_message"===t.message_type&&this.tcp.Switch&&(this.tcp.ReturnContent=t.tcp_message.content)}},beforeDestroy(){this.SocketStatus=!1}},F={render:function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("el-container",[e("div",{staticClass:"main"},[e("div",{staticClass:"main-box"},[e("div",{staticClass:"tools-box"},[e("el-row",{attrs:{gutter:15}},[e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("远程控制")]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger","data-tools":"远程控制"},on:{click:function(s){return t.onTools(1)}}},[t._v("打开工具")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("远程消息")]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger","data-tools":"远程消息"},on:{click:function(s){return t.onTools(2)}}},[t._v("打开工具")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("串口调试")]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger","data-tools":"串口调试"},on:{click:function(s){return t.onTools(3)}}},[t._v("打开工具")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("tcp/socket通讯")]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger","data-tools":"tcp/socket通讯"},on:{click:function(s){return t.onTools(4)}}},[t._v("打开工具")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("陀螺仪")]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger","data-tools":"陀螺仪"},on:{click:function(s){return t.onTools(5)}}},[t._v("打开工具")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("总线舵机")]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger","data-tools":"总线舵机"},on:{click:function(s){return t.onTools(6)}}},[t._v("打开工具")])],1)])]),t._v(" "),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"grid-content"},[e("div",{staticClass:"hd"},[t._v("PWM")]),t._v(" "),e("div",{staticClass:"button-box"},[e("el-button",{staticClass:"tools-button",attrs:{type:"danger","data-tools":"PWM"},on:{click:function(s){return t.onTools(7)}}},[t._v("打开工具")])],1)])])],1)],1)])]),t._v(" "),e("el-dialog",{attrs:{title:t.toolsTitle,visible:t.dialogTools,width:"1000px","before-close":t.HandleClose,closeOnClickModal:!1,closeOnPressEscape:!1},on:{"update:visible":function(s){t.dialogTools=s}}},[2===t.toolsIndex?e("el-form",{ref:"form",attrs:{model:t.remote,"label-width":"0px"}},[e("el-form-item",{attrs:{label:""}},[e("el-input",{staticStyle:{width:"100%",resize:"none"},attrs:{type:"textarea",placeholder:"请输入要发送的消息数据",autocomplete:"off",rows:"10",resize:"none"},model:{value:t.remote.content,callback:function(s){t.$set(t.remote,"content",s)},expression:"remote.content"}}),t._v(" "),e("span",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("根据自定义的消息通讯协议，填写要发送的数据。")])],1),t._v(" "),e("el-form-item",[e("el-button",{attrs:{type:"danger",loading:t.remote.ButtonStatus},on:{click:t.onSubmitRemote}},[t._v("发送消息")])],1)],1):t._e(),t._v(" "),3===t.toolsIndex?e("el-form",{ref:"form",attrs:{model:t.serial,"label-width":"0px"}},[e("el-form-item",{attrs:{label:""}},[e("el-input",{staticStyle:{width:"300px"},attrs:{placeholder:"请输入串口地址",autocomplete:"off"},model:{value:t.serial.port,callback:function(s){t.$set(t.serial,"port",s)},expression:"serial.port"}}),t._v(" "),e("el-input",{staticStyle:{width:"120px"},attrs:{placeholder:"波特率",autocomplete:"off"},model:{value:t.serial.rate,callback:function(s){t.$set(t.serial,"rate",s)},expression:"serial.rate"}}),t._v(" "),e("el-input",{staticStyle:{width:"100px"},attrs:{placeholder:"数据位",autocomplete:"off",disabled:""},model:{value:t.serial.bits,callback:function(s){t.$set(t.serial,"bits",s)},expression:"serial.bits"}}),t._v(" "),e("span",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("硬件设备的串口地址、波特率、数据位，例如：/dev/ttyACM0。")])],1),t._v(" "),e("el-form-item",{attrs:{label:""}},[e("el-input",{staticStyle:{width:"100%",resize:"none"},attrs:{type:"textarea",placeholder:"请输入要发送的串口数据",autocomplete:"off",rows:"10",resize:"none"},model:{value:t.serial.content,callback:function(s){t.$set(t.serial,"content",s)},expression:"serial.content"}}),t._v(" "),e("span",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("根据串口协议，填写要发送的数据。")])],1),t._v(" "),e("el-form-item",[e("el-button",{attrs:{type:"danger",loading:t.serial.ButtonStatus},on:{click:t.onSubmitSerial}},[t._v("发送串口数据")])],1),t._v(" "),e("el-form-item",{attrs:{label:""}},[e("el-switch",{attrs:{"active-color":"#F56C6C","inactive-color":"#30333d"},model:{value:t.serial.Switch,callback:function(s){t.$set(t.serial,"Switch",s)},expression:"serial.Switch"}}),t._v(" "),e("span",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("是否开启串口数据监听，开启后，可以实时获取串口发送回来的数据信息。")])],1),t._v(" "),t.serial.Switch?e("el-form-item",{attrs:{label:""}},[e("el-input",{staticStyle:{width:"100%",resize:"none"},attrs:{type:"textarea",placeholder:"这里会实时显示串口发来的数据信息",autocomplete:"off",rows:"12",resize:"none"},model:{value:t.serial.ReturnContent,callback:function(s){t.$set(t.serial,"ReturnContent",s)},expression:"serial.ReturnContent"}})],1):t._e()],1):t._e(),t._v(" "),4===t.toolsIndex?e("el-form",{ref:"form",attrs:{model:t.tcp,"label-width":"0px"}},[e("el-form-item",{attrs:{label:""}},[e("el-input",{staticStyle:{width:"300px"},attrs:{placeholder:"请输入通讯IP地址",autocomplete:"off"},model:{value:t.tcp.ip,callback:function(s){t.$set(t.tcp,"ip",s)},expression:"tcp.ip"}}),t._v(" "),e("el-input",{staticStyle:{width:"160px"},attrs:{placeholder:"写入端口",autocomplete:"off"},model:{value:t.tcp.port,callback:function(s){t.$set(t.tcp,"port",s)},expression:"tcp.port"}}),t._v(" "),e("span",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("tcp/socket通讯配置，包含IP、写入端口。")])],1),t._v(" "),e("el-form-item",{attrs:{label:""}},[e("el-input",{staticStyle:{width:"100%",resize:"none"},attrs:{type:"textarea",placeholder:"请输入要发送的tcp/socket数据，为空则视为只读数据",autocomplete:"off",rows:"10",resize:"none"},model:{value:t.tcp.content,callback:function(s){t.$set(t.tcp,"content",s)},expression:"tcp.content"}}),t._v(" "),e("span",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("根据tcp/socket协议，填写要发送的数据。")])],1),t._v(" "),e("el-form-item",[e("el-button",{attrs:{type:"danger",loading:t.tcp.ButtonStatus},on:{click:t.onSubmitTcp}},[t._v("发送数据")])],1),t._v(" "),e("el-form-item",{attrs:{label:""}},[e("el-switch",{attrs:{"active-color":"#F56C6C","inactive-color":"#30333d"},model:{value:t.tcp.Switch,callback:function(s){t.$set(t.tcp,"Switch",s)},expression:"tcp.Switch"}}),t._v(" "),e("span",{staticClass:"el-input__span"},[e("i",{staticClass:"iconfont icon-feedback_fill"}),t._v("是否开启串口数据监听，开启后，可以实时获取串口发送回来的数据信息。")])],1),t._v(" "),t.tcp.Switch?e("el-form-item",{attrs:{label:""}},[e("el-input",{staticStyle:{width:"100%",resize:"none"},attrs:{type:"textarea",placeholder:"这里会实时显示tcp/socket发来的数据信息",autocomplete:"off",rows:"12",resize:"none"},model:{value:t.tcp.ReturnContent,callback:function(s){t.$set(t.tcp,"ReturnContent",s)},expression:"tcp.ReturnContent"}})],1):t._e()],1):t._e()],1)],1)},staticRenderFns:[]};var M=e("VU/8")(E,F,!1,function(t){e("FWHM")},"data-v-73a2e15a",null).exports;a.default.use(n.a);var I=new n.a({routes:[{path:"/",component:d,children:[{path:"/",name:"HomeIndex",component:k},{path:"/robot",name:"HomeRobot",component:y},{path:"/skill",name:"HomeSkill",component:z},{path:"/ability",name:"HomeAbility",component:B},{path:"/tools",name:"HomeTools",component:M}]}]}),O=e("V8mf"),U=e.n(O),H=e("wvfG"),A=e.n(H),L=e("gBx8"),q=e.n(L);e("tvR6"),e("TJvI"),e("4/hK");a.default.config.productionTip=!1,a.default.use(p.a),a.default.use(q.a),a.default.use(D.a,{options:{theme:""}}),a.default.directive("highlight",function(t){t.querySelectorAll("pre code").forEach(t=>{U.a.highlightBlock(t)})}),A.a.config.autoSetContainer=!0,a.default.use(A.a),new a.default({el:"#app",router:I,components:{App:i},template:"<App/>"})},PsxY:function(t,s){},TJvI:function(t,s){},UM8r:function(t,s){},Yokd:function(t,s){},Z45g:function(t,s){},bP0V:function(t,s){},bf6V:function(t,s){},oP5K:function(t,s,e){t.exports=e.p+"static/img/qrcode_for_gh_a2c92d370ff8_258.585ea61.jpg"},tvR6:function(t,s){},uslO:function(t,s,e){var a={"./af":"3CJN","./af.js":"3CJN","./ar":"3MVc","./ar-dz":"tkWw","./ar-dz.js":"tkWw","./ar-kw":"j8cJ","./ar-kw.js":"j8cJ","./ar-ly":"wPpW","./ar-ly.js":"wPpW","./ar-ma":"dURR","./ar-ma.js":"dURR","./ar-sa":"7OnE","./ar-sa.js":"7OnE","./ar-tn":"BEem","./ar-tn.js":"BEem","./ar.js":"3MVc","./az":"eHwN","./az.js":"eHwN","./be":"3hfc","./be.js":"3hfc","./bg":"lOED","./bg.js":"lOED","./bm":"hng5","./bm.js":"hng5","./bn":"aM0x","./bn.js":"aM0x","./bo":"w2Hs","./bo.js":"w2Hs","./br":"OSsP","./br.js":"OSsP","./bs":"aqvp","./bs.js":"aqvp","./ca":"wIgY","./ca.js":"wIgY","./cs":"ssxj","./cs.js":"ssxj","./cv":"N3vo","./cv.js":"N3vo","./cy":"ZFGz","./cy.js":"ZFGz","./da":"YBA/","./da.js":"YBA/","./de":"DOkx","./de-at":"8v14","./de-at.js":"8v14","./de-ch":"Frex","./de-ch.js":"Frex","./de.js":"DOkx","./dv":"rIuo","./dv.js":"rIuo","./el":"CFqe","./el.js":"CFqe","./en-SG":"oYA3","./en-SG.js":"oYA3","./en-au":"Sjoy","./en-au.js":"Sjoy","./en-ca":"Tqun","./en-ca.js":"Tqun","./en-gb":"hPuz","./en-gb.js":"hPuz","./en-ie":"ALEw","./en-ie.js":"ALEw","./en-il":"QZk1","./en-il.js":"QZk1","./en-nz":"dyB6","./en-nz.js":"dyB6","./eo":"Nd3h","./eo.js":"Nd3h","./es":"LT9G","./es-do":"7MHZ","./es-do.js":"7MHZ","./es-us":"INcR","./es-us.js":"INcR","./es.js":"LT9G","./et":"XlWM","./et.js":"XlWM","./eu":"sqLM","./eu.js":"sqLM","./fa":"2pmY","./fa.js":"2pmY","./fi":"nS2h","./fi.js":"nS2h","./fo":"OVPi","./fo.js":"OVPi","./fr":"tzHd","./fr-ca":"bXQP","./fr-ca.js":"bXQP","./fr-ch":"VK9h","./fr-ch.js":"VK9h","./fr.js":"tzHd","./fy":"g7KF","./fy.js":"g7KF","./ga":"U5Iz","./ga.js":"U5Iz","./gd":"nLOz","./gd.js":"nLOz","./gl":"FuaP","./gl.js":"FuaP","./gom-latn":"+27R","./gom-latn.js":"+27R","./gu":"rtsW","./gu.js":"rtsW","./he":"Nzt2","./he.js":"Nzt2","./hi":"ETHv","./hi.js":"ETHv","./hr":"V4qH","./hr.js":"V4qH","./hu":"xne+","./hu.js":"xne+","./hy-am":"GrS7","./hy-am.js":"GrS7","./id":"yRTJ","./id.js":"yRTJ","./is":"upln","./is.js":"upln","./it":"FKXc","./it-ch":"/E8D","./it-ch.js":"/E8D","./it.js":"FKXc","./ja":"ORgI","./ja.js":"ORgI","./jv":"JwiF","./jv.js":"JwiF","./ka":"RnJI","./ka.js":"RnJI","./kk":"j+vx","./kk.js":"j+vx","./km":"5j66","./km.js":"5j66","./kn":"gEQe","./kn.js":"gEQe","./ko":"eBB/","./ko.js":"eBB/","./ku":"kI9l","./ku.js":"kI9l","./ky":"6cf8","./ky.js":"6cf8","./lb":"z3hR","./lb.js":"z3hR","./lo":"nE8X","./lo.js":"nE8X","./lt":"/6P1","./lt.js":"/6P1","./lv":"jxEH","./lv.js":"jxEH","./me":"svD2","./me.js":"svD2","./mi":"gEU3","./mi.js":"gEU3","./mk":"Ab7C","./mk.js":"Ab7C","./ml":"oo1B","./ml.js":"oo1B","./mn":"CqHt","./mn.js":"CqHt","./mr":"5vPg","./mr.js":"5vPg","./ms":"ooba","./ms-my":"G++c","./ms-my.js":"G++c","./ms.js":"ooba","./mt":"oCzW","./mt.js":"oCzW","./my":"F+2e","./my.js":"F+2e","./nb":"FlzV","./nb.js":"FlzV","./ne":"/mhn","./ne.js":"/mhn","./nl":"3K28","./nl-be":"Bp2f","./nl-be.js":"Bp2f","./nl.js":"3K28","./nn":"C7av","./nn.js":"C7av","./pa-in":"pfs9","./pa-in.js":"pfs9","./pl":"7LV+","./pl.js":"7LV+","./pt":"ZoSI","./pt-br":"AoDM","./pt-br.js":"AoDM","./pt.js":"ZoSI","./ro":"wT5f","./ro.js":"wT5f","./ru":"ulq9","./ru.js":"ulq9","./sd":"fW1y","./sd.js":"fW1y","./se":"5Omq","./se.js":"5Omq","./si":"Lgqo","./si.js":"Lgqo","./sk":"OUMt","./sk.js":"OUMt","./sl":"2s1U","./sl.js":"2s1U","./sq":"V0td","./sq.js":"V0td","./sr":"f4W3","./sr-cyrl":"c1x4","./sr-cyrl.js":"c1x4","./sr.js":"f4W3","./ss":"7Q8x","./ss.js":"7Q8x","./sv":"Fpqq","./sv.js":"Fpqq","./sw":"DSXN","./sw.js":"DSXN","./ta":"+7/x","./ta.js":"+7/x","./te":"Nlnz","./te.js":"Nlnz","./tet":"gUgh","./tet.js":"gUgh","./tg":"5SNd","./tg.js":"5SNd","./th":"XzD+","./th.js":"XzD+","./tl-ph":"3LKG","./tl-ph.js":"3LKG","./tlh":"m7yE","./tlh.js":"m7yE","./tr":"k+5o","./tr.js":"k+5o","./tzl":"iNtv","./tzl.js":"iNtv","./tzm":"FRPF","./tzm-latn":"krPU","./tzm-latn.js":"krPU","./tzm.js":"FRPF","./ug-cn":"To0v","./ug-cn.js":"To0v","./uk":"ntHu","./uk.js":"ntHu","./ur":"uSe8","./ur.js":"uSe8","./uz":"XU1s","./uz-latn":"/bsm","./uz-latn.js":"/bsm","./uz.js":"XU1s","./vi":"0X8Q","./vi.js":"0X8Q","./x-pseudo":"e/KL","./x-pseudo.js":"e/KL","./yo":"YXlc","./yo.js":"YXlc","./zh-cn":"Vz2w","./zh-cn.js":"Vz2w","./zh-hk":"ZUyn","./zh-hk.js":"ZUyn","./zh-tw":"BbgG","./zh-tw.js":"BbgG"};function o(t){return e(i(t))}function i(t){var s=a[t];if(!(s+1))throw new Error("Cannot find module '"+t+"'.");return s}o.keys=function(){return Object.keys(a)},o.resolve=i,t.exports=o,o.id="uslO"},w2vp:function(t,s){}},["NHnr"]);