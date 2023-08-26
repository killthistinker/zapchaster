<template>
  <div id="app">
    <router-view />
    <Footer />
  </div>
</template>

<script>
import Footer from '../src/components/Footer.vue'
export default {
  name: 'App',
  components: {
    Footer
  },
  mounted() {
    // Изменение названия сайта
    document.title = 'Zapchastik.kz';
  }
}
</script>

<style>
body {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url(../public/background.jpg);
  height: 100%;
  background-size: cover;
  background-repeat: no-repeat;

}

#id {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

h1,
h2,
h3,
h4,
h5 {
  margin: 0;
  padding: 0;
}

:root {
  --swiper-theme-color: #949c9c !important;
}
</style>
