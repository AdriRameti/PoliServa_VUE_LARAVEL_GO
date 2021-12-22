import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {path: '/court',name: 'Court',component:()=> import('../views/Court/Court.vue')},
  {path: '/dashboard',name: 'Dashboard',component:()=> import('../views/Dashboard/Dashboard.vue')},
  {path: '/login',name: 'Login',component:()=> import('../views/Login/Login.vue')},
  {path: '/register',name: 'Register',component:()=> import('../views/Login/Register.vue')},
  {path: '/reservation',name: 'Reservation',component:()=> import('../views/Reservation/Reservation.vue')},
  {path: '/sport',name: 'Sport',component:()=> import('../views/Sport/Sport.vue')}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
