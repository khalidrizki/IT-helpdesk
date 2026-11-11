import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import DashboardView from '@/views/DashboardView.vue'
import TicketsView from '@/views/TicketsView.vue'
import CreateTicketView from '@/views/CreateTicketView.vue'
import MainLayout from '@/layouts/MainLayout.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import AccountsView from '@/views/AccountsView.vue'
import ForgotPasswordView from '@/views/ForgotPasswordView.vue'
import NotFoundView from '@/views/NotFoundView.vue'
import ForbiddenView from '@/views/ForbiddenView.vue'
import { useAuth } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: MainLayout,
      redirect: '/dashboard',
      children: [
        {
          path: '/dashboard',
          name: 'dashboard',
          component: DashboardView,
          meta: {
            requiresAuth: true,
          },
        },
        {
          path: 'tickets',
          name: 'tickets',
          component: TicketsView,
          meta: {
            requiresAuth: true,
          },
        },
        {
          path: 'create',
          name: 'create',
          component: CreateTicketView,
          meta: {
            requiresAuth: true,
          },
        },
        {
          path: 'accounts',
          name: 'accounts',
          component: AccountsView,
          meta: {
            requiresAuth: true,
          },
        },
      ],
    },
    {
      path: '/login',
      component: LoginView,
    },
    {
      path: '/register',
      component: RegisterView,
    },
    {
      path: '/forgot-password',
      component: ForgotPasswordView,
    },
    {
      path: '/forbidden',
      name: 'forbidden',
      component: ForbiddenView,
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: NotFoundView,
    },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue'),
    // },
  ],
})

router.beforeEach(async (to) => {
  const auth = useAuth()
  if (to.meta.requiresAuth) {
    const isAuthenticated = await auth.fetchUser()
    if (!isAuthenticated) {
      return '/login'
    }
  }
  return true
})

export default router
