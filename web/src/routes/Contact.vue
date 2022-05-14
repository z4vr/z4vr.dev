<template>
  <div class="contact">
    <h1 class="page-header mb-5">CONTACT</h1>
    <b-container class="d-flex">
      <div class="d-flex mx-auto">
        <a
          v-for="sl in socialLinks"
          :key="sl.link"
          :href="sl.link"
          target="_blank"
          class="mx-2 social-btn"
        >
          <img :src="sl.img" height="60" width="60" />
        </a>
      </div>
    </b-container>
    <div class="built-with-banner text-center">
      <div class="hr"></div>
      <h3 class="mt-5">This page is built with</h3>
      <div class="mt-5 mx-auto">
        <img class="mx-3" height="45" src="../assets/vue-logo-450x450.png" />
        <img class="mx-3" height="45" src="../assets/js-logo-480x480.png" />
        <img class="mx-3" height="45" src="../assets/html5-logo-512x512.png" />
        <img class="mx-3" height="45" src="../assets/node-logo-266x266.png" />
      </div>
    </div>
  </div>
</template>

<script>
/** @format */

import logoDiscord from '../assets/social/discord-128x128.png';
import logoTwitter from '../assets/social/twitter-128x128.png';
import logoGithub from '../assets/social/github-128x128.png';

import { RouterEventBus } from '../js/router-eventbus';

export default {
  name: 'Contact',

  props: {},

  data() {
    return {
      socialLinks: [
        {
          img: logoDiscord,
          link: 'https://discord.com/users/974794021081841747',
        },
        {
          img: logoTwitter,
          link: 'https://twitter.com/z4vr1337',
        },
        {
          img: logoGithub,
          link: 'https://github.com/z4vr',
        },
      ],
    };
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
td {
  padding-right: 20px;
}
.contact {
  margin-top: 120px;
}
.social-btn {
  transition: all 0.25s ease-in-out;
}
.social-btn:hover {
  transform: scale(1.06);
}
.tfont {
  font-size: 20px;
}
.built-with-banner {
  margin-top: 100px;
}
.built-with-banner > .hr {
  width: 80%;
  height: 2px;
  background-color: rgba(255, 255, 255, 0.25);
  margin: 0px auto;
}
@media screen and (max-width: 690px) {
}
</style>