<template>
 <header>
    <div class="header">
        <div class="empty-div"></div>
        <div class="header-text">
            <router-link class="company-text" to="/">Authostore</router-link>
        </div>
        <div class="navbar">
            <Navbar/>
        </div>
    </div>
 </header>
</template>
<script>
import Navbar from '../components/Navbar.vue'
export default {
    name: 'MainLayout',
    components: {
        Navbar,
    },
}
</script>
<style>
    .header{
        margin: 30px 0;
        padding: 0;
        display: flex;
        width: 100%;
        justify-content: space-around;
        align-items: center;
    }
    .company-text{
        font-size: 38px;
        font-family: 'Unbounded';
        font-weight: 300;
        color: aliceblue;
        text-decoration: none;
    }
    
@media screen and (max-width: 600px) {
  .company-text{
    font-size: 20px;
  }

  .navbar {
    font-size: 7px;
  }
}

@media screen and (max-width: 350px) {
  .company-text{
    font-size: 12px;
  }

  .navbar {
    font-size: 6px;
  }
}
</style>