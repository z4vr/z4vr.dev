/** @format */

import Router from 'vue-router';

export default new Router({
  mode: 'history',

  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../routes/Home'),
    },
    {
      path: '/contact',
      name: 'contact',
      component: () => import('../routes/Contact'),
    },
    {
      path: '/projects',
      name: 'projects',
      component: () => import('../routes/Projects'),
    },
  ],
});
