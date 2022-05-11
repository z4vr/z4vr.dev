/** @format */

import Router from 'vue-router';
import Main from '../routes/Main';

export default new Router({
  mode: 'history',

  routes: [
    {
      path: '/',
      name: 'func main()',
      component: Main,
    },
  ],
});
