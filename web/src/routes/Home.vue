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
        Hi! I'm Michael! 👋
      <br>
        I'm a {{ calculateAge(new Date('2001-06-28')) }} years old programmer and currently a student at the Technische Universität Wien.
        I love playing with operating systems, creating useless applications and trying different methods to mess up things. Most of the time I code in Golang, Python or when it's needed Java.
      <br>
      <br>
        Currently I am working on <a href="https://github.com/z4vr/subayai">Subayai</a>, a discord bot that offers leveling, automanaging of roles and many more features. In addition I am working
        on this website! You can find the source code for it and all related things <a href="https://github.com/z4vr/z4vr.dev">here</a>.
      <br>
      <br>
        Since I am hosting most of my stuff myself, I have a lot of experience with Linux in general and currently I am getting into Docker and other containerization technologies.
        To expand my knowledge I am currently learning about web development (as you can see it's going well, or else you could not read that website).
      </p>
    </b-container>
    <b-container id="gear" class="content mt-6">
      
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
