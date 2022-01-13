import { createApp } from 'vue'
import Toast from "vue-toastification"
import "vue-toastification/dist/index.css"
import App from './App.vue'
import router from './router'
import store from './store'
import 'bootstrap'
import './assets/app.scss'
import VueChartkick from 'vue-chartkick'
import 'chartkick/chart.js'
createApp(App).use(router).use(store).use(Toast).use(VueChartkick).mount('#app')


