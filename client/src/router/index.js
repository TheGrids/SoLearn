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
      },
      {
        path: '/profile',
        name: 'profile',
        component: () => import('../views/ProfileView.vue')
      },
      {
        path: '/registration',
        name: 'registration',
        component: () => import('../views/RegistrationView.vue')
      },
      {
        path: '/login',
        name: 'login',
        component: () => import('../views/LoginView.vue')
      },
      {
        path: '/courses',
        name: 'courses',
        component: () => import('../views/CoursesView.vue')
      }]
    }
  ]
})

export default router
