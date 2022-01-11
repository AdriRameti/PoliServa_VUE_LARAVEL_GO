import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import AuthGuards from '../services/guards/AuthGuards'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue')
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
  {path: '/dashboard',name: 'Dashboard',component:()=> import('../views/Dashboard/Dashboard.vue'), beforeEnter: AuthGuards.authGuardAdmin},
  {path: '/login',name: 'Login',component:()=> import('../views/Login/Login.vue')},
  {path: '/otp',name: 'Otp',component: ()=> import('../views/Login/OTP.vue'), beforeEnter: AuthGuards.authGuardUser},
  {path: '/profile', name:'Profile', component: () => import('../views/Profile/Profile.vue'), beforeEnter: AuthGuards.authGuardUser},
  // {path: '/register',name: 'Register',component:()=> import('../views/Login/Register.vue')},
  {path: '/reservation',name: 'Reservation',component:()=> import('../views/Reservation/Reservation.vue')},
  {path: '/sport',name: 'Sport',component:()=> import('../views/Sport/Sport.vue')}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
