/** @format */

import Router from 'vue-router';
import Main from '../routes/Main';

export default new Router({
  mode: 'history',

  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main,
    },
  ],
});