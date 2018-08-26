import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: () => import('@/views/Home')
    },
    {
      path: '/sign-in',
      component: () => import('@/views/SignIn')
    }
  ]
});

export default router;
