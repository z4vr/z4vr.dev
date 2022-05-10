<template>
    <div>
        <div class="main-header">
        <b-container fluid class="header-cont">
        <div class="logo mx-auto text-center">
          <img width="300" height="300" src="../assets/logo-1000x1000.png" />
          <h1 class="font-montserrat mt-3 mb-0">z4vr</h1>
          <h2 class="font-montserrat">development</h2>
          <button class="down" @click="scrollTo('#scroll-about')">
            <p>Read more</p>
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
      <h1 class="page-header">About me</h1>
      <p class="mt-5">
        <!-- TODO write about my self -->
      </p>
    </b-container>
    </div>
</template>