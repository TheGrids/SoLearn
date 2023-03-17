import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../layouts/MainLayout.vue'),
      children: [{
        path: '',
        component: () => import('../views/HomeView.vue')
      }]
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('../layouts/MainLayout.vue'),
      children: [{
        path: '',
        component: () => import('../views/ProfileView.vue')
      }]
    },
    {
      path: '/registration',
      name: 'registration',
      component: () => import('../layouts/MainLayout.vue'),
      children: [{
        path: '',
        component: () => import('../views/RegistrationView.vue')
      }]
    },
  ]
})

export default router
