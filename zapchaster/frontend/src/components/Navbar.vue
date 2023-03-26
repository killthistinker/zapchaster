<!-- eslint-disable -->
<template>
    <nav>
      <ul>
        <li v-for="route in routes" :key="route.path">
          <router-link class="navbar-text" :to="route.path">{{ route.name }}</router-link>
        </li>
      </ul>
    </nav>
  </template>
  
  <script>
  export default {
    data() {
      return {
        routes: [
          {
            path: '/',
            name: 'Главная',
          },
          {
            path: '/about',
            name: 'О нас',
          },
          {
            path: '/contacts',
            name: 'Контакты',
          },
          {
            path: '/delivery',
            name: 'Доставка',
          },
        ],
      };
    },
  };
</script> 
<style>
    nav{
        padding: 0;
        margin: 0;
    }
    ul{
        padding: 0;
        margin: 0;
        display: flex;
        flex-direction: row;
    }
    li{
        padding: 0 15px;
        list-style-type: none;
        text-decoration: none;
        font-family: 'Unbounded';
        font-weight: 300;
        color: aliceblue;
    }

    .navbar-text{
      text-decoration: none;
      color: aliceblue;
    }
</style>