import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 使用axios
import Axios from 'axios'
import VueAxios from "vue-axios"

// 高亮markdown里的语法
import 'highlight.js/styles/agate.css'

// 使用antd
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
import '@/assets/common.css'

Axios.defaults.baseURL = '/api'

import mixin from "@/mixin";

createApp(App).use(store).use(router).use(VueAxios, Axios).use(Antd).mixin(mixin).mount('#app')
