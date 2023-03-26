<template>
  <div id="app">
    <router-view />
    <Footer/>
  </div>
</template>

<script>
import Footer from '../src/components/Footer.vue'
export default {
  name: 'App',
  components: {
    Footer
  }
}
</script>

<style>
body{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  background-color: #4b4b4b;
  background-image: url(../public/background.jpg);
  background-size: cover;
  background-repeat: no-repeat;
}

#id{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

h1, h2, h3, h4, h5{
  margin: 0;
  padding: 0;
}

:root {
    --swiper-theme-color: #949c9c !important;
}
</style>
