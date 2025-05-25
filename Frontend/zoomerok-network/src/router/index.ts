import { createRouter, createWebHistory } from 'vue-router'
import { authGuard } from './guards'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: `/news`,
    },
    {
      path: '/news',
      name: 'News',
      component: () => import('../views/NewsView.vue'),
    },
    {
      path: '/profile/:login',
      name: 'Profile',
      component: () => import('../views/ProfileView.vue'),
    },
    {
      path: '/friends',
      name: 'Friends',
      component: () => import('../views/FriendsView.vue'),
    },
  ],
})

router.beforeEach(authGuard)

export default router
