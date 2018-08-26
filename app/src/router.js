import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: () => import('@/views/Splash')
    },
    {
      path: '/profile',
      component: () => import('@/views/Profile')
    },
    {
      path: '/supporter',
      component: () => import('@/views/Supporter/Container'),
      children: [
        {
          path: '',
          component: () => import('@/views/Supporter/Index')
        }
      ]
    },
    {
      path: '/entrepreneur',
      component: () => import('@/views/Entrepreneur/Container'),
      children: [
        {
          path: '',
          component: () => import('@/views/Entrepreneur/Index')
        }
      ]
    }
  ]
});

export default router;
