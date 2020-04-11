<template>
    <el-container>
        <div class="main" v-if="Data" v-loading.fullscreen.lock="fullscreenLoading" element-loading-background="rgba(0, 0, 0, .8)">
            <div class="main-box">
                <div class="skill-box">
                    <el-row :gutter="15">
                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd master">机器人主技能程序 <span>Golang</span></div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" @click="editMaster"><i class="iconfont icon-brush_fill"></i> 编辑程序</el-button>
                                </div>
                            </div>
                        </el-col>

                        <el-col :span="6">
                            <div class="grid-content">
                                <div class="hd">测试小技能 <span>Python</span></div>
                                <div class="button-box">
                                    <el-button class="tools-button" type="danger" @click="editSkill"><i class="iconfont icon-brush_fill"></i> 编辑技能</el-button>
                                    <el-button class="tools-button delete" type="danger" @click="editSkill"><i class="iconfont icon-trash_fill"></i> 删除技能</el-button>
                                </div>
                            </div>
                        </el-col>

                        <el-col :span="6">
                            <div class="grid-content end" @click="addSkill">
                                <i class="iconfont icon-brush_fill"></i> 创建机器人技能
                            </div>
                        </el-col>
                    </el-row>
                </div>
            </div>
        </div>

        <el-dialog :visible.sync="dialogBoxStatus" width="1200px" :before-close="HandleClose" :closeOnClickModal="false" :closeOnPressEscape="false">
            <div class="edit-box">
                <div class="edit-tools">
                    <div class="item">
                        <div class="run-status-box">
                            <div class="title">{{ProgressText}}</div>
                            <el-progress :percentage="Progress" color="#409eff" :text-inside="true" :show-text="false"></el-progress>
                        </div>
                    </div>
                    <div class="item right">
                        <el-button class="el-button-tools" :loading="saveStatus" type="danger" @click="onSave"><i class="iconfont icon-order_fill" v-if="!saveStatus"></i>保存程序</el-button>
                        <el-button class="el-button-tools run" :loading="runStatus" type="danger" @click="onRun"><i class="iconfont icon-play_fill" v-if="!runStatus"></i>运行程序</el-button>
                    </div>
                </div>
                <codemirror ref="codemirrorBox" v-model="codeData" :options="cmOption"></codemirror>
            </div>
        </el-dialog>
    </el-container>
</template>

<script>
    import{GetHomeSkill,GetHomeSkillEdit,GetHomeSkillSave,GetHomeSkillBuild,GetHomeSkillRestart} from "../../api/index";
    import CodeMirror from 'vue-codemirror';
    // language
    import "codemirror/mode/go/go.js";
    // theme css
    import "codemirror/theme/monokai.css";
    // require active-line.js
    import "codemirror/addon/selection/active-line.js";
    // styleSelectedText
    import "codemirror/addon/selection/mark-selection.js";
    import "codemirror/addon/search/searchcursor.js";
    // hint
    import "codemirror/addon/hint/show-hint.js";
    import "codemirror/addon/hint/show-hint.css";
    import "codemirror/addon/hint/anyword-hint.js";
    import "codemirror/addon/selection/active-line.js";
    // highlightSelectionMatches
    import "codemirror/addon/scroll/annotatescrollbar.js";
    import "codemirror/addon/search/matchesonscrollbar.js";
    import "codemirror/addon/search/searchcursor.js";
    import "codemirror/addon/search/match-highlighter.js";
    // keyMap
    import "codemirror/mode/clike/clike.js";
    import "codemirror/addon/edit/matchbrackets.js";
    import "codemirror/addon/comment/comment.js";
    import "codemirror/addon/dialog/dialog.js";
    import "codemirror/addon/dialog/dialog.css";
    import "codemirror/addon/search/searchcursor.js";
    import "codemirror/addon/search/search.js";
    import "codemirror/keymap/sublime.js";
    // foldGutter
    import "codemirror/addon/fold/foldgutter.css";
    import "codemirror/addon/fold/brace-fold.js";
    import "codemirror/addon/fold/comment-fold.js";
    import "codemirror/addon/fold/foldcode.js";
    import "codemirror/addon/fold/foldgutter.js";
    import "codemirror/addon/fold/indent-fold.js";
    import "codemirror/addon/fold/markdown-fold.js";
    import "codemirror/addon/fold/xml-fold.js";
    export default {
        name: 'App',
        components: {CodeMirror},
        data() {
            return {
                Data:false,
                dialogBoxStatus:false,
                codeType:"",
                codeData:"",
                cmOption: {
                    tabSize: 4,
                    indentUnit:4,
                    styleActiveLine: false,
                    lineNumbers: true,
                    styleSelectedText: false,
                    line: true,
                    foldGutter: true,
                    gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter"],
                    highlightSelectionMatches: { showToken: /\w/, annotateScrollbar: true },
                    mode: "text/x-go",
                    hintOptions: {
                        completeSingle: false,
                    },
                    keyMap: "sublime",
                    matchBrackets: true,
                    showCursorWhenSelecting: true,
                    theme: "monokai",
                    extraKeys: {
                        "Ctrl": "autocomplete"
                    }
                },
                fullscreenLoading: false,
                Progress: 0,
                ProgressText: "准备就绪",
                saveStatus: false,
                runStatus: false
            }
        },
        created(){
            this.init();
        },
        mounted() {

        },
        methods: {
            init(){
                GetHomeSkill().then(res=>{
                    if(res.data.code === -1){
                        this.$router.push({path: '/'});
                    }else{
                        this.Data = res.data.data;
                        this.$emit("main",this.Data);
                    }
                });
            },
            HandleClose(done){
                done();
                this.codeData = "";
                this.codeType = "";
            },
            editMaster(){
                this.fullscreenLoading = true;
                this.codeType = "Master";
                GetHomeSkillEdit(this.codeType).then(res=>{
                    if(res.data.code === -1) {
                        this.$router.push({path: '/'});
                    }else if(res.data.code === 0){
                        this.codeData = res.data.data.code;
                        this.dialogBoxStatus = true;
                        this.fullscreenLoading = false;
                    }else{
                        this.$message.error(res.data.msg);
                        this.fullscreenLoading = false;
                        this.codeType = "";
                    }
                });
            },
            editSkill(){
                this.fullscreenLoading = true;
                this.codeType = "Skill";
                GetHomeSkillEdit(this.codeType).then(res=>{
                    if(res.data.code === -1) {
                        this.$router.push({path: '/'});
                    }else if(res.data.code === 0){
                        this.codeData = res.data.data.code;
                        this.dialogBoxStatus = true;
                        this.fullscreenLoading = false;
                    }else{
                        this.$message.error(res.data.msg);
                        this.fullscreenLoading = false;
                        this.codeType = "";
                    }
                });
            },
            onSave(){
                this.saveStatus = true;
                GetHomeSkillSave({code:this.codeData,type:this.codeType}).then(res=>{
                    if(res.data.code === -1) {
                        this.$router.push({path: '/'});
                    }else if(res.data.code === 0){
                        this.$message({
                            message: '程序代码保存成功',
                            type: 'success'
                        });
                        this.saveStatus = false;
                    }else{
                        this.$message.error(res.data.msg);
                        this.saveStatus = false;
                    }
                });
            },
            onRun(){
                let that = this;
                that.runStatus = true;
                that.ProgressText = "正在编译技能程序";
                that.Progress = 10;
                GetHomeSkillBuild(that.codeType).then(res=>{
                    if(res.data.code === -1) {
                        that.$router.push({path: '/'});
                    }else if(res.data.code === 0){
                        that.ProgressText = "技能程序编译成功";
                        that.Progress = 30;
                        setTimeout(function(){
                            that.ProgressText = "正在重启机器人程序";
                            that.Progress = 55;
                            GetHomeSkillRestart(that.codeType).then(res=>{
                                if(res.data.code === -1) {
                                    that.$router.push({path: '/'});
                                }else if(res.data.code === 0){
                                    that.ProgressText = "机器人技能程序重启成功";
                                    that.Progress = 100;
                                    setTimeout(function(){
                                        that.ProgressText = "准备就绪";
                                        that.Progress = 0;
                                    },1500);
                                    that.runStatus = false;
                                }else{
                                    that.$message.error(res.data.msg);
                                    that.runStatus = false;
                                    that.ProgressText = "准备就绪";
                                    that.Progress = 0;
                                }
                            });
                        },900);
                    }else{
                        that.$message.error(res.data.msg);
                        that.runStatus = false;
                        that.ProgressText = "准备就绪";
                        that.Progress = 0;
                    }
                });
            },
            addSkill(){

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
            margin: 0 auto;
            .skill-box{
                width: 100%;
                min-height: 800px;
                .grid-content{
                    width: 100%;
                    line-height: 35px;
                    background-color: #333843;
                    border-radius: 6px;
                    border: 0;
                    box-shadow: 0 2px 2px 0 rgba(0,0,0,.1);
                    margin-bottom: 30px;
                    padding-bottom: 20px;
                    .hd{
                        width: 100%;
                        height: 100px;
                        background-color: #E6A23C;
                        border-radius: 6px 6px 0 0;
                        text-align: center;
                        line-height: 100px;
                        margin-bottom: 30px;
                        position: relative;
                        color: #ffffff;
                        font-weight: 500;
                        &.master{
                            color: #ffffff;
                            background-color: #409EFF;
                        }
                        span{
                            height: 24px;
                            line-height: 24px;
                            font-size: 9px;
                            position: absolute;
                            top: 10px;
                            left: 10px;
                            background-color: #ffffff;
                            border-radius: 12px;
                            padding: 0 10px;
                            color: #606266;
                            font-weight: 400;
                        }
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
                                background-color: #409EFF;
                            }
                            &.delete{
                                &:hover{
                                    color: #ffffff;
                                    background-color: #F56C6C;
                                }
                            }
                        }
                    }
                    &.end{
                        height: 185px;
                        text-align: center;
                        line-height: 185px;
                        box-shadow: 0 0 0 0 rgba(0,0,0,.1);
                        cursor: pointer;
                        &:hover{
                            color: #ffffff;
                        }
                    }
                }
            }
        }
    }
    .edit-box{
        .edit-tools{
            width: 100%;
            height: 55px;
            background-color: #000000;
            border-radius: 4px 4px 0 0;
            margin-bottom: 1px;
            padding: 10px;
            font-size: 0;
            .item{
                width: 50%;
                height: 35px;
                display: inline-block;
                vertical-align: top;
                font-size: 12px;
                .run-status-box{
                    width: 500px;
                    height: 35px;
                    border-radius: 4px;
                    background-color: #0a0a0a;
                    .title{
                        width: 100%;
                        height: 29px;
                        line-height: 29px;
                        text-align: center;
                    }
                }
                &.right{
                    text-align: right;
                    .el-button-tools{
                        line-height: 35px;
                        padding: 0 15px;
                        height: 35px;
                        background-color: #409EFF;
                        border: 0;
                        .iconfont{
                            margin-right: 5px;
                        }
                        &.run{
                            background-color: #F56C6C;
                        }
                    }
                    .el-switch{
                        margin-right: 10px;
                    }
                    span{
                        height: 35px;
                        line-height: 37px;
                        font-size: 12px;
                        margin-right: 10px;
                        display: inline-block;
                    }
                }
            }
        }
    }
</style>