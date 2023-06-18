<!-- eslint-disable -->
<template>
    <Layout/>
      <div class="container">
         <div class="slider-info">
            <div class="slider">
                <MainSlider/>
            </div>
          </div>
          <div class="product-block">
            <Product 
            v-for="product in products" 
            :key="product.id"
            :product_data="product"
            />
          </div>
      </div>
   </template>
   
<script>
import Layout from '../Layouts/Layout.vue'
import MainSlider from './sliders/MainSlider.vue'
import Product from './Product.vue'
export default {
     name: 'HomeV',
     components: {
    Layout,
    MainSlider,
    Product,
     },
    data (){
      return{
         products: [
      ]
      } 
    },
    mounted (){
      this.$axios.get('http://localhost:8080/carpart').then(response => (this.products = response.data))
     },
}
</script>
<style>
    img{     
      width: 100%;
      height: 100%;
    }

   .container{
      margin: 0 auto;
      margin-top: 5%;
      width: 85%;
   }

   .slider{
      margin-top: 2%;
      height: 200px;
   }

   .slider-text{
      text-align: center;
   }

   .product-block{
      margin-top: 5%;
      display: flex;
      justify-content: space-between;
      flex-direction: row;
      flex-wrap: wrap;
      text-align: center;
   }

   @media screen and (max-width: 600px) {
  .product-block{
   margin-top: 0%;
  }
}

</style>