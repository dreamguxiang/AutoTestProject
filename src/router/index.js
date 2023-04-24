import { createRouter, createWebHashHistory } from 'vue-router'
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {path: '/', component:() => import('../views/LoginView.vue')},
    {path: '/home', component:() => import('../views/HomeView.vue')},
  ]
})

export default router
