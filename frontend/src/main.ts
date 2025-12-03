import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// 样式
import './styles/main.css'

// 指令
import { permissionDirective } from './directives/permission'

// 创建应用
const app = createApp(App)

// Pinia
const pinia = createPinia()
app.use(pinia)

// 路由
app.use(router)

// 注册指令
app.directive('permission', permissionDirective)

// 挂载
app.mount('#app')
