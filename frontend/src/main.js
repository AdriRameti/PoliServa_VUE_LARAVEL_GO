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
import i18n from './locales/i18n'

createApp(App).use(router).use(store).use(Toast).use(VueChartkick).use(i18n).mount('#app')


