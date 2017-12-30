<template>
  <div id="app">
    <router-view/>
    <vue-snotify class="panel_notifications"></vue-snotify>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
  @import "~vue-snotify/styles/material.css";

  * {
    font-family: 'Open Sans', sans-serif;
  }

  body {
    padding: 0;
    margin: 0;
  }

  #app {
    min-height: 100vh;
    min-width: 100vw;
  }

  .panel_notifications { z-index: 100000 }

</style>
