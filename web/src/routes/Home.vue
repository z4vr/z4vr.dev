<template>
  <div>
    <div class="main-header">
      <b-container fluid class="header-cont">
        <div class="logo mx-auto text-center">
          <img width="300" height="300" src="../assets/logo-1000x1000.png" />
          <h1 class="font-montserrat">z4vr.dev</h1>
          <h2 class="font-montserrat">welcome to my page</h2>
          <button class="down" @click="scrollTo('#scroll-about')">
            <div style="top: -20px;">
              <img width="80" src="../assets/down-vec.svg" />
            </div>
            <div style="top: -60px;">
              <img width="80" src="../assets/down-vec.svg" />
            </div>
          </button>
        </div>
      </b-container>
    </div>
    <a id="scroll-about"></a>
    <b-container id="about" class="content mt-6">
      <h1 class="page-header">about me</h1>
      <p class="mt-5">
        hi, my name is michael, i am also known as z4vr.
        <br>
        i am a {{ calculateAge(new Date('2001-06-28')) }} years old computer science student at the university of technology in vienna.
        most of the time i spend my time coding, but i also love to play games, both casually and competitively.

        currently i am working on a project called <a href="https://github.com/z4vr/subayai">subayai</a>, a discord bot that offers many different functions. i also worked on 
      </p>
    </b-container>
    <Footer />
  </div>
</template>

<script>
/** @format */

import { RouterEventBus } from '../js/router-eventbus';

export default {
  name: 'Home',

  props: {
    msg: String,
  },

  components: {},

  methods: {
    scrollTo(el) {
      var elem = this.$el.querySelector(el);
      if (elem) elem.scrollIntoView();
    },

    calculateAge(birthday) {
      var ageDifMs = Date.now() - birthday.getTime();
      var ageDate = new Date(ageDifMs);
      return Math.abs(ageDate.getUTCFullYear() - 1970);
    },
  },

  mounted() {
    RouterEventBus.$emit('mounted', this.$options.name);
  },
};
</script>

<style scoped>
/** @format */

a {
  text-decoration: underline;
}

.main-header {
  width: 100%;
  height: 100vh;
  background-image: url('../assets/main-bg.jpg');
  background-repeat: no-repeat;
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;

  transition: transform 0.25s ease-in-out;
}

.header-cont {
  padding: 20vh 0px;
}

.logo {
  opacity: 0.9;
}

.logo > h1 {
  font-weight: 300;
  font-size: 80px;
}

.logo > h2 {
  font-weight: 200;
  font-size: 27px;
}

.down {
  position: absolute;
  left: calc(50% - 50px);
  top: calc(100% - 150px);
  margin-top: auto;
  color: white !important;
}

.down > div {
  position: relative;
}

.mt-6 {
  margin-top: 100px;
}

@media screen and (max-height: 750px), screen and (max-width: 690px) {
  .logo > img {
    width: 150px;
    height: 150px;
  }

  .logo > h1 {
    font-weight: 300;
    font-size: 55px;
  }

  .logo > h2 {
    font-weight: 200;
    font-size: 19px;
  }
}
</style>
