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
import { library } from '@fortawesome/fontawesome-svg-core'
import { faHatWizard,faCalendar } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(faCalendar);
createApp(App).use(router).use(store).use(Toast).use(VueChartkick).component('font-awesome-icon', FontAwesomeIcon).mount('#app')


