<template>
  <div id="app">
    <Header />
    <div class="router-cont">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
/** @format */
import './css/consts.css';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import Header from './components/Header';
import Router from './js/router';
export default {
  name: 'app',
  router: Router,
  components: {
    Header,
  },
};
</script>
