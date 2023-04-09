<template>
    <div class="product">
    <v-popup
        v-if="isProductInfoVisible"
        @closePopup="closePopup"
    />
     <div class="image"><img alt="photo" :src="'data:image/jpeg;base64,/'+ product_data.image "></div>
     <div class="price"><p class="text">Цена: {{product_data.price}}</p></div>
     <div class="description"><p class="text">{{product_data.description}}</p></div>
         <div class="button-block">
            <button class="product-button">Купить</button>
            <button 
            class="product-button"
            @click="showProductInfo"
            >
            Информация
            </button>
        </div>
    </div>
</template>
   
<script>
   import vPopup from '../components/popups/v-popup.vue'

   export default {
       name: 'Product-vue',
       components:{
        vPopup,
       },
       props: {
        product_data:{
            type: Object,
            default() {
                return {}
            }
        }
       },
       data() { 
        return {
            isProductInfoVisible: false,
        } 
     },
     methods: {
        showProductInfo(){
            this.isProductInfoVisible = true
        },
        closePopup(){
            this.isProductInfoVisible = false
        }
 }
}

</script>

<style>
  .product{
      color: aliceblue;
      width: 23%;
      margin: 10px 10px;
   }

   .product-button{
      margin: 0 5px;
   }

   .product-image{
      width: 100%;
   }
</style>
   