import {createRouter, createWebHistory} from 'vue-router'
import {useAuthStore} from "@/stores/auth";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('../layouts/MainLayout.vue'),
            children: [
                {
                    path: '',
                    component: () => import('../views/HomeView.vue')
                },
            ],
        },
        {
            path: '/create-course',
            component: () => import('../layouts/MainLayout.vue'), children: [{
                path: '',
                component: () => import('../views/CreateCourseView.vue')
            }],
            meta: {
                requiredAuth: true,
            },
        },
        {
            path: '/profile',
            name: 'profile',
            component: () => import('../layouts/MainLayout.vue'),
            children: [
                {
                    path: '',
                    component: () => import('../views/ProfileView.vue')
                },
                {
                    path: '/courses',
                    name: 'courses',
                    component: () => import('../views/CoursesView.vue')
                },
                {
                    path: '/courses/:id',
                    name: 'course',
                    component: () => import('../views/CourseView.vue')
                },
            ],
            meta: {
                requiredAuth: true,
            },
        },
        {
            path: '/available-courses',
            component: () => import('../layouts/MainLayout.vue'),
            children: [
                {
                    path: '',
                    name: 'available-courses',
                    component: () => import('../views/AvailableCoursesView.vue')
                },
            ],
            meta: {
                requiredAuth: true,
            },
        },
        {
            path: '/lesson',
            name: 'lesson',
            component: () => import('../views/LessonView.vue'),
            meta: {
                requiredAuth: true,
            },
        },
        {
            path: '/registration',
            component: () => import('../layouts/MainLayout.vue'),
            children: [
                {
                    path: '',
                    name: 'registration',
                    component: () => import('../views/RegistrationView.vue')
                }
            ],
        },
        {
            path: '/login',
            component: () => import('../layouts/MainLayout.vue'),
            children: [
                {
                    path: '',
                    name: 'login',
                    component: () => import('../views/LoginView.vue')
                }
            ],
        },
        {
            path: '/activate/:id',
            component: () => import('../views/EmailActivationView.vue'),
        },
        {
            path: '/:pathMatch(.*)*',
            component: () => import('../views/NotFound.vue')
        }
    ]
})
router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();
    const requiredAuth = to.meta.requiredAuth;
    const isAuthRoutes = to.name === 'login' || to.name === 'registration';
    if (!authStore.user && requiredAuth) {
        next({name: 'login'});
    } else if (authStore.user && isAuthRoutes) {
        next({path: '/available-courses'});
    } else {
        next();
    }
});

export default router
