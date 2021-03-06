import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import AuthGuards from '../services/guards/AuthGuards'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'), beforeEnter: [AuthGuards.checkIfToken]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('../views/About.vue')
  },
  {path:'/contact',name:'Contact', component:()=>import('../views/Contact.vue')},
  {path: '/court',name: 'Court',component:()=> import('../views/Court/Court.vue')},
  {path: '/dashboard',name: 'Dashboard',component:()=> import('../views/Dashboard/Dashboard.vue'), beforeEnter: [AuthGuards.checkIfToken, AuthGuards.authGuardAdmin]},
  {path: '/login',name: 'Login',component:()=> import('../views/Login/Login.vue'), beforeEnter: [AuthGuards.checkIfToken]},
  {path: '/otp',name: 'Otp',component: ()=> import('../views/Login/OTP.vue'), beforeEnter: [AuthGuards.checkIfToken, AuthGuards.authGuardUser]},
  {path: '/profile', name:'Profile', component: () => import('../views/Profile/Profile.vue'), beforeEnter: [AuthGuards.checkIfToken, AuthGuards.authGuardUser]},
  // {path: '/register',name: 'Register',component:()=> import('../views/Login/Register.vue')},
  {path: '/reservation',name: 'Reservation',component:()=> import('../views/Reservation/Reservation.vue'), beforeEnter: [AuthGuards.checkIfToken]},
  {path: '/sport',name: 'Sport',component:()=> import('../views/Sport/Sport.vue')},
  {path: '/report-incident', name: 'Report', component: ()=> import('../views/ReportIncident.vue'), beforeEnter: [AuthGuards.checkIfToken]}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
