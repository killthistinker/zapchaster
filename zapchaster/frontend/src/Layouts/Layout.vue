<template>
 <header>
    <div class="header">
        <div class="header-text">
            <router-link class="company-text" to="/">Zapchastik.kz</router-link>
            <h2 class="sub-title">Запчасти из Казахстана по выгодным ценам!</h2>
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
    .header-text{
        font-size: 38px;
        font-family: 'Unbounded';
        font-weight: 300;
    }

    .company-text{
      text-decoration: none;
      color: aliceblue;
    }

    .sub-title{
      font-size: 16px;
      color: aliceblue;
      font-weight: 200;
    }
    
@media screen and (max-width: 600px) {
  .header-text{
    width: 25%;
    font-size: 16px;
  }

  li{
    padding: 0 5px
  }

  .sub-title{
    font-size: 12px;
  }

  .navbar {
    font-size: 10px;
  }
}

@media screen and (max-width: 400px) {
  .header-text{
    font-size: 12px;
  }

  .sub-title{
    font-size: 8px;
  }

  .navbar {
    font-size: 10px;
  }
}
</style>