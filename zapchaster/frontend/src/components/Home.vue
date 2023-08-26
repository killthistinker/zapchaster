<template>
   <Layout />
   <div class="container">
      <div class="slider-info">
         <div class="slider">
            <MainSlider />
         </div>
      </div>
      <div class="search">
         <div class="input-group">
            <div class="input-group-text">
               <input class="form-check-input mt-0" type="radio" value=""
                  aria-label="Radio button for following text input">
            </div>
            <input type="text" class="form-control" aria-label="Text input with radio button">
         </div>
      </div>
      <LoadingSpinner v-if="loading"></LoadingSpinner>
      <div v-else class="product-block">
         <Product v-for="product in products" :key="product.id" :product_data="product" />
      </div>
   </div>
</template>
   
<script>
import Layout from '../Layouts/Layout.vue'
import MainSlider from './sliders/MainSlider.vue'
import Product from './Product.vue'
import getParts from '../api/parts'
import LoadingSpinner from './LoadingSpinner.vue';
export default {
   name: 'HomeV',
   components: {
      Layout,
      MainSlider,
      Product,
      LoadingSpinner,
   },
   data() {
      return {
         products: [
         ],
         loading: true
      }
   },
   async mounted() {
      const resp = await getParts()
      this.products = resp.data
      this.loading = false
   },
}
</script>
<style>
img {
   width: 100%;
   height: 100%;
}

.container {
   margin: 0 auto;
   margin-top: 5%;
   width: 85%;
}

.slider {
   margin-top: 2%;
   height: 200px;
}

.slider-text {
   text-align: center;
}

.product-block {
   margin-top: 5%;
   display: flex;
   justify-content: space-between;
   flex-direction: row;
   flex-wrap: wrap;
   text-align: center;
}

@media screen and (max-width: 600px) {
   .product-block {
      margin-top: 0%;
   }

   .slider {
      height: 110px;
   }

   .preLoader {
      margin-top: 10vh;
      align-items: flex-start;
   }
}
</style>