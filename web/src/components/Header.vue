<template>
  <div>
    <!-- DESKTOP HEADER -->
    <b-container fluid class="header d-flex">
      <img
        src="../assets/logo-1000x1000.png"
        width="40"
        height="40"
        class="header-img mr-3 my-auto"
      />
      <h4 class="header-title mr-3 my-auto">z4vr.dev</h4>
      <button
        v-for="btn in btns"
        class="header-btn"
        :class="{ 'active': btn.active }"
        :key="btn.title"
        @click="btnClick(btn)"
      >{{ btn.title }}</button>
    </b-container>

    <!-- BURGER HEADER -->
    <b-container fluid class="header header-burger d-none">
      <img
        src="../assets/logo-1000x1000.png"
        width="30"
        height="30"
        class="header-img mr-4 my-auto"
      />
      <BurgerButton @click="burgerBtnClick" ref="burgerBtn" />
      <b-container fluid :class="burgerClasses" class="burger-buttons">
        <button
          v-for="btn in btns"
          class="header-btn burger-btn"
          :class="{ 'active': btn.active }"
          :key="btn.title"
          @click="btnClick(btn)"
        >{{ btn.title }}</button>
      </b-container>
    </b-container>

    <div class="burger-closer" :class="{ 'd-none': !burgerActive }" @click="burgerClose()"></div>
  </div>
</template>

<script>
/** @format */

import BurgerButton from '../components/BurgerButton';
import { RouterEventBus } from '../js/router-eventbus';
import { setTimeout } from 'timers-browserify';

export default {
  name: 'Header',

  props: {
    msg: String,
  },

  components: {
    BurgerButton,
  },

  data() {
    return {
      burgerActive: false,
      setBurgerActive: false,
      burgerClasses: {
        'burger-active': false,
        'd-none': true,
      },
      btns: [
        {
          name: 'HOME',
          title: 'HOME',
          window_title: 'z4vr.dev | home',
          route: '/',
          active: true,
        },
        {
          name: 'CONTACT',
          title: 'CONTACT',
          window_title: 'z4vr.dev | contact',
          route: '/contact',
          active: true,
        },
      ],
    };
  },

  methods: {
    btnClick(btn) {
      this.$router.push(btn.route);
      window.document.title = btn.window_title;
    },

    burgerBtnClick(active) {
      this.burgerActive = active;
      if (active) {
        this.burgerClasses['d-none'] = false;
        setTimeout(() => {
          this.burgerClasses['burger-active'] = true;
        }, 25);
      } else {
        this.burgerClasses['burger-active'] = false;
        setTimeout(() => {
          this.burgerClasses['d-none'] = true;
        }, 250);
      }
    },

    burgerClose() {
      this.$refs.burgerBtn.setActive(false);
    },
  },

  created() {
    RouterEventBus.$on('mounted', (name) => {
      this.btns.forEach((b) => {
        if (b.name === name) {
          b.active = true;
        } else {
          b.active = false;
        }
      });
    });
  },
};
</script>

<style scoped>
/** @format */

.header {
  position: fixed;
  height: 60px;
  top: 0px;
  z-index: 5;
  color: white !important;
  background-color: var(--c-bg-header-o);
  font-family: 'Fira Sans', sans-serif;
}

.header-btn {
  position: relative;
  margin: 0;
  padding: 0 30px;
  height: 100%;
  background-color: #00000000;
  z-index: 1;
  font-family: 'Fira Sans', sans-serif;
  font-weight: 200;
  color: white;
  transition: all 0.25s ease;
}

.header-btn::before {
  content: '';
  position: absolute;
  z-index: -1;
  left: 0;
  bottom: 0;
  height: 0%;
  width: 100%;
  background-color: var(--c-highlight);
  transition: all 0.25s ease;
}

.header-btn:hover::before {
  height: 100%;
}

.active::before {
  height: 100%;
}

.burger-btn {
  width: 100%;
  margin: 0px;
  padding: 10px 0px;
  background-color: var(--c-bg-header-o);
}

.burger-buttons {
  position: relative;
  top: -50px;
  opacity: 0;

  transition: all 0.25s ease-in-out;
}

.burger-active {
  top: 12px;
  opacity: 1;
}

.burger-closer {
  position: fixed;
  width: 100%;
  height: 100%;
  z-index: 4;
}

@media screen and (max-width: 850px) {
  .header {
    display: none !important;
  }

  .header-burger {
    display: block !important;
  }

  /* .header-btn {
      width: 100%;
      margin: 0;
      padding: 5px 0px;
    }

    .header-img {
      display: none;
    } */
}
</style>
