<template>
  <div>
    <div class="main-header">
      <b-container fluid class="header-cont">
        <div class="logo mx-auto text-center">
          <img width="300" height="300" src="../assets/logo-1000x1000.png"/>
          <h1 class="mt-6">Hi! I am Michael 👋</h1>
          <h2>Scroll down to learn more about me!</h2>
          <button class="to-about" @click="scrollTo('#scroll-about')">
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
      <h1 class="chapter-header">ABOUT ME</h1>
      <p class="mt-5 font-firasans">
      <br>
        I'm a {{ calculateAge(new Date('2001-06-28')) }} years old programmer and currently a student at the Technische Universität Wien.
        I started out writing simple scripts in Python, then moved to building bigger projects in Python still, like a Discord Music Bot and plugins for the <a href="https://github.com/Cog-Creators/Red-DiscordBot">Red Discord Bot</a>.
        Then I moved onto Go and Java, with university forcing Java down my throat (I still didn't get over it), learning a lot about CI/CD and now Docker. And now I am here, writing a Vue.js website and a Go backend so that you can
        enjoy this website.
      <br>
      <br>
        Currently I am working on <a href="https://github.com/z4vr/subayai">Subayai</a>, a discord bot that offers leveling, automanaging of roles and many more features. In addition I am working
        on this website! You can find the source code for it and all related things <a href="https://github.com/z4vr/z4vr.dev">here</a>. With this website I am also dipping my toes into Docker and
        HTTP server stuff in Golang. For serving this website I am using <a href="https://github.com/z4vr/shitsuji">shitsuji</a>, a fasthttp wrapper written in Go. You can find my other projects <a href="#scrool-projects">here</a>.
      <br>
      <br>
        Since you are reading this, I hope you enjoy my website! I am always looking for new problems to tackle (or create new ones to tackle). If you have any questions, feel free to contact me via Mail, Discord or Twitter,
        you can find all my contact information <a href="https://z4vr.dev/contact">here</a>.
      </p>
    </b-container>
    <b-container id="projects" class="content mt-6">
      <h1 class="chapter-header">PROJECTS</h1>
    </b-container>
    <b-container id="gear" class="content mt-6">
      <h1 class="chapter-header mb-5">MY GEAR</h1>
      <b-row v-for="(row, i) in rows" :key="`row-${i}`" class="data-row">
        <b-col v-for="col in row" :key="col.title">
          <h2 class="mb-3">{{ col.title }}</h2>
          <ul>
            <li v-for="(item, j) in col.items" :key="`item-${j}`">
              <a :href="item.link" target="_blank">{{ item.text }}</a>
            </li>
          </ul>
        </b-col>
      </b-row>
    </b-container>
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

  created() {
    this.rows = [
      [
        {
          title: 'PC Hardware',
          items: [
            {
              text: 'Ryzen 7 5800X / CPU',
              link: 'https://de.pcpartpicker.com/product/qtvqqs/amd-ryzen-7-5800x-38-ghz-8-core-processor-100-100000063wof',
            },
            {
              text: 'Gigabyte B550 AORUS ELITE AX V2 / Motherboard',
              link: 'https://de.pcpartpicker.com/product/Vb3mP6/gigabyte-b550-aorus-elite-ax-v2-atx-am4-motherboard-b550aorus-elite-axv2',
            },
            {
              text: 'G.Skill Trident Z Neo 32 GB @ 3800MT/s CL14 / RAM',
              link: 'https://de.pcpartpicker.com/product/w3FKHx/gskill-trident-z-neo-32-gb-2-x-16-gb-ddr4-3600-cl16-memory-f4-3600c16d-32gtznc',
            },
            {
              text: 'Noctua NH-D15 chromax.black / CPU Cooler',
              link: 'https://de.pcpartpicker.com/product/84MTwP/noctua-nh-d15-chromaxblack-8252-cfm-cpu-cooler-nh-d15-chromaxblack',
            },
            {
              text: 'EVGA GeForce RTX 2080S XC Ultra / GPU',
              link: 'https://de.pcpartpicker.com/product/n2RgXL/evga-geforce-rtx-2080-super-8-gb-xc-ultra-gaming-video-card-08g-p4-3183-kr',
            },
            {
              text: 'Samsung 980 Pro 1 TB / SSD',
              link: 'https://de.pcpartpicker.com/product/DDWBD3/samsung-980-pro-1-tb-m2-2280-nvme-solid-state-drive-mz-v8p1t0bam',
            },
            {
              text: 'Corsair RM (2021) 850 W / PSU',
              link: 'https://de.pcpartpicker.com/product/FPYmP6/corsair-rm-2021-850-w-80-gold-certified-fully-modular-atx-power-supply-cp-9020235-na'
            },
            {
              text: 'Corsair  4000D Airflow / Case',
              link: 'https://de.pcpartpicker.com/product/bCYQzy/corsair-4000d-airflow-atx-mid-tower-case-cc-9011200-ww',
            },
          ]
        },
        {
          title: 'Software',
          items: [
            {
              text: 'VS Code',
              link: 'https://code.visualstudio.com/',
            },
            {
              text: 'Goland',
              link: 'https://www.jetbrains.com/go/',
            },
            {
              text: 'IntelliJ',
              link: 'https://www.jetbrains.com/idea/',
            },
            {
              text: 'PyCharm',
              link: 'https://www.jetbrains.com/pycharm/',
            },
            {
              text: 'Windows 10 (with WSL2)',
              link: 'https://www.microsoft.com/en-us/windows/',
            },
            {
              text: 'Arch Linux',
              link: 'https://www.archlinux.org/',
            },
            {
              text: 'Docker',
              link: 'https://www.docker.com/',
            },
          ],
        },
        {
          title: 'Peripherals & Accessories',
          items: [
            {
              text: 'Logitech G Pro X Superlight / Mouse',
              link: 'https://www.logitechg.com/de-at/products/gaming-mice/pro-x-superlight-wireless-mouse.html',
            },
            {
              text: 'Custom Keyboard - Q2 - Gat Oil Kings - Traitorskeys KABUKI / Keyboard',
              link: 'https://www.keychron.com/products/keychron-q2-qmk-custom-mechanical-keyboard?variant=39610691518553'
            },
            {
              text: 'MSI Optix MAG274QRFDE-QD / Monitor',
              link: 'https://de.pcpartpicker.com/product/bWWBD3/msi-optix-mag274qrf-qd-270-2560x1440-165-hz-monitor-optix-mag274qrf-qd',
            },
            {
              text: 'Corsair MM300 / Mousepad',
              link: 'https://www.corsair.com/de/de/Kategorien/Produkte/Gaming-Mauspads/Mauspads-mit-Gewebeoberfl%C3%A4che/MM300-Gaming-Mauspad-mit-nicht-fransender-Stoffoberfl%C3%A4che-%E2%80%93-Extrabreit/p/CH-9000108-WW',
            },
          ]
        },
      ],
    ]
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
  font-weight: 200;
  font-size: 80px;
}

.logo > h2 {
  font-weight: 100;
  font-size: 27px;
}

.to-about {
  position: absolute;
  left: calc(50% - 50px);
  top: calc(100% - 150px);
  margin-top: auto;
  color: white !important;
}

.to-about > div {
  position: relative;
}

.mt-6 {
  margin-top: 150px;
}

.chapter-header {
  text-align: center;
  font-weight: 100;
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
