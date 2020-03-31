import Vue from 'vue';
import App from './App';
import router from './router/index';
import ElementUI from 'element-ui';
import Highlight from 'highlight.js';
import VueClipboard from 'vue-clipboard2';
import Vdd from 'vddl';
import VueCodemirror from 'vue-codemirror';

import 'element-ui/lib/theme-chalk/index.css';
import 'highlight.js/styles/monokai-sublime.css';
import 'codemirror/lib/codemirror.css';

Vue.config.productionTip = false;

Vue.use(ElementUI);

Vue.use(Vdd);

Vue.use(VueCodemirror,{
    options: { theme: ''}
});

Vue.directive('highlight',function (el) {
    let highlight = el.querySelectorAll('pre code');
    highlight.forEach((block)=>{
        Highlight.highlightBlock(block);
    });
});

VueClipboard.config.autoSetContainer = true;

Vue.use(VueClipboard);

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: {
        App
    },
    template: '<App/>'
})
